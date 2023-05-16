package main

import (
	"log"
	"net/http"

	"example.com/go-demo-1/middleware"
	"github.com/uptrace/bunrouter"
)


func main(){
 
    router := bunrouter.New()

    router.GET("/", func(w http.ResponseWriter, req bunrouter.Request) error {

        w.Write([]byte("index"))
        return nil
    })

    middL := middleware.ne

    log.Println("listening on http://localhost:8080")
    log.Println(http.ListenAndServe(":8080", router))
}