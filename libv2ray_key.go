package libv2ray

const APP_KEY = "573d23e9ff036761bfb8d179717453173d8191834b73cb9388d12a00228cb62e"
const SERVER_KEY = "afa2d1267fa1609aa58e268312ca2940d7793c5688eb6a2fe9e00b5f9f68e014"
const APP_STORE_FINGERPRINT = "C0:6F:83:53:83:12:07:D2:D6:76:C7:4C:E6:89:57:BB:B4:18:C6:23:EA:91:9C:9F:AE:D5:B4:F9:C1:89:22:9F"

type KeyService interface {
	GetAppKey() string
	GetServerKey() string
}

func GetAppKey() string {
	return APP_KEY
}

func GetServerKey() string {
	return SERVER_KEY
}

func GetFingerPrint() string {
	return APP_STORE_FINGERPRINT
}