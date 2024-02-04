package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zRich/go-backend/internal/db"
	"github.com/zRich/go-backend/internal/db/models"
	"github.com/zRich/go-backend/internal/server"
)

type GetTasksEndpoint struct {
}

func (e *GetTasksEndpoint) Method() string {
	return "GET"
}

func (e *GetTasksEndpoint) Path() string {
	return "tasks"
}

func (e *GetTasksEndpoint) LoginVerify() bool {
	return false
}

func (e GetTasksEndpoint) Handler(ctx *gin.Context) {
	var tasks []models.Task
	var response server.HttpResonpose

	_end, _ := strconv.Atoi(ctx.DefaultQuery("_end", "10"))
	_start, _ := strconv.Atoi(ctx.DefaultQuery("_start", "1"))

	studentNo := ctx.Param("studentNo")

	_db := db.DB.Limit(_end - _start).Offset(_start)

	if studentNo != "" {
		_db = _db.Where("student_no = ?", studentNo)
	}

	if err := _db.Find(&tasks).Error; err != nil {
		response = server.WrapResponse(http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	// response = server.WrapResponse(http.StatusOK, "success", courses)
	ctx.JSON(http.StatusOK, tasks)
}

// CreateTaskEndpoint is a struct for creating a task.
type CreateTaskEndpoint struct {
}

func (e *CreateTaskEndpoint) Method() string {
	return "POST"
}

func (e *CreateTaskEndpoint) Path() string {
	return "tasks"
}

func (e *CreateTaskEndpoint) LoginVerify() bool {
	return false
}

func (e CreateTaskEndpoint) Handler(ctx *gin.Context) {
	var task models.Task
	var response server.HttpResonpose
	if err := ctx.ShouldBindJSON(&task); err != nil {
		response = server.WrapResponse(http.StatusBadRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if err := db.DB.Create(&task).Error; err != nil {
		response = server.WrapResponse(http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	ctx.JSON(http.StatusOK, task)
}

// UpdateTaskEndpoint is a struct for updating a task.
type UpdateTaskEndpoint struct {
}

func (e *UpdateTaskEndpoint) Method() string {
	return "PUT"
}

func (e *UpdateTaskEndpoint) Path() string {
	return "tasks/:id"
}

func (e *UpdateTaskEndpoint) LoginVerify() bool {
	return false
}

func (e UpdateTaskEndpoint) Handler(ctx *gin.Context) {
	var task models.Task
	var response server.HttpResonpose
	if err := ctx.ShouldBindJSON(&task); err != nil {
		response = server.WrapResponse(http.StatusBadRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if err := db.DB.Save(&task).Error; err != nil {
		response = server.WrapResponse(http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	ctx.JSON(http.StatusOK, task)
}

// DeleteTaskEndpoint is a struct for deleting a task.
type DeleteTaskEndpoint struct {
}

func (e *DeleteTaskEndpoint) Method() string {
	return "DELETE"
}

func (e *DeleteTaskEndpoint) Path() string {
	return "tasks/:id"
}

func (e *DeleteTaskEndpoint) LoginVerify() bool {
	return false
}

func (e DeleteTaskEndpoint) Handler(ctx *gin.Context) {
	var task models.Task
	var response server.HttpResonpose
	id := ctx.Param("id")
	if err := db.DB.Where("id = ?", id).First(&task).Error; err != nil {
		response = server.WrapResponse(http.StatusBadRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if err := db.DB.Delete(&task).Error; err != nil {
		response = server.WrapResponse(http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	ctx.JSON(http.StatusOK, task)
}
