package main

import (
	"github.com/cnugroho11/movie_api/controllers"
	_ "github.com/cnugroho11/movie_api/docs"
	"github.com/cnugroho11/movie_api/initializers"
	"github.com/cnugroho11/movie_api/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	server *gin.Engine

	MovieController      controllers.MovieController
	MovieRouteController routes.MovieRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load env", err)
	}

	initializers.ConnectDB(&config)

	MovieController = controllers.NewMovieController(initializers.DB)
	MovieRouteController = routes.NewMovieRouteController(MovieController)

	server = gin.Default()
}

// @title           Simple movie API
// @version         1.0
// @description     Simple api app using go, gin, and gorm.
// @termsOfService  http://swagger.io/terms/
// @contact.name   Cahyo Nugroho
// @contact.url    http://www.swagger.io/support
// @contact.email  cnugroho211@gmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8000
// @BasePath  /api
// @securityDefinitions.basic  BasicAuth
func main() {
	server.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "http://localhost:8000/docs/index.html")
	})
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router := server.Group("/api")

	// Movie Route
	MovieRouteController.MovieRoute(router)

	log.Fatal(server.Run(":" + "8000"))
}
