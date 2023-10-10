package config

type System struct {
	Addr string `json:"addr" yaml:"addr"`
}

type Server struct {
	System    System `json:"system" yaml:"system"`
	DB        DB     `yaml:"db"`
	JwtSecret string `json:"jwtSecret" yaml:"jwtSecret"`
}

type DB struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dataname string `yaml:"dataname"`
}
