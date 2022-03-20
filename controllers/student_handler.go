package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/henilthakor/goapi/domain"
	"github.com/henilthakor/goapi/services"
	"github.com/labstack/echo/v4"
)

func AddNewStudentDetail(ctx echo.Context) error {
	s := new(domain.Student)

	if err := ctx.Bind(s); err != nil {
		return errors.New("Bad Request")
	}

	services.AddNewStudentDetail(s)

	return ctx.JSON(http.StatusOK, s)
}

func GetStudentDetail(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errors.New("Bad Request")
	}

	s, err := services.GetStudentDetail(id)
	if err != nil {
		return errors.New("Wrong Student Id")
	}

	return ctx.JSON(http.StatusOK, s)
}

func UpdateStudentDetail(ctx echo.Context) error {
	s := new(domain.Student)
	if err := ctx.Bind(s); err != nil {
		return errors.New("Bad Request")
	}

	err := services.UpdateStudentDetail(s)

	if err != nil {
		return errors.New("Wrong Student Id")
	}

	return ctx.JSON(http.StatusOK, s)
}

func DeleteStudentDetail(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errors.New("Bad Request")
	}

	if err := services.DeleteStudentDetail(id); err != nil {
		return errors.New("Wrong Student Id")
	}

	return ctx.NoContent(http.StatusOK)
}
