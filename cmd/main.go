package main

import (
	"fmt"
	"net/http"
	"mls/config"
	"mls/controller"
	"mls/storage"
)

func main() {
	cfg := config.Load()
	store, err := storage.New(cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.DB.Close()

	con := controller.NewController(store)

	http.HandleFunc("/Branch",con.Branch)
	http.HandleFunc("/Teacher",con.Teacher)


	fmt.Println("programm is running on localhost:8008...")
	http.ListenAndServe(":8008", nil)

}