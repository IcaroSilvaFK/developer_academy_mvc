package main

import (
	"log"
	"os"
	"path"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/routes"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	e := gin.Default()

	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}

	expPath := path.Dir(ex)

	pages := path.Join(expPath, "public", "pages")
	e.Static("/assets", "./public/assets")

	paths, err := utils.ReadTemplatesFiles(pages)

	if err != nil {
		log.Fatal(err)
	}

	e.LoadHTMLFiles(paths...)

	routes.NewWebRoutes(e)

	log.Fatal(e.Run())
}
