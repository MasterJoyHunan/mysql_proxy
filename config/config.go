package config

type Log struct {
	Dir   string
	Level string
}

type ProxyMysql struct {
	Host string
	Port int
	User string
	Pwd  string
}

type RealMysql struct {
	Host string
	Port int
	User string
	Pwd  string
	Db   string
}
