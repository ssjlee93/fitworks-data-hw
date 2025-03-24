package main

import (
	"github.com/ssjlee93/fitworks-data-hw/configs"
	"github.com/ssjlee93/fitworks-data-hw/repositories"
	"log"
	"net/http"
)

func main() {
	// Code
	log.Println("Starting the application...")
	db := configs.GetConnection()
	defer db.Close()

	// repositories
	exerciseRepository := repositories.NewExerciseRepository(db)
	log.Printf("Initializing %T", exerciseRepository)
	//
	//userRepo := repositories.NewUserRepository(*userDao)
	//
	//userController := controllers.NewUserController(*userRepo)
	//
	//http.HandleFunc("/users", userController.ReadAllHandler)
	//http.HandleFunc("/user/", userController.Handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
