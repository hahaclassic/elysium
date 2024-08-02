package service

import (
	"context"

	"github.com/google/uuid"
	tgclient "github.com/hahaclassic/elysium/internal/client/telegram"
	"github.com/hahaclassic/elysium/internal/model"
)

func (p *Processor) Menu(ctx context.Context, chatID int) error {
	keyboard := [][]*tgclient.InlineKeyboardButton{
		{
			{
				Text:         "Folders",
				CallbackData: "/folders",
			},
		},
		{
			{
				Text:         "Keys",
				CallbackData: "/keys",
			},
		},
		{
			{
				Text:         "Settings",
				CallbackData: "/settings",
			},
		},
	}

	msg := &tgclient.OutputMessage{
		ChatID: chatID,
		Text:   "main",
		ReplyMarkup: &tgclient.InlineKeyboardMarkup{
			InlineKeyboard: keyboard,
		},
	}

	p.tg.SendMessage(ctx, msg)
}

func (p *Processor) AllFolders(ctx context.Context, userID int64, chatID int) error {
	folders, err := p.storage.GetAllUserFolders(ctx, userID)
	if err != nil {
		return err
	}

	keyboard := [][]*tgclient.InlineKeyboardButton{}
	for _, folder := range folders {
		keyboard = append(keyboard, []*tgclient.InlineKeyboardButton{
			{
				Text:         folder.Name,
				CallbackData: folder.ID.String(),
			},
		})
	}

	msg := &tgclient.OutputMessage{
		ChatID: chatID,
		Text:   "main",
		ReplyMarkup: &tgclient.InlineKeyboardMarkup{
			InlineKeyboard: keyboard,
		},
	}

	p.tg.SendMessage(ctx, msg)
}

func (p *Processor) Folder(ctx context.Context, chatID int, userID int64, folderID uuid.UUID) error {
	accessLevel, err := p.storage.GetAccessLevel(ctx, userID, folderID)
	if err != nil {
		return err
	}

	keyboard := [][]*tgclient.InlineKeyboardButton{}

	switch accessLevel {
	case model.OwnerLvl:
		keyboard = [][]*tgclient.InlineKeyboardButton{
			{
				{
					Text:         "Keys",
					CallbackData: "/keys",
				},
			},
			{
				{
					Text:         "Update folder name",
					CallbackData: "/update_folder_name",
				},
				{
					Text:         "Update tag",
					CallbackData: "/update_tag",
				},
			},
			{
				{
					Text:         "Delete folder",
					CallbackData: "/delete_folder",
				},
				{
					Text:         "Delete link",
					CallbackData: "/delete_links",
				},
			},
		}
	case model.EditorLvl:
		keyboard = [][]*tgclient.InlineKeyboardButton{
			{
				{
					Text:         "Update tag",
					CallbackData: "/update_tag",
				},
			},
			{
				{
					Text:         "Delete link",
					CallbackData: "/delete_links",
				},
			},
		}
	}

	keyboard = append(keyboard, []*tgclient.InlineKeyboardButton{
		{
			Text:         "Back to all folders",
			CallbackData: "/back",
		},
	})

	links, err := p.storage.GetAllUserFolders(ctx, userID)
	if err != nil {
		return err
	}

	keyboard := [][]*tgclient.InlineKeyboardButton{}
	for _, folder := range folders {
		keyboard = append(keyboard, []*tgclient.InlineKeyboardButton{
			{
				Text:         folder.Name,
				CallbackData: folder.ID.String(),
			},
		})
	}

	msg := &tgclient.OutputMessage{
		ChatID: chatID,
		Text:   "main",
		ReplyMarkup: &tgclient.InlineKeyboardMarkup{
			InlineKeyboard: keyboard,
		},
	}

	p.tg.SendMessage(ctx, msg)
}

func (p *Processor) Keys(ctx context.Context, chatID int) {
}
