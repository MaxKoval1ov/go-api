package lib

// Env has environment stored
type Env struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENV"`
	LogOutput   string `mapstructure:"LOG_OUTPUT"`
	DBUsername  string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASS"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBName      string `mapstructure:"DB_NAME"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
}

// NewEnv creates a new environment
func Config() Env {

	env := Env{
		ServerPort:  "5432",
		Environment: "DEVELOPMENT",
		DBUsername:  "admin",
		DBPassword:  "admin",
		DBHost:      "localhost",
		DBPort:      "5432",
		DBName:      "test",
		JWTSecret:   "SUPER_SECRET",
	}
	// viper.SetConfigFile(".env")

	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Fatal("☠️ cannot read configuration")
	// }

	// err = viper.Unmarshal(&env)
	// if err != nil {
	// 	log.Fatal("☠️ environment can't be loaded: ", err)
	// }

	return env
}
