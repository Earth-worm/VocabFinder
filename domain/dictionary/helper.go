package dictionary

import (
	"encoding/json"
	"net/http"

	"github.com/Earth-worm/VocabFinder/util"
	"github.com/pkg/errors"
)

const (
	DictionaryApiURL        = "https://api.dictionaryapi.dev/api/v2/entries/en/"
	TitleNoDefinitionsFound = "No Definitions Found"
)

type Word struct {
	Word      string `json:"word"`
	Phonetic  string `json:"phonetic"`
	Phonetics []struct {
		Text      string `json:"text"`
		Audio     string `json:"audio"`
		SourceURL string `json:"sourceUrl"`
		Lisence   []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"lisence"`
	} `json:"phonetics"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		definitions  []struct {
			Definition string   `json:"definition"`
			Synonyms   []string `json:"synonyms"`
			Antonyms   []string `json:"antonyms"`
			Example    string   `json:"example"`
		}
		Synonyms []string `json:"synonyms"`
		Antonyms []string `json:"antonyms"`
	} `json:"meanings"`
}

type ErrorResponse struct {
	Title      string `json:"title"`
	Message    string `json:"message"`
	Resolution string `json:"Resplution"`
}

func FindWord(word string) ([]Word, error) {
	statusCode, body, err := util.HttpRequest(util.MethodGet, DictionaryApiURL+word, func(req *http.Request) (*http.Request, error) { return req, nil })
	if err != nil {
		return nil, err
	}
	if statusCode == 404 {
		errResp, err := ParseErrorResponse(body)
		if err != nil {
			return nil, err
		}
		//単語のスペルが間違っている場合
		if errResp.Title == TitleNoDefinitionsFound {
			return nil, nil
		} else {
			return nil, errors.Errorf("dictionary api error, error message :%s,resolution:%s", errResp.Message, errResp.Resolution)
		}
	}
	words, err := ParseWords(body)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func ParseWords(body []byte) ([]Word, error) {
	var words []Word
	err := json.Unmarshal(body, &words)
	if err != nil {
		return nil, errors.Wrap(err, "parse data error")
	}
	return words, nil
}

func ParseErrorResponse(body []byte) (ErrorResponse, error) {
	var errorResponse ErrorResponse
	err := json.Unmarshal(body, &errorResponse)
	if err != nil {
		return errorResponse, errors.Wrap(err, "parse data error")
	}
	return errorResponse, nil
}