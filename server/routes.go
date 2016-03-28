package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

// Config file struct definition
type Config struct {
	PORT string
}

func main() {

	/* Read config.json file */
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		color.Red("Read file error")
		os.Exit(-1)
	}

	var config Config
	json.Unmarshal(file, &config)

	router := gin.Default()

	api := router.Group("/api")

	api.GET("/name/:name", UserControllerGET)

	router.StaticFile("/", "../client/dist/index.html")
	router.Static("/assets", "../client/dist")

	router.Run(":" + config.PORT) // listen and server on 0.0.0.0:8080
}
