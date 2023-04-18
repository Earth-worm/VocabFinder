package usecase

import (
	"context"
	"fmt"

	"github.com/Earth-worm/VocabFinder/domain/dictionary"
	"github.com/Earth-worm/VocabFinder/domain/line"
	"github.com/Earth-worm/VocabFinder/usecase/input"
)

func SendWordDetail(ctx context.Context, global *input.Global, replyToken string, word string) (responseBody string, err error) {
	words, err := dictionary.FindWord(word)
	if err != nil {
		return "", err
	}
	if len(words) == 0 {
		err = line.ReplyMessage(global.LineConfig, replyToken, fmt.Sprintf("not found '%s'", word))
		if err != nil {
			return "", err
		}
		return "ok", nil
	}
	msgTemplate := line.AssembleWordDetailCarouselTemplate(words)
	err = line.ReplyTemplate(global.LineConfig, replyToken, msgTemplate)
	if err != nil {
		return "", err
	}
	return "ok", nil
}
