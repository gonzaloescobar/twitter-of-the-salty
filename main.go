package main

import (
	"log"
	"github.com/gonzaloescobar/twitter-of-the-salty/handlers"
	"github.com/gonzaloescobar/twitter-of-the-salty/db"
)


func main()  {
	if db.CheckConnection() == 0{
		log.Fatal("Connection refused")
		return
	}
	handlers.Handlers()
}