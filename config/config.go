package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type DBConfig struct{
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMODE bool
}

type Config struct{
	Version  		 string
	ServiceName  string
	HttpPort 		 int
	JwtSecretKey string
	DB           *DBConfig
}

func loadConfig(){
	err := godotenv.Load()
	if err != nil{
		fmt.Println("Failed to load the env variable", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == ""{
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == ""{
		fmt.Println("Service name is required")
		os.Exit(1)
	}

	httpPoert := os.Getenv("HTTP_PORT")
	if httpPoert == ""{
		fmt.Println("HttpPort is required")
		os.Exit(1)
	}
  
	poet, err := strconv.ParseInt(httpPoert, 10, 64)

	if err != nil{
		fmt.Println("Port must me number")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	if jwtSecretKey == ""{
		fmt.Println("Jwt secret key is required")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == ""{
		fmt.Println("Host is required")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == ""{
		fmt.Println("Db Port is required")
		os.Exit(1)
	}

	dbPrt, err := strconv.ParseInt(dbPort,10, 64)
	if err != nil{
		fmt.Println("Port must be Number")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == ""{
		fmt.Println("Db Name is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == ""{
		fmt.Println("Db User is required")
		os.Exit(1)
	}

	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == ""{
		fmt.Println("Db Passord is required")
		os.Exit(1)
	}

	enableSslMode := os.Getenv("DB_ENABLE_SSL_MODE")

	enbSSlMode, err := strconv.ParseBool(enableSslMode)
	if err != nil{
		fmt.Println("Invalid enable ssl mode value : ", err)
		os.Exit(1)
	}
	// if enableSslMode == ""{
	// 	fmt.Println("Db enableSslMode is required")
	// 	os.Exit(1)
	// }

	dbConfig := &DBConfig{
		Host: dbHost,
		Port: int(dbPrt),
		Name: dbName,
		User: dbUser,
		Password: dbPass,
		EnableSSLMODE: enbSSlMode,
  
	}

	configurations = &Config{
		Version:       version,
		ServiceName:  serviceName,
		HttpPort:     int(poet),
		JwtSecretKey: jwtSecretKey,
		DB: dbConfig,
	}
}

func GetConfig() *Config{
	if configurations == nil{
		loadConfig()
	}
	return configurations
}