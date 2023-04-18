package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Earth-worm/VocabFinder/domain/aws"
	"github.com/Earth-worm/VocabFinder/domain/line"
	"github.com/Earth-worm/VocabFinder/usecase/input"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"go.uber.org/zap"
)

type LambdaStartHandler func(context.Context, events.APIGatewayProxyRequest) (interface{}, error)

func Middleware(next LambdaStartHandler) LambdaStartHandler {
	return LambdaStartHandler(func(ctx context.Context, request events.APIGatewayProxyRequest) (interface{}, error) {
		return next(ctx, request)
	})
}

type Handler func(context.Context, events.APIGatewayProxyRequest, *input.Global) (events.APIGatewayProxyResponse, error)

func MiddlewareHandler(h Handler) LambdaStartHandler {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (interface{}, error) {
		var err error
		logger, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		zap.ReplaceGlobals(logger)

		lc, ok := lambdacontext.FromContext(ctx)

		if ok != true {
			err = fmt.Errorf("middleware lambda context error")
			logger.Error("new logger error", zap.Error(err))
			return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: err.Error()}, nil
		}
		// line
		lineChannelSecret, err := aws.GetParameter("VOCAB_FINDER_LINE_CHANNEL_SECRET")
		if err != nil {
			logger.Error("new logger error", zap.Error(err))
			return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
		}
		lineChannelToken, err := aws.GetParameter("VOCAB_FINDER_LINE_CHANNEL_TOKEN")
		if err != nil {
			zap.L().Error(err.Error())
			log.Print(err)
			log.Print(err)
			return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
		}
		lineConfig := line.NewConfig(lineChannelSecret, lineChannelToken)

		global := &input.Global{
			LambdaContext: lc,
			LineConfig:    lineConfig,
		}
		return h(ctx, request, global)
	}
}