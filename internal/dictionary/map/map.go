package dictionarymap

import (
	"errors"

	"github.com/hahaclassic/elysium/config"
	"github.com/hahaclassic/elysium/internal/dictionary"
	"github.com/hahaclassic/elysium/pkg/errwrap"
	"github.com/hahaclassic/elysium/pkg/syncmap"
	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EngLang = "eng"
	RusLang = "rus"
)

var ErrGetMessage = errors.New("can't get message from map")

type BotMessage struct {
	Name    string `json:"msg_name"`
	EngText string `json:"eng_text"`
	RusText string `json:"rus_text"`
}

type MessagesDictionary struct {
	dict *syncmap.Map
}

func New(conf *config.DictConfig) (*MessagesDictionary, error) {
	msgDict := &MessagesDictionary{
		dict: syncmap.NewMap(),
	}
	messages := []*BotMessage{}

	err := cleanenv.ReadConfig(conf.Path, messages)
	if err != nil {
		return nil, errwrap.Wrap(dictionary.ErrLoadingDictionary, err)
	}

	for _, msg := range messages {
		msgDict.dict.Store(msg.Name, msg)
	}

	return msgDict, nil
}

// TODO: convert to tgclient message with premium emoji.
func (m *MessagesDictionary) Message(language string, messageName string) (string, error) {
	msg, ok := m.dict.Load(messageName)
	if !ok {
		return "", dictionary.ErrInvalidKey
	}

	msgValue, ok := msg.(*BotMessage)
	if !ok {
		return "", ErrGetMessage
	}

	switch language {
	case RusLang:
		return msgValue.RusText, nil

	default:
		return msgValue.EngText, nil
	}
}
