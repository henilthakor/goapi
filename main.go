package main

import (
	"fmt"

	"github.com/henilthakor/goapi/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Welcome to my first API.")

	//Defining new echo server
	e := echo.New()
	fmt.Println("Echo Created")
	//Defining Paths

	e.GET("/studentdetail/:id", controllers.GetStudentDetail)

	e.POST("/addnewstudent", controllers.AddNewStudentDetail)

	e.PUT("/updatestudent", controllers.UpdateStudentDetail)

	e.DELETE("/deletestudent/:id", controllers.DeleteStudentDetail)

	fmt.Println("Paths Defined")
	//Starting the server on Port :5000
	err := e.Start(":5000")
	if err != nil {
		panic("Error occured in starting the server")
	}

}
