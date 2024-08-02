package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrUnpackAccessData = errors.New("cant decode access data")

const (
	numOfCallbackParams = 3
	numOfMessageParams  = 4
)

type AccessData struct {
	FolderID    string
	FolderName  string
	AccessLevel AccessLevel
	UserID      int
	Username    string
}

func CreateAccessData(folderID string, folderName string,
	accessLvl AccessLevel, userID int, username string) *AccessData {

	return &AccessData{
		FolderID:    folderID,
		FolderName:  folderName,
		AccessLevel: accessLvl,
		UserID:      userID,
		Username:    username,
	}
}

// UnpackAccessData returns folderID, userID, accessLvl from callbackData and
// username, folderName from messageData.
func UnpackAccessData(callbackData string, message string) (*AccessData, error) {
	callbackParam := strings.Split(callbackData, ",")
	if len(callbackParam) != numOfCallbackParams {
		return nil, ErrUnpackAccessData
	}

	messageParam := strings.Split(message, "'")
	if len(messageParam) != numOfMessageParams {
		return nil, ErrUnpackAccessData
	}

	folderID := callbackParam[1]
	accessLevel := ToAccessLvl(callbackParam[2])
	username, folderName := messageParam[1], messageParam[3]

	userID, err := strconv.Atoi(callbackParam[2])
	if err != nil {
		return nil, err
	}

	return &AccessData{
		FolderID:    folderID,
		FolderName:  folderName,
		AccessLevel: accessLevel,
		UserID:      userID,
		Username:    username,
	}, nil
}

func (data *AccessData) PackCallbackData() string {
	return fmt.Sprintf("%s,%s,%d,%s", GetAccessCmd, data.FolderID, data.UserID, data.AccessLevel)
}
