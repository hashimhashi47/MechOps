package db

import (
	"log"
	"github.com/joho/godotenv"
)

//Connection
func Connection(){

	err:=godotenv.Load()
	if err!=nil{
		  log.Fatal("Error loading .env file")
	}


}