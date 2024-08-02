package model

type UserSettings struct {
	ID       int64
	Name     int64
	Premium  bool
	Format   string
	Language string
}

type UserCacheSettings struct {
	Format   string
	Language string
}
