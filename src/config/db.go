package config

type db struct {
	DriverName  string `json:"DriverName"`
	Host        string `json:"Host"`
	Port        string `json:"Port"`
	DBName      string `json:"DBName"`
	User        string `json:"User"`
	PW          string `json:"PW"`
	AdminDBName string `json:"AdminDBName"`
}

type redis struct {
	Host string `json:"Host"`
	Port string `json:"Port"`
	PW   string `json:"PW"`
}

const (
	// redis key and the format of value
	RedisDefaultExpire = 3600 * 24 * 7
)
