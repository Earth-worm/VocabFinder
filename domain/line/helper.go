package line

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/pkg/errors"
)

type Config struct {
	ChannelSecret string
	ChannelToken  string
}

type TextMessageTrigger string

func (v TextMessageTrigger) String() string {
	return string(v)
}

func NewConfig(channelSecret, channelToken string) *Config {
	return &Config{
		ChannelToken:  channelToken,
		ChannelSecret: channelSecret,
	}
}

const (
	TextMessageTriggerTest = TextMessageTrigger("テスト")
)

func ReplyMessage(config *Config, replyToken, message string) error {
	bot, err := linebot.New(config.ChannelSecret, config.ChannelToken)
	if err != nil {
		return errors.Wrap(err, "line api reply message error")
	}
	_, err = bot.ReplyMessage(replyToken, linebot.NewTextMessage(message)).Do()
	if err != nil {
		return errors.Wrap(err, "line api reply message error")
	}
	return nil
}

func ReplyTemplate(config *Config, replyToken string, template linebot.Template) error {
	bot, err := linebot.New(config.ChannelSecret, config.ChannelToken)
	if err != nil {
		return errors.Wrap(err, "line api reply template error")
	}
	_, err = bot.ReplyMessage(replyToken, linebot.NewTemplateMessage("template", template)).Do()
	if err != nil {
		return errors.Wrap(err, "line api reply template error")
	}
	return nil
}