package configs

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Load() map[string]string {
	var env map[string]string = make(map[string]string)

	validEnv := []string{"MODE", "OPENWEATHERMAP_API_KEY", "LATITUDE", "LONGITUDE", "PORT"}

	envpath := "./.env"

	if _, err := os.Stat(envpath); err == nil {

		dotenv, err := godotenv.Read(envpath)
		if err != nil {
			fmt.Println("Error loading .env file: ", err)
		}

		env = dotenv
	} else {
		fmt.Println("No .env file found", err)
	}

	for _, key := range validEnv {
		tempenv := os.Getenv(key)
		if tempenv != "" {
			env[key] = tempenv
		}
	}

	if len(env) == 0 {
		fmt.Println("no environment variables are set")
		os.Exit(1)
	}
	
	if val, ok := env["MODE"]; ok {
		if val == "" {
			env["MODE"] = "calc"
		}
	} else {
		env["MODE"] = "calc"
	}

	env["MODE"] = strings.ToLower(env["MODE"])
	switch env["MODE"] {
		case "both":
			checkEnvWeather(env)
			checkEnvCalc(env)
		case "weather":
			checkEnvWeather(env)
		default:
			env["MODE"] = "calc"
	}

	return env
}

func checkEnvWeather(env map[string]string) {
	// Is there an API key for openweathermap?
	if val, ok := env["OPENWEATHERMAP_API_KEY"]; ok {
		if val == "" {
			print("OPENWEATHERMAP_API_KEY is not set")
			os.Exit(1)
		}
	} else {
		fmt.Println("OPENWEATHERMAP_API_KEY is not set")
		os.Exit(1)
	}

	// Is there a LATITUDE and LONGITUDE?
	if val, ok := env["LATITUDE"]; ok {
		if val == "" {
			print("LATITUDE is not set")
			os.Exit(1)
		}
	} else {
		fmt.Println("LATITUDE is not set")
		os.Exit(1)
	}
	if val, ok := env["LONGITUDE"]; ok {
		if val == "" {
			print("LONGITUDE is not set")
			os.Exit(1)
		}
	} else {
		fmt.Println("LONGITUDE is not set")
		os.Exit(1)
	}
}

func checkEnvCalc(env map[string]string) {
	// Check for calc variables
}
