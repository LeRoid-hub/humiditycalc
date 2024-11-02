package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Load() map[string]string {
	var env map[string]string = make(map[string]string)

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

	mode := os.Getenv("MODE")
	if mode != "" {
		env["MODE"] = os.Getenv("MODE")
	}

	openweathermapAPIKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if openweathermapAPIKey != "" {
		env["OPENWEATHERMAP_API_KEY"] = openweathermapAPIKey
	}

	latitude := os.Getenv("LATITUDE")
	if latitude != "" {
		env["LATITUDE"] = latitude
	}

	longitude := os.Getenv("LONGITUDE")
	if longitude != "" {
		env["LONGITUDE"] = longitude
	}

	port := os.Getenv("PORT")
	if port != "" {
		env["PORT"] = port
	}

	if len(env) == 0 {
		fmt.Println("no environment variables are set")
		os.Exit(1)
	}

	return env
}
