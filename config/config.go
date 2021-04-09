package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var (
	// Server ip and port
	ServerAddress string

	// Name godoc
	Name string

	// Age godoc
	Age int

	// Skills godoc
	Skills []interface{}

	// Currencies godoc
	Currencies map[string]interface{}

	// Male godoc
	Male bool
	// Birthday date
	Birthday time.Time

	// Taxonomies godoc
	Taxonomies map[string]string

	DbPort     = []byte("")
	DbHost     = []byte("")
	DbUsername = []byte("")
	DbPassword = []byte("")
	DbName     = []byte("")
)

func readConfigYAML() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func setDefaultValues() {
	viper.SetDefault("name", "Default Name")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
}

// Load config.yaml
func Load() {
	readConfigYAML()
	setDefaultValues()
	ServerAddress = viper.GetString("server_address")
	Name = viper.GetString("name")
	Age = viper.GetInt("age")
	Skills = viper.Get("skills").([]interface{})
	Currencies = viper.Get("currencies").(map[string]interface{})
	Male = viper.GetBool("male")
	Birthday = viper.GetTime("birthday")
	Taxonomies = viper.GetStringMapString("Taxonomies")

	// .env file
	env := godotenv.Load()
	if env != nil {
		fmt.Println("Can't load configurations!")
		fmt.Println(env)
	}

	DbPort = []byte(os.Getenv("db_port"))
	DbHost = []byte(os.Getenv("db_host"))
	DbUsername = []byte(os.Getenv("db_username"))
	DbPassword = []byte(os.Getenv("db_password"))
	DbName = []byte(os.Getenv("db_name"))
}
