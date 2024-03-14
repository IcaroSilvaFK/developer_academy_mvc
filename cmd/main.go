package main

import (
	"log"
	"os"
	"path"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/routes"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	gin.ForceConsoleColor()

	e := gin.Default()
	store := cookie.NewStore([]byte("secret"))

	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(sessions.Sessions("mysession", store))

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
	routes.NewApiRoutes(e)

	log.Fatal(e.Run())
}
