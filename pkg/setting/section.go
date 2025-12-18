package setting

type Config struct {
	Server     ServerSetting     `mapstructure:"server"`
	Logger     LoggerSetting     `mapstructure:"logger"`
	PostgreSQL PostgreSQLSetting `mapstructure:"postgresql"`
}

type ServerSetting struct {
	Port string `mapstructure:"port"`
}

type PostgreSQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	Log_level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_size      int    `mapstructure:"max_size"`
	Max_age       int    `mapstructure:"max_age"`
	Compress      bool   `mapstructure:"compress"`
}
