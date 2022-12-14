package controllers

import (
	"fmt"
	"github.com/cnugroho11/movie_api/models"
	"github.com/cnugroho11/movie_api/utils"
	"net/http"
	_ "net/http/httputil"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieController struct {
	DB *gorm.DB
}

func NewMovieController(DB *gorm.DB) MovieController {
	return MovieController{DB}
}

// UpdateMovie
// @Summary      Edit movie data
// @Description  Edit movie data in the database
// @Tags         PATCH
// @Accept       json
// @Produce      json
// @Param payload body models.MovieUpdate true "payload movie"
// @Success      200  {object}  models.MovieInput
// @Router       /movie/edit [patch]
func (mc *MovieController) UpdateMovie(ctx *gin.Context) {
	var payload *models.Movie
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "cannot edit movie",
		})
		return
	}

	var movie models.Movie
	getMovie := mc.DB.First(&movie, payload.ID)
	if getMovie.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": fmt.Sprintf("cannot found movie with id %d", payload.ID),
		})
		return
	}

	updateMovie := mc.DB.Model(&movie).Updates(
		models.Movie{
			Title:       payload.Title,
			Description: payload.Description,
			Rating:      payload.Rating,
			Image:       payload.Image,
			UpdatedAt:   time.Now(),
		},
	)
	if updateMovie.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": fmt.Sprintf("cannot edit movie with id %d", payload.ID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "success update data movie",
	})

}

// DeleteMovie
// @Summary      Delete movie data
// @Description  Delete movie data in the database
// @Tags         DELETE
// @Accept       json
// @Produce      json
// @Param payload body models.MovieUpdate true "payload movie"
// @Success      200  {object}  models.MovieInput
// @Router       /movie/delete [delete]
func (mc *MovieController) DeleteMovie(ctx *gin.Context) {
	var payload *models.Movie
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "cannot delete movie",
		})
		return
	}

	var movie models.Movie
	getMovie := mc.DB.First(&movie, payload.ID)
	if getMovie.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": fmt.Sprintf("cannot found movie with id %d", payload.ID),
		})
		return
	}

	deleteMovie := mc.DB.Delete(&movie)
	if deleteMovie.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": fmt.Sprintf("cannot delete movie with id %d", payload.ID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "success delete movie",
	})
}

// FetchAllMovie
// @Summary      Fetch all movies
// @Description  Fetch all movies in database
// @Tags         GET
// @Accept       json
// @Produce      json
// @Param        page    query     string  false  "page"
// @Param        limit    query     string  false  "limit"
// @Param        sort    query     string  false  "sort"
// @Success      200  {object}  models.Movie
// @Router       /movie/all [get]
func (mc *MovieController) FetchAllMovie(ctx *gin.Context) {
	var movies []models.Movie

	pagination := utils.Pagination(ctx)
	offset := (pagination.Page - 1) * pagination.Limit

	getMovies := mc.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&movies)
	if getMovies.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Failed to get movies",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"pagination": gin.H{
			"page":  pagination.Page,
			"limit": pagination.Limit,
			"sort":  pagination.Sort,
		},
		"data": gin.H{
			"movies": movies,
		},
	})

}

// FetchMovie
// @Summary      Fetch movie
// @Description  Fetch movie
// @Tags         GET
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Movie ID"
// @Success      200  {object}  models.Movie
// @Router       /movie/{id} [get]
func (mc *MovieController) FetchMovie(ctx *gin.Context) {
	movieId := ctx.Param("id")

	var movie models.Movie
	getMovie := mc.DB.First(&movie, movieId)
	if getMovie.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "cannot get movie by id " + movieId,
		})
		return
	}

	response := models.Movie{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Rating:      movie.Rating,
		Image:       movie.Image,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"movie": response,
		},
	})
}

// InsertMovie
// @Summary      Insert movie data
// @Description  Insert movie data to the database
// @Tags         POST
// @Accept       json
// @Produce      json
// @Param payload body models.MovieInput true "payload movie"
// @Success      200  {object}  models.MovieInput
// @Router       /movie/add [post]
func (mc *MovieController) InsertMovie(ctx *gin.Context) {
	var payload *models.MovieInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	newMovie := models.Movie{
		Title:       payload.Title,
		Description: payload.Description,
		Rating:      payload.Rating,
		Image:       payload.Image,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	insertDB := mc.DB.Create(&newMovie)
	if insertDB.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "error insert data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Success input data",
	})
}
