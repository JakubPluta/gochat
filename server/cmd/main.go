package main

import (
	"log"
	"server/db"
	"server/internal/user"
	"server/router"
)

func main() {

	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	userRepository := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)
	router.InitRouter(userHandler)
	router.Start(":8080")

}
