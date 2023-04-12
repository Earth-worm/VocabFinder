package main

import (
	"context"

	"github.com/Earth-worm/VocabFinder/usecase/input"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func MessageHandler(ctx context.Context, global *input.Global, event *linebot.Event) (responseBody string, err error) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		responseBody, err = MessageTextHandler(ctx, global, event.ReplyToken, message)
		/*
			case linebot.MessageTypeImage:
			case linebot.MessageTypeVideo:
			case linebot.MessageTypeAudio:
			case linebot.MessageTypeFile:
			case linebot.MessageTypeLocation:
			case linebot.MessageTypeSticker:
			case linebot.MessageTypeTemplate:
			case linebot.MessageTypeImagemap:
			case linebot.MessageTypeFlex:
		*/
	}
	if err != nil {
		return "", err
	}
	return responseBody, nil
}