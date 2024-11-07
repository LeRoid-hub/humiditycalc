package main

import (
	"fmt"
	"os"

	"github.com/LeRoid-hub/humiditycalc/configs"
	"github.com/LeRoid-hub/humiditycalc/server"
)

func main() {
	env := configs.Load()
	
	switch env["MODE"] {
		case "both":
			server.Run(env)
		case "weather":
			server.Run(env)
			// weather.Run()
		case "calc":
			server.Run(env)
			// calc.Run()
		default:
			server.Run(env)
	}
}

