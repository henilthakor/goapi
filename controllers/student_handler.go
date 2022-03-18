package controllers

import (
	"net/http"

	"github.com/henilthakor/goapi/domain"
	"github.com/labstack/echo/v4"
)

type StudentController interface {
	AddNewStudentDetail(echo.Context) error
	GetStudentDetail(echo.Context) error
	UpdateStudentDetail(echo.Context) error
	DeleteStudentDetail(echo.Context) error
}

//----------------------------------------------------------//
//HAD TO CREATE NEW STRUCT BECAUSE IT OTHERWISE SHOWED ERROR//
//----------------------------------------------------------//

type studentService struct {
	domain.Student
}

func (*studentService) AddNewStudentDetail(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hey there")
}

func (*studentService) GetStudentDetail(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hey there")
}

func (*studentService) UpdateStudentDetail(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hey there")
}

func (*studentService) DeleteStudentDetail(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hey there")
}
