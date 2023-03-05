package settings

/**
 * @Description
 * @Author 枯木.
 * @Date 2023/3/5 13:57
 **/

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
	Port     int    `mapstructure:"port"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"host"`
	DB       string `mapstructure:"host"`
	Port     int    `mapstructure:"host"`
}
