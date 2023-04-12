package line

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// ParseRequest func
func ParseRequest(channelSecret string, r *events.APIGatewayProxyRequest) ([]*linebot.Event, error) {
	byteBody := []byte(r.Body)
	if !validateSignature(channelSecret, r.Headers["x-line-signature"], byteBody) {
		return nil, linebot.ErrInvalidSignature
	}

	request := &struct {
		Events []*linebot.Event `json:"events"`
	}{}
	if err := json.Unmarshal(byteBody, request); err != nil {
		return nil, err
	}
	return request.Events, nil
}

func validateSignature(channelSecret, signature string, body []byte) bool {
	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}
	hash := hmac.New(sha256.New, []byte(channelSecret))

	_, err = hash.Write(body)
	if err != nil {
		return false
	}

	return hmac.Equal(decoded, hash.Sum(nil))
}