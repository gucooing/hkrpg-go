package pkg

const (
	AppVersion  = "1.8.0-dev"
	GameVersion = "2.7.5x"
)

func GetAppVersion() string {
	return AppVersion
}

func GetGameVersion() string {
	return GameVersion
}
