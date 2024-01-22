package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zRich/go-backend/internal/db"
	"github.com/zRich/go-backend/internal/db/models"
	"github.com/zRich/go-backend/internal/server"
)

type GetCoursesEndpoint struct {
}

func (e *GetCoursesEndpoint) Method() string {
	return "GET"
}

func (e *GetCoursesEndpoint) Path() string {
	return "/api/courses"
}

func (e *GetCoursesEndpoint) LoginVerify() bool {
	return true
}

func (e GetCoursesEndpoint) Handler(ctx *gin.Context) {
	var courses []models.Course
	var response server.HttpResonpose

	_end, _ := strconv.Atoi(ctx.DefaultQuery("_end", "10"))
	_start, _ := strconv.Atoi(ctx.DefaultQuery("_start", "1"))

	if err := db.DB.Limit(_end - _start).Offset(_start).Find(&courses).Error; err != nil {
		response = server.WrapResponse(http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	// response = server.WrapResponse(http.StatusOK, "success", courses)
	ctx.JSON(http.StatusOK, courses)
}

// CreateCourseEndpoint is a struct for creating a course.
type CreateCourseEndpoint struct {
}

func (e *CreateCourseEndpoint) Method() string {
	return "POST"
}

func (e *CreateCourseEndpoint) Path() string {
	return "/api/courses"
}

func (e *CreateCourseEndpoint) LoginVerify() bool {
	return false
}

func (e CreateCourseEndpoint) Handler(ctx *gin.Context) {
	var course models.Course
	var response server.HttpResonpose
	if err := ctx.ShouldBindJSON(&course); err != nil {
		response = server.WrapResponse(http.StatusBadRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if err := db.DB.Create(&course).Error; err != nil {
		response = server.WrapResponse(http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	// response = server.WrapResponse(http.StatusOK, "success", course)
	ctx.JSON(http.StatusOK, course)
}

//GetCourseByNameEndpoint is a struct for getting a course by name.

type GetCourseByNameEndpoint struct {
}

func (e *GetCourseByNameEndpoint) Method() string {
	return "GET"
}

func (e *GetCourseByNameEndpoint) Path() string {
	return "/api/courses/:name"
}

func (e *GetCourseByNameEndpoint) LoginVerify() bool {
	return false
}

func (e GetCourseByNameEndpoint) Handler(ctx *gin.Context) {
	var course models.Course
	var response server.HttpResonpose
	name := ctx.Param("name")
	if err := db.DB.Where("name = ?", name).First(&course).Error; err != nil {
		response = server.WrapResponse(http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	// response = server.WrapResponse(http.StatusOK, "success", course)
	ctx.JSON(http.StatusOK, course)
}

// UpdateCourseByNameEndpoint is a struct for updating a course.
type UpdateCourseByNameEndpoint struct {
}

func (e *UpdateCourseByNameEndpoint) Method() string {
	return "PUT"
}

func (e *UpdateCourseByNameEndpoint) Path() string {
	return "/api/courses/:name"
}

func (e *UpdateCourseByNameEndpoint) LoginVerify() bool {
	return false
}

func (e UpdateCourseByNameEndpoint) Handler(ctx *gin.Context) {
	var course models.Course
	var response server.HttpResonpose
	name := ctx.Param("name")
	if err := db.DB.Where("name = ?", name).First(&course).Error; err != nil {
		response = server.WrapResponse(http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	if err := ctx.ShouldBindJSON(&course); err != nil {
		response = server.WrapResponse(http.StatusBadRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if err := db.DB.Save(&course).Error; err != nil {
		response = server.WrapResponse(http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	// response = server.WrapResponse(http.StatusOK, "success", course)
	ctx.JSON(http.StatusOK, course)
}

// DeleteCourseByNameEndpoint is a struct for deleting a course.

type DeleteCourseByNameEndpoint struct {
}

func (e *DeleteCourseByNameEndpoint) Method() string {
	return "DELETE"
}

func (e *DeleteCourseByNameEndpoint) Path() string {
	return "/api/courses/:name"
}

func (e *DeleteCourseByNameEndpoint) LoginVerify() bool {
	return false
}

func (e *DeleteCourseByNameEndpoint) Handler(ctx *gin.Context) {
	var course models.Course
	var response server.HttpResonpose
	name := ctx.Param("name")
	if err := db.DB.Where("name = ?", name).First(&course).Error; err != nil {
		response = server.WrapResponse(http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	if err := db.DB.Delete(&course).Error; err != nil {
		response = server.WrapResponse(http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	// response = server.WrapResponse(http.StatusOK, "success", course)
	ctx.JSON(http.StatusOK, course)
}
