package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	"github.com/IcaroSilvaFK/developer_academy_mvc/application/http/middlewares"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/routes"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/services"
	"github.com/IcaroSilvaFK/developer_academy_mvc/application/utils"
	"github.com/IcaroSilvaFK/developer_academy_mvc/cmd/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			DEVELOPER ACADEMY
//	@version		1.0
//	@description	This api was designed to be able to use the DEVELOPER ACADEMY platform externally
//	@termsOfService	http://swagger.io/support

//	@contact.email	icarovsilva1@gmail.com
//	@contact.name	Icaro Vieira
//	@contact.url	http://

//	@license.name	MIT
//	@license.url	https://mit-license.org/

//	@host		localhost:8080
//	@basePath	/api/v1

//	@securityDefinitions.basic	Session

func main() {

	// f, _ := os.Create("trace.out")

	// defer f.Close()

	// trace.Start(f)

	// defer trace.Stop()
	docs.SwaggerInfo.BasePath = "/api/v1"
	mode := os.Getenv(utils.MODE)
	err := godotenv.Load()

	if mode != utils.PRODUCTION_MODE && err != nil {
		log.Println(err)
	}

	e := gin.Default()
	store := cookie.NewStore([]byte(os.Getenv(utils.SESSION_KEY)))

	e.Use(middlewares.Throttle(1000, 20))
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(sessions.Sessions(utils.SESSION_NAME, store))
	e.Use(middlewares.AddCurrentInContextRequest(services.NewSessionService()))
	e.Use(middlewares.Logger())
	e.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

	routes.NewApiRoutes(e)
	routes.NewWebRoutes(e)

	port := os.Getenv(utils.PORT)

	if port == "" {
		port = "8080"
	}

	log.Printf("Server runing at %s", port)
	log.Fatal(e.Run(fmt.Sprintf(":%v", port)))
}
