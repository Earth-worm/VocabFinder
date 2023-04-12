package usecase

import (
	"context"

	"github.com/Earth-worm/VocabFinder/domain/line"
	"github.com/Earth-worm/VocabFinder/usecase/input"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Test(ctx context.Context, global *input.Global, replyToken string, message *linebot.TextMessage) (responseBody string, err error) {
	err = line.ReplyMessage(global.LineConfig, replyToken, "hello world")
	if err != nil {
		return "", err
	}
	return "", nil
}
