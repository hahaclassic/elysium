package cache

type SettingsCache interface {
	SetLanguage()
	SetOutputFormat()

	GetLanguage()
	GetOutputFormat()
}
