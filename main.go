package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"example.com/go-demo-1/middleware"
	"example.com/go-demo-1/repository/mysql"
	"example.com/go-demo-1/usecase"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bunrouter"
)


func main(){
 
    router := bunrouter.New()

    router.GET("/", func(w http.ResponseWriter, req bunrouter.Request) error {

        w.Write([]byte("index"))
        return nil
    })

	dbHost :=  ""
	dbPort :=  ""
	dbUser :=  ""
	dbPass :=  ""
	dbName :=  ""
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

    e := echo.New()
    middL := middleware.InitMiddleware()
    e.Use(middL.CORS)
	if err != nil {
		log.Fatal(err)
	}
	mysqlRepo  := mysql.NewMySqlCrudRepository(dbConn)
    
	crdUseCase := usecase.NewCrudUsecase(mysqlRepo)

	middleware.NewCrudHalder(e,crdUseCase)

    log.Println("listening on http://localhost:8080")
    log.Println(http.ListenAndServe(":8080", router))
}