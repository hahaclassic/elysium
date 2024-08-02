package tgclient

type UpdatesResponse struct {
	Ok     bool      `json:"ok"`
	Result []*Update `json:"result"`
}

type PostRequestResponse struct {
	Ok     bool           `json:"ok"`
	Result *OutputMessage `json:"result"`
}

type Update struct {
	ID            int            `json:"update_id"`
	Message       *InputMessage  `json:"message"`
	CallbackQuery *CallbackQuery `json:"callback_query"`
}

type CallbackQuery struct {
	QueryID string        `json:"id"`
	From    *From         `json:"from"`
	Message *InputMessage `json:"message"`
	Data    string        `json:"data"`
}

type InputMessage struct {
	Text     string           `json:"text"`
	From     *From            `json:"from"`
	Chat     *Chat            `json:"chat"`
	Entities []*MessageEntity `json:"entities"` // Tempo
}

type MessageEntity struct {
	Type          string `json:"type"`
	Offset        int    `json:"offset"`
	Length        int    `json:"length"`
	URL           string `json:"url"`
	CustomEmojiID string `json:"custom_emoji_id"`
}

type From struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
}

type OutputMessage struct {
	ChatID      int                   `json:"chat_id"`
	Text        string                `json:"text"`
	MessageID   int                   `json:"message_id"`
	Entities    []*MessageEntity      `json:"entities,omitempty"`
	ParseMode   string                `json:"parse_mode"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}
