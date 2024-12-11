# humiditycalc

This service allows you to compare the absolute humidity from indoor and outdoor.
  
## Setup 
The setup is aimed for docker.

### DOCKER COMPOSE EXAMPLE

```
services:
    humiditycalc:
        restart: always
        container_name: humiditycalc
        image: ghcr.io/leroid-hub/humiditycalc:latest
        environment:
          OPENWEATHERMAP_API_KEY: "YOURKEY"
          LATITUDE: "40.07"
          LONGITUDE: "9.8"
          MODE: "BOTH"
```

### ENVIROMENT VARS 
MODES: CALC, WEATHER, BOTH  
BOTH is standard  
  
OPENWEATHERMAP_API_KEY is required in WEATHER and BOTH mode  
  
DEFAULT:  
MODE=CALC  
  
Standalone:  
MODE=BOTH  
OPENWEATHERMAP_API_KEY=yourkey  
LATITUDE=  
LONGITUDE=  
  
WEATHER:  
MODE=WEATHER  
OPENWEATHERMAP_API_KEY=yourkey  
DB_HOST=  
DB_PORT=  
DB_NAME=  
DB_PASSWORD=  
DB=  
  
OPTINAL:  
PORT=80  