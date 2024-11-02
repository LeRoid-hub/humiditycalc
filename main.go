package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/LeRoid-hub/humiditycalc/server"
	"github.com/joho/godotenv"
)

func main() {
	env := loadEnv()
	if val, ok := env["MODE"]; ok {
		if strings.ToLower(val) == "both" {
			checkEnv(env)
			server.Run()
		} else if strings.ToLower(val) == "weather" {
			checkEnv(env)
			// weather.Run()
		} else if strings.ToLower(val) == "calc" {
			// calc.Run()
			server.Run()
		}
	}
}

func loadEnv() map[string]string {
	var env map[string]string
	env, err := godotenv.Read()
	if err != nil {
		fmt.Println("Error loading .env file: ", err)
		os.Exit(1)
	}
	if len(env) == 0 {
		fmt.Println(".env file is empty")
		os.Exit(1)
	}
	return env
}

func checkEnv(env map[string]string) {
	// Is there an API key for openweathermap?
	if val, ok := env["OPENWEATHERMAP_API_KEY"]; ok {
		fmt.Println(val)
		if val == "" {
			print("OPENWEATHERMAP_API_KEY is not set")
			os.Exit(1)
		}
	} else {
		fmt.Println("OPENWEATHERMAP_API_KEY is not set")
		os.Exit(1)
	}
}
