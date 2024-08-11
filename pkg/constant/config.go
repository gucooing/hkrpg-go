package constant

type AppList struct {
	App map[string]App `json:"app"`
}
type App struct {
	Port      string `json:"port"`
	InnerAddr string `json:"inner_addr"`
	OuterAddr string `json:"outer_addr"`
}
type MysqlConf struct {
	Dsn string `json:"dsn"`
}
type RedisConf struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}
