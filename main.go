package main

import (
	"github.com/Radical-Egg/Lol-Prediction-Backend-API/database"
)

func main() {
	db := database.NewDataBaseConnection()

	defer database.Close(db.Client, db.Ctx, db.Cancel)

	database.Ping(db.Client, db.Ctx)
}
