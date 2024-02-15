package main

import (
	"log"
	"server/db"
	"server/internal/user"
	"server/internal/ws"
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
	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start(":8080")

}
