package main

import (
	"log"
	"server/db"
	"server/modules/app"
	"server/modules/shared"
	"server/modules/user"
)

func main() {
	conn, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("could not initialize database connection %s", err)
	}

	sharedSvc := shared.NewService()

	userRep := user.NewRepository(conn.GetDB())
	userSvc := user.NewService(userRep, sharedSvc)
	userHandler := user.NewHandler(userSvc)

	app.InitRouter(userHandler)
}
