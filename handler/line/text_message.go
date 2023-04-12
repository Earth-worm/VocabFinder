package main

import (
	"context"

	"github.com/Earth-worm/VocabFinder/domain/line"
	"github.com/Earth-worm/VocabFinder/usecase"
	"github.com/Earth-worm/VocabFinder/usecase/input"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func MessageTextHandler(ctx context.Context, global *input.Global, replyToken string, message *linebot.TextMessage) (responseBody string, err error) {
	switch message.Text {
	case line.TextMessageTriggerTest.String():
		responseBody, err = usecase.Test(ctx, global, replyToken, message)
	}
	if err != nil {
		return "", err
	}
	return "ok", nil
}
