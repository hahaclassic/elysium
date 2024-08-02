package dictionary

import "errors"

const (
	EngLang = "eng"
	RusLang = "rus"
)

var (
	ErrLoadingDictionary = errors.New("err loading dictionary")
	ErrInvalidKey        = errors.New("err invalid key")
)

type MessagesDictionary interface {
	Message(language string, messageName string) (string, error)
}
