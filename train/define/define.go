package define

type DatabaseConfig struct {
	User     string `config:"user" validate:"required"`
	Password string `config:"password" validate:"required"`
	Host     string `config:"host" validate:"required"`
	Port     int    `config:"port"`
	DbName   string `config:"dbname" validate:"required"`
}

type LogConfig struct {
	Dir        string `config:"dir" validate:"required"`
	MaxSize    int    `config:"max-logger-size"`
	MaxBackups int    `config:"max-logger-backups"`
	MaxAge     int    `config:"days-to-keep"`
}
type AgentConfig struct {
	Host string `config:"host"`
}

type EmConfig struct {
	DockerCompose string `config:"docker-compose"`
}

type Operator struct {
	Product string   `config:"product-name"`
	Action  string   `config:"action"`
	Order   []string `config:"order"`
}

type Config struct {
	MysqlDb      DatabaseConfig `config:"mysqldb" validate:"required"`
	Log          LogConfig      `config:"log" validate:"required"`
	Agent        AgentConfig    `config:"agent" validate:"required"`
	Em           EmConfig       `config:"em" validate:"required"`
	OperatorList []Operator     `config:"operators" validate:"required"`
}
