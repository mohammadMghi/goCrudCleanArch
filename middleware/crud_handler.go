package middleware

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/go-demo-1/domain"
	"github.com/labstack/echo/v4"
)



type ResponseError struct{
	Message string `json:message`
}


type CrudHandler struct{
	CrudUseCase domain.CrudUsecase
}

func NewCrudHalder(e *echo.Echo  , crudUsercase domain.CrudUsecase)CrudHandler{
	handler := &CrudHandler{
		CrudUseCase: crudUsercase,
	}
	
	e.DELETE("/remove" ,handler.Remove)
	e.POST("/update" , handler.Update)
	e.GET("/read:id" ,handler.Read )
	e.PUT("/create" , handler.Create)
}

func (cr *CrudHandler)Update(ctx echo.Context )( e error) {
	
 
	var user domain.User
	
	json.NewDecoder(ctx.Request().Body).Decode(&user)

	var ctxa = ctx.Request().Context()
	u , e  := cr.CrudUseCase.Update(ctxa  , user) 
	if  e != nil{
		return
	}
	
	return ctx.JSON(http.StatusOK , u)
}

func (cr CrudHandler)Remove(ctx echo.Context ) (  e error) {

 
	var user domain.User  
	
	json.NewDecoder(ctx.Request().Body).Decode(&user)
	e  =cr.CrudUseCase.Remove( user)

	if e != nil{
		return ctx.JSON(http.StatusNotFound , "Error")
	}

	return ctx.JSON(http.StatusOK , "success")

}


func (cr CrudHandler)Read(ctx echo.Context )( e error) {
 

	var c = ctx.Request().Context()
	id, e := strconv.Atoi(ctx.Param("id"))
	

	u , e  := cr.CrudUseCase.Read(c,int64(id)) 

	if  e != nil{
		return ctx.JSON(http.StatusNotFound , "Error")
	}

	return ctx.JSON(http.StatusOK , u)
}



func (cr CrudHandler)Create(ctx echo.Context )( error) {

 
	var user domain.User 
	
	json.NewDecoder(ctx.Request().Body).Decode(&user)

	var ctxa = ctx.Request().Context()

	u , e  := cr.CrudUseCase.Update(ctxa  , user) 

	if  e != nil{
		return ctx.JSON(http.StatusOK , "Error")
	}

	return ctx.JSON(http.StatusOK , u)
}