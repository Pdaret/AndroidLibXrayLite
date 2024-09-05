package libv2ray

const SERVER_KEY = "afa2d1267fa1609aa58e268312ca2940d7793c5688eb6a2fe9e00b5f9f68e014"

type KeyService interface {
	GetAppKey() string
	GetServerKey() string
}

func GetServerKey() string {
	return SERVER_KEY
}
