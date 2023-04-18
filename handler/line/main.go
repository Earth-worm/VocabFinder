package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Earth-worm/VocabFinder/domain/line"
	"github.com/Earth-worm/VocabFinder/middleware"
	"github.com/Earth-worm/VocabFinder/usecase/input"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.uber.org/zap"
)

func LineHandler(ctx context.Context, request events.APIGatewayProxyRequest, global *input.Global) (events.APIGatewayProxyResponse, error) {
	log.Printf("RequestBody: %+v", request.Body)
	var responseBody string
	var err error

	lineEvents, err := line.ParseRequest(global.LineConfig.ChannelSecret, &request)
	if err != nil {
		zap.L().Error(fmt.Sprintf("aws request id:%s", global.LambdaContext.AwsRequestID), zap.Error(err))
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: err.Error()}, nil
	}
	for _, event := range lineEvents {
		switch event.Type {
		case linebot.EventTypeMessage:
			responseBody, err = MessageHandler(ctx, global, event)
			/*
				case linebot.EventTypeFollow:
				case linebot.EventTypeUnfollow:
				case linebot.EventTypeJoin:
				case linebot.EventTypeLeave:
				case linebot.EventTypeMemberJoined:
				case linebot.EventTypeMemberLeft:
				case linebot.EventTypePostback:
				case linebot.EventTypeBeacon:
				case linebot.EventTypeAccountLink:
				case linebot.EventTypeThings:
				case linebot.EventTypeUnsend:
				case linebot.EventTypeVideoPlayComplete:
			*/
		}
		if err != nil {
			zap.L().Error("error msg", zap.Error(err))
			return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: err.Error()}, nil
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       responseBody,
	}, nil
}

func main() {
	lambda.Start(
		middleware.Middleware(
			middleware.MiddlewareHandler(
				LineHandler,
			),
		),
	)
}