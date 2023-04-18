package line

import (
	"fmt"

	"github.com/Earth-worm/VocabFinder/domain/dictionary"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func AssembleWordDetailCarouselTemplate(words []dictionary.Word) *linebot.CarouselTemplate {
	columns := []*linebot.CarouselColumn{}
	for _, word := range words {
		exampleStr := ""
		for _, meaning := range word.Meanings {
			for _, definition := range meaning.Definitions {
				exampleStr += fmt.Sprintf("%s", definition.Example)
			}
		}
		columns = append(columns, linebot.NewCarouselColumn(
			"",
			word.Word,
			exampleStr[0:50],
			linebot.NewMessageAction("Action?", word.Word),
		))
	}
	return linebot.NewCarouselTemplate(columns...)
}
