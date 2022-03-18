package main

import (
	"fmt"

	"github.com/henilthakor/goapi/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Welcome to my first API.")

	var controller controllers.StudentController

	//Defining new echo server
	e := echo.New()

	//Defining Paths

	e.POST("/newstudent", controller.AddNewStudentDetail)
	e.GET("/studentdetail/:id", controller.GetStudentDetail)
	e.PUT("/updatestudent", controller.UpdateStudentDetail)
	e.DELETE("/deletestudent/:id", controller.DeleteStudentDetail)

	//Starting the server on Port :5000
	e.Logger.Fatal(e.Start(":5000"))

}
