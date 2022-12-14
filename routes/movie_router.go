package routes

import (
	"github.com/cnugroho11/movie_api/controllers"
	"github.com/gin-gonic/gin"
)

type MovieRouteController struct {
	movieController controllers.MovieController
}

func NewMovieRouteController(movieController controllers.MovieController) MovieRouteController {
	return MovieRouteController{movieController}
}

func (mc *MovieRouteController) MovieRoute(rg *gin.RouterGroup) {
	router := rg.Group("/movie")
	// POST Route
	router.POST("/add", mc.movieController.InsertMovie)

	// GET Route
	router.GET("/:id", mc.movieController.FetchMovie)
	router.GET("/all", mc.movieController.FetchAllMovie)

	// PATCH Route
	router.PATCH("/edit", mc.movieController.UpdateMovie)

	// DELETE Route
	router.DELETE("/delete", mc.movieController.DeleteMovie)
}
