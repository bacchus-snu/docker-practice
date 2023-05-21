package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, u.Username)
			log.Printf("Served a request!\n");
		},
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	log.Println("Listening on port", port)
	http.ListenAndServe(":"+port, nil)
}
