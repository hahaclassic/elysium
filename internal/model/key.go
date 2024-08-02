package model

import "github.com/google/uuid"

type Key struct {
	FolderID  uuid.UUID
	AccessLvl AccessLevel
	Key       string
}
