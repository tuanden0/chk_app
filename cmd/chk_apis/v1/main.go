package main

import (
	v1 "chk/internal/chk_apis/v1"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {

	db, err := v1.NewDB()
	if err != nil {
		log.Println(err)
	}

	repo := v1.NewRepo(db)

	handler := v1.NewHandler(repo)

	v1.NewApp(handler).RunServer()
}
