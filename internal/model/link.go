package model

import "github.com/google/uuid"

type Link struct {
	URL      string
	Tag      string
	FolderID uuid.UUID
}
