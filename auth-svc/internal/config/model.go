package config

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Logger   LoggerConfig   `yaml:"logger"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	Auth     AuthConfig     `yaml:"auth"`
}

type LoggerConfig struct {
	Path                 string   `yaml:"path"`
	File                 string   `yaml:"file"`
	Prefix               string   `yaml:"prefix"`
	Level                string   `yaml:"level"`
	TimeFormat           string   `yaml:"timeFormat"`
	CtxKeys              []string `yaml:"ctxKeys"`
	Header               bool     `yaml:"header"`
	StSkip               int      `yaml:"stSkip"`
	Stdout               bool     `yaml:"stdout"`
	RotateSize           int      `yaml:"rotateSize"`
	RotateExpire         int      `yaml:"rotateExpire"`
	RotateBackupLimit    int      `yaml:"rotateBackupLimit"`
	RotateBackupExpire   int      `yaml:"rotateBackupExpire"`
	RotateBackupCompress int      `yaml:"rotateBackupCompress"`
	RotateCheckInterval  string   `yaml:"rotateCheckInterval"`
	StdoutColorDisabled  bool     `yaml:"stdoutColorDisabled"`
	WriterColorEnable    bool     `yaml:"writerColorEnable"`
	Flags                int      `yaml:"flags"`
}
type ServerConfig struct {
	Address         string `yaml:"address"`
	OpenAPIPath     string `yaml:"openapiPath"`
	SwaggerPath     string `yaml:"swaggerPath"`
	DumpRouterMap   bool   `yaml:"dumpRouterMap"`
	ErrorStack      bool   `yaml:"errorStack"`
	ErrorLogEnabled bool   `yaml:"errorLogEnabled"`
	ErrorLogPattern string `yaml:"errorLogPattern"`
}
type DatabaseConfig struct {
	Logger struct {
		Level  string `yaml:"level"`
		Stdout bool   `yaml:"stdout"`
	} `yaml:"logger"`
	Default struct {
		Link        string `yaml:"link"`
		Debug       bool   `yaml:"debug"`
		MaxIdle     string `yaml:"maxIdle"`
		MaxOpen     string `yaml:"maxOpen"`
		MaxLifetime string `yaml:"maxLifetime"`
	} `yaml:"default"`
}

type RedisConfig struct {
	Default struct {
		Address         string `yaml:"address"`
		Db              int    `yaml:"db"`
		IdleTimeout     string `yaml:"idleTimeout"`
		MaxConnLifetime string `yaml:"maxConnLifetime"`
		WaitTimeout     string `yaml:"waitTimeout"`
		DialTimeout     string `yaml:"dialTimeout"`
		ReadTimeout     string `yaml:"readTimeout"`
		WriteTimeout    string `yaml:"writeTimeout"`
		MaxActive       int    `yaml:"maxActive"`
		MaxPoolSize     int    `yaml:"maxPoolSize"`
		MinIdleConns    int    `yaml:"minIdleConns"`
		MaxRetries      int    `yaml:"maxRetries"`
	} `yaml:"default"`
}

type AuthConfig struct {
	SecretKey                string `yaml:"secretKey"`
	AccessTokenExpireMinute  int    `yaml:"accessTokenExpireMinute"`
	RefreshTokenExpireMinute int    `yaml:"refreshTokenExpireMinute"`
}
