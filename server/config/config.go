package config

type DBConfig struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

func LoadDBConfig() *DBConfig {
	return &DBConfig{
		Host:     "localhost",
		Port:     "3306",
		Name:     "tododb",
		Username: "tododbuser",
		Password: "GHcn=zKhNr5o",
	}
}
