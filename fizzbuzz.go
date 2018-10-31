package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{"+Fizz+"}/{"+Buzz+"}/{"+Limit+"}/{"+FizzString+"}/{"+BuzzString+"}", GetFizzBuzz).Methods("GET")
	router.HandleFunc("/int1/{"+Fizz+"}/int2/{"+Buzz+"}/limit/{"+Limit+"}/string1/{"+FizzString+"}/string2/{"+BuzzString+"}", GetFizzBuzz).Methods("GET")
	router.HandleFunc("/", Usage).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(Usage)

	srv := &http.Server{
		Addr:         ":8000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: router,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
			os.Exit(-1)
		}
	}()

	WaitForCtrlC()
	ctx := context.Background()
	go func() {
		srv.Shutdown(ctx)
	}()

	<-ctx.Done()
	os.Exit(0)
}
