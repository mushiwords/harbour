package config

type MyDB struct {
	DbPath   string `json:"DbPath"`
	UserName string `json:"UserName"`
	PassWord string `json:"PassWord"`
	DbName   string `json:"DbName"`
}

type MyLog struct {
	LogLevel  string `json:"LogLevel"`
	AccLogFile   string `json:"AccLogFile"`
	InfoLogFile   string `json:"InoLogFile"`
	LogMaxAge int    `json:"LogMaxAge"`
}

type Service struct {
	ListenPort    int    `json:"ListenPort"`    // Server 端口
}

type OpenApi struct {
	Scheme   string `json:"Scheme"`
	EndPoint string `json:"EndPoint"`
	AK       string `json:"AK"`
	SK       string `json:"SK"`
}

type OSS struct {
	Endpoint   string `json:"Endpoint"`
	Region     string `json:"Region"`
	AK         string `json:"AK"`
	SK         string `json:"SK"`
	DisableSSL bool   `json:"DisableSSL"`
	Bucket     string `json:"Bucket"`
}

type Redis struct {
	Addr     string `json:"Addr"`
	PassWord string `json:"PassWord"`
	Timeout  int    `json:"Timeout"`
}

type Config struct {
	MyDB      *MyDB      `json:"MyDB"`
	MyLog     *MyLog     `json:"MyLog"`
	Service   *Service   `json:"Service"`
	OpenApi   *OpenApi   `json:"OpenApi"`
	OSS       *OSS       `json:"OSS"`
	Redis     *Redis     `json:"Redis"`
}
