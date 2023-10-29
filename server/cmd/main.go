package main

import (
	"log"
	"server"
	"server/db"
	"server/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	// userRep := user.NewRepository(dbConn.GetDB())
	// userSev := user.NewService(userRep)
	// userHandler := user.NewHandler(userSev)

	userHandler, _ := server.InitializeHandler(dbConn.GetDB())

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")
}
