package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/LeRoid-hub/humiditycalc/configs"
	"github.com/LeRoid-hub/humiditycalc/server"
)

func main() {
	env := configs.Load()
	if val, ok := env["MODE"]; ok {
		if strings.ToLower(val) == "both" {
			checkEnv(env)
			server.Run(env)
		} else if strings.ToLower(val) == "weather" {
			checkEnv(env)
			// weather.Run()
		} else if strings.ToLower(val) == "calc" {
			// calc.Run()
			server.Run(env)
		}
	} else {
		// calc.Run()
		server.Run(env)
	}
}

func checkEnv(env map[string]string) {
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
}
