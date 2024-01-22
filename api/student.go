package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zRich/go-backend/internal/db"
	"github.com/zRich/go-backend/internal/db/models"
	"github.com/zRich/go-backend/internal/server"
)

type GetStudentsEndpoint struct {
}

func (e *GetStudentsEndpoint) Method() string {
	return "GET"
}

func (e *GetStudentsEndpoint) Path() string {
	return "/api/students"
}

func (e *GetStudentsEndpoint) LoginVerify() bool {
	return false
}

func (e GetStudentsEndpoint) Handler(ctx *gin.Context) {
	var students []models.Student
	var response server.HttpResonpose
	_end, _ := strconv.Atoi(ctx.DefaultQuery("_end", "10"))
	_start, _ := strconv.Atoi(ctx.DefaultQuery("_start", "1"))

	if err := db.DB.Limit(_end - _start).Offset(_start).Find(&students).Error; err != nil {

		response = server.WrapResponse(http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	// response = server.WrapResponse(http.StatusOK, "success", courses)
	ctx.JSON(http.StatusOK, students)
}

// CreateStudentEndpoint is a struct for creating a student.
type CreateStudentEndpoint struct {
}

func (e *CreateStudentEndpoint) Method() string {
	return "POST"
}

func (e *CreateStudentEndpoint) Path() string {
	return "/api/students"
}

func (e *CreateStudentEndpoint) LoginVerify() bool {
	return false
}

func (e CreateStudentEndpoint) Handler(ctx *gin.Context) {
	var student models.Student
	var response server.HttpResonpose
	if err := ctx.ShouldBindJSON(&student); err != nil {
		response = server.WrapResponse(http.StatusBadRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if err := db.DB.Create(&student).Error; err != nil {
		response = server.WrapResponse(http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	ctx.JSON(http.StatusOK, student)
}

// UpdateStudentByStudentNoEndpoint is a struct for updating a student.
type UpdateStudentByStudentNoEndpoint struct {
}

func (e *UpdateStudentByStudentNoEndpoint) Method() string {
	return "PUT"
}

func (e *UpdateStudentByStudentNoEndpoint) Path() string {
	return "/api/students/:studentNo"
}

func (e *UpdateStudentByStudentNoEndpoint) LoginVerify() bool {
	return false
}

func (e UpdateStudentByStudentNoEndpoint) Handler(ctx *gin.Context) {
	var student models.Student
	var response server.HttpResonpose
	studentNo := ctx.Param("studentNo")
	if err := db.DB.Where("student_no = ?", studentNo).First(&student).Error; err != nil {
		response = server.WrapResponse(http.StatusBadRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if err := ctx.ShouldBindJSON(&student); err != nil {
		response = server.WrapResponse(http.StatusBadRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if err := db.DB.Save(&student).Error; err != nil {
		response = server.WrapResponse(http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	ctx.JSON(http.StatusOK, student)
}

// DeleteStudentEndpoint is a struct for deleting a student.

type DeleteStudentEndpoint struct {
}

func (e *DeleteStudentEndpoint) Method() string {
	return "DELETE"
}

func (e *DeleteStudentEndpoint) Path() string {
	return "/api/students/:studentNo"
}

func (e *DeleteStudentEndpoint) LoginVerify() bool {
	return false
}

func (e DeleteStudentEndpoint) Handler(ctx *gin.Context) {
	var student models.Student
	var response server.HttpResonpose
	studentNo := ctx.Param("studentNo")
	if err := db.DB.Where("student_no = ?", studentNo).First(&student).Error; err != nil {
		response = server.WrapResponse(http.StatusBadRequest, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if err := db.DB.Delete(&student).Error; err != nil {
		response = server.WrapResponse(http.StatusInternalServerError, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	ctx.JSON(http.StatusOK, student)
}
