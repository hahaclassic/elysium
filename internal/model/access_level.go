package model

import "fmt"

type AccessLevel int

const (
	UndefinedLvl       AccessLevel = iota
	BannedLvl                      // The folder owner does not receive requests from banned users
	SuspectedLvl                   // Last chance to gain access. In case of another refusal, the status will change to banned
	ReaderLvl                      // Reading only
	ConfirmedReaderLvl             // Reading only after confirmation
	EditorLvl                      // Editor: add/delete links (always after confirmation)
	OwnerLvl                       // Owner: All possible actions with the folder and its contents are available
)

func (lvl AccessLevel) String() string {
	return []string{"Undefined", "Banned", "Suspected", "Reader", "Confirmed reader", "Editor", "Owner"}[lvl]
}

func (lvl AccessLevel) IsValid() bool {
	return lvl >= BannedLvl && lvl <= OwnerLvl
}

func ToAccessLvl(s string) AccessLevel {
	for lvl := UndefinedLvl; lvl <= OwnerLvl; lvl++ {
		if s == fmt.Sprint(lvl) {
			return lvl
		}
	}

	return UndefinedLvl
}
