package main

import (
	"context"
	"fmt"
	"log"
	"task-tracker/internal/repository"
)

func main(){
	db, err := repository.NewDB()
	if err != nil{
		log.Fatal(err)
	}
	
	fmt.Println("Connected to database!")

	defer db.Close(context.Background())
}