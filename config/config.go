package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type Config struct{
	Version  		 string
	ServiceName  string
	HttpPort 		 int
	JwtSecretKey string
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

	configurations = &Config{
		Version:       version,
		ServiceName:  serviceName,
		HttpPort:     int(poet),
		JwtSecretKey: jwtSecretKey,
	}
}

func GetConfig() *Config{
	if configurations == nil{
		loadConfig()
	}
	return configurations
}