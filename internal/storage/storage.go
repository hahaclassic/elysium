package storage

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hahaclassic/elysium/internal/model"
)

type LinkStorage interface {
	CreateLink(ctx context.Context, link *model.Link) error
	DeleteLink(ctx context.Context, url string, folderID uuid.UUID) error
	UpdateTag(ctx context.Context, link *model.Link) error
	PickRandom(ctx context.Context, userID int64) error
	GetLinks(ctx context.Context, folderID uuid.UUID) ([]*model.Link, error)
}

type FolderStorage interface {
	CreateFolder(ctx context.Context, folder *model.Folder) error
	DeleteFolder(ctx context.Context, folderID uuid.UUID) error
	Owner(ctx context.Context, folderID string) (userID int64, err error) // хз
	RenameFolder(ctx context.Context, folderID uuid.UUID, newName string) error
	GetAllUserFolders(ctx context.Context, userID int64) ([]*model.Folder, error)
}

type AccessStorage interface {
	GetAccessLevel(ctx context.Context, userID int64, folderID uuid.UUID) (model.AccessLevel, error)
	SetAccess(ctx context.Context, userID int64, folderID uuid.UUID, lvl model.AccessLevel) error
	DeleteAccess(ctx context.Context, userID int64, folderID uuid.UUID) error

	SetKey(ctx context.Context, key *model.Key) error
	DeleteKey(ctx context.Context, key string) error
	GetKeys(ctx context.Context, folderID uuid.UUID) (keys []string, err error)
}

type SettingsStorage interface {
	SetUsername(ctx context.Context, userID int64, username string) error
	SetLanguage(ctx context.Context, userID int64, language string) error
	SetOutputFormat(ctx context.Context, userID int64, format string) error
	SetPremium(ctx context.Context, userID int64, isPremium bool) error
	UserSettings(ctx context.Context, userID int64) (*model.UserSettings, error)
}

type Storage interface {
	LinkStorage
	FolderStorage
	AccessStorage
	SettingsStorage
}

// type Storage interface {
// 	CreateLink(ctx context.Context, link *model.Link) error
// 	DeleteLink(ctx context.Context, url string, folderID uuid.UUID) error
// 	UpdateTag(ctx context.Context, link *model.Link) error
// 	PickRandom(ctx context.Context, userID int64) error
// 	GetLinks(ctx context.Context, folderID uuid.UUID) ([]*model.Link, error)

// 	CreateFolder(ctx context.Context, folder *model.Folder) error
// 	DeleteFolder(ctx context.Context, folderID uuid.UUID) error
// 	Owner(ctx context.Context, folderID string) (userID int64, err error) // хз
// 	RenameFolder(ctx context.Context, folderID uuid.UUID, newName string) error
// 	GetAllUserFolders(ctx context.Context, userID int64) ([]*model.Folder, error)

// 	GetAccessLevel(ctx context.Context, userID int64, folderID uuid.UUID) (AccessLevel, error)
// 	SetAccess(ctx context.Context, userID int64, jwtClaims *model.JWTClaims) error
// 	DeleteAccess(ctx context.Context, userID int64, folderID uuid.UUID) error

// 	SetKey(ctx context.Context, jwtClaims *model.JWTClaims, key string) error
// 	DeleteKey(ctx context.Context, jwtClaims *model.JWTClaims) error
// 	GetKeys(ctx context.Context, folderID uuid.UUID) (keys []string, err error)

// 	ChangeLanguage(ctx context.Context, userID int64, language string) error
// 	UserLanguage(ctx context.Context, userID int64) (lang string)
// 	SetOutputFormat(ctx context.Context, userID int64, format string) error
// 	UserOutputFormat(ctx context.Context, userID int64) (format string)
// }

var (
	ErrNoFolders         = errors.New("no folders")
	ErrNoSavedPages      = errors.New("no saved pages")
	ErrIvalidAccessLvl   = errors.New("invalid access level")
	ErrNoPasswords       = errors.New("no passwords")
	ErrNoRows            = errors.New("err no rows")
	ErrNoRowsAffected    = errors.New("no rows affected")
	ErrNotFound          = errors.New("not found")
	ErrStorageConnection = errors.New("error connecting to the storage")
)
