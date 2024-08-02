package tgclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/hahaclassic/elysium/pkg/errwrap"
	"github.com/mailru/easyjson"
)

var (
	ErrCreateHTTPRequest = errors.New("error creating http request")
	ErrExecHTTPRequest   = errors.New("error executing http request")
	ErrReadData          = errors.New("error reading response body")
	ErrMarshalJSON       = errors.New("error while marshaling json")
	ErrUnmarshalJSON     = errors.New("error while unmarshaling json")

	ErrNoData      = errors.New("error no data")
	ErrInvalidData = errors.New("error invalid data")
)

const (
	getUpdatesMethod          = "getUpdates"
	sendMessageMethod         = "sendMessage"
	AnswerCallbackQueryMethod = "answerCallbackQuery"
	deleteMessageMethod       = "deleteMessage"
	editMessageMethod         = "editMessageText"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: "bot" + token,
		client:   http.Client{},
	}
}

func (c *Client) Updates(ctx context.Context, offset int, limit int) (upd []*Update, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("[Updates]: %w", err)
		}
	}()

	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doGetRequest(ctx, getUpdatesMethod, q)
	if err != nil {
		return nil, errwrap.Wrap(ErrExecHTTPRequest, err)
	}

	res := &UpdatesResponse{}

	if err := easyjson.Unmarshal(data, res); err != nil {
		return nil, errwrap.Wrap(ErrUnmarshalJSON, err)
	}

	return res.Result, nil
}

// func (c *Client) SendMessage(ctx context.Context, chatID int, text string) (err error) {
// 	defer func() {
// 		if err != nil {
// 			err = fmt.Errorf("[SendMessage]: %w", err)
// 		}
// 	}()

// 	data := &OutputMessage{
// 		ChatID:    chatID,
// 		Text:      text,
// 		ParseMode: "HTML",
// 	}

// 	EncodedData, err := easyjson.Marshal(data)
// 	if err != nil {
// 		return errwrap.Wrap(ErrMarshalJSON, err)
// 	}

// 	_, err = c.doPostRequest(sendMessageMethod, EncodedData)
// 	if err != nil {
// 		return errwrap.Wrap(ErrExecHTTPRequest, err)
// 	}

// 	return nil
// }

// func CreateInlineKeyboardMarkup(buttons [][]*InlineKeyboardButton) (keyboard *InlineKeyboardMarkup, err error) {
// 	defer func() {
// 		if err != nil {
// 			err = fmt.Errorf("[CreateInlineKeyboardMarkup]: %w", err)
// 		}
// 	}()

// 	// for i := 0; i < len(buttonsText); i++ {
// 	// 	inline := []InlineKeyboardButton{}
// 	// 	inline = append(inline, InlineKeyboardButton{
// 	// 		Text:         buttonsText[i],
// 	// 		CallbackData: callbackData[i],
// 	// 	})
// 	// 	buttons = append(buttons, inline)
// 	// }

// 	replyMarkup := &InlineKeyboardMarkup{
// 		InlineKeyboard: buttons}

// 	return replyMarkup, nil
// }

// data := OutputMessage{
// 	ChatID:      chatID,
// 	Text:        text,
// 	ParseMode:   "HTML",
// 	ReplyMarkup: replyMarkup,
// }

func (c *Client) SendMessage(ctx context.Context, msg *OutputMessage) (sendedMsg *OutputMessage, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("[SendCallbackMessage]: %w", err)
		}
	}()

	encodedData, err := easyjson.Marshal(msg)
	if err != nil {
		return nil, errwrap.Wrap(ErrMarshalJSON, err)
	}

	bodyData, err := c.doPostRequest(ctx, sendMessageMethod, encodedData)
	if err != nil {
		return nil, errwrap.Wrap(ErrExecHTTPRequest, err)
	}

	res := PostRequestResponse{}
	if err := json.Unmarshal(bodyData, &res); err != nil {
		return nil, errwrap.Wrap(ErrUnmarshalJSON, err)
	}

	return res.Result, nil
}

func (c *Client) EditMessage(ctx context.Context, msg *OutputMessage) error {
	encodedData, err := easyjson.Marshal(msg)
	if err != nil {
		return errwrap.Wrap(ErrMarshalJSON, err)
	}

	bodyData, err := c.doPostRequest(ctx, editMessageMethod, encodedData)
	if err != nil {
		return errwrap.Wrap(ErrExecHTTPRequest, err)
	}

	res := PostRequestResponse{}
	if err := json.Unmarshal(bodyData, &res); err != nil {
		return errwrap.Wrap(ErrUnmarshalJSON, err)
	}

	return nil
}

func (c *Client) DeleteMessage(ctx context.Context, chatID int, messageID int) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("message_id", strconv.Itoa(messageID))

	_, err := c.doGetRequest(ctx, deleteMessageMethod, q)

	return err
}

func (c *Client) AnswerCallbackQuery(ctx context.Context, callbackQueryID string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("[AnswerCallbackQuery]: %w", err)
		}
	}()

	q := url.Values{}
	q.Add("callback_query_id", callbackQueryID)

	_, err = c.doGetRequest(ctx, AnswerCallbackQueryMethod, q)

	return err
}

// // doPostRequest() sends a post request to the server. Accepts data in json format
func (c *Client) doPostRequest(ctx context.Context, method string, jsonData []byte) (data []byte, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("[doPostRequest]: %w", err)
		}
	}()

	url := c.createMethodURL(method)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, errwrap.Wrap(ErrCreateHTTPRequest, err)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errwrap.Wrap(ErrExecHTTPRequest, err)
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errwrap.Wrap(ErrReadData, err)
	}

	return body, nil
}

func (c *Client) doGetRequest(ctx context.Context, method string, query url.Values) (data []byte, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("[doGetRequest]: %w", err)
		}
	}()

	url := c.createMethodURL(method)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errwrap.Wrap(ErrCreateHTTPRequest, err)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errwrap.Wrap(ErrExecHTTPRequest, err)
	}

	defer func() { _ = resp.Body.Close() }()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, errwrap.Wrap(ErrReadData, err)
	}

	return data, nil
}

func (c *Client) createMethodURL(method string) string {
	methodURL := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	return methodURL.String()
}
