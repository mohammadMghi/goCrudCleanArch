package middleware

import (
	"encoding/json"
	"net/http"

	"example.com/go-demo-1/domain"
	"github.com/labstack/echo/v4"
)



type ResponseError struct{
	Message string `json:message`
}


type CrudHandler struct{
	CrudUseCase domain.CrudUsecase
}

func NewCrudHalder(e *echo.Echo  , crudUsercase domain.CrudUsecase){
	handler := &CrudHandler{
		CrudUseCase: crudUsercase,
	}
	
	e.DELETE("/remove" ,handler.Remove)
	e.POST("/update" , handler.Update)
	e.GET("/read" ,handler.Read )
	e.PUT("/create" , handler.Create)
}

func (cr *CrudHandler)Update(ctx echo.Context )( e error) {
	
 
	var users domain.Users  
	
	json.NewDecoder(ctx.Request().Body).Decode(&users)

	var ctxa = ctx.Request().Context()

	u , e  := cr.CrudUseCase.Update(ctxa  , users) 
	
	if  e != nil{
		return
	}

	return ctx.JSON(http.StatusOK , u)
}

func (cr CrudHandler)Remove(ctx echo.Context ) (  e error) {

	var context = ctx.Request().Context()

	 u,e :=cr.CrudUseCase.Remove(context)

	if e != nil{
		return ctx.JSON(http.StatusNotFound , "Error")
	}

	return ctx.JSON(http.StatusOK , u)

}


func (cr CrudHandler)Read(ctx echo.Context )( e error) {
 

	var c = ctx.Request().Context()

	u , e  := cr.CrudUseCase.Read(c) 

	if  e != nil{
		return ctx.JSON(http.StatusNotFound , "Error")
	}

	return ctx.JSON(http.StatusOK , u)
}



func (cr CrudHandler)Create(ctx echo.Context )( error) {

 
	var users domain.Users  
	
	json.NewDecoder(ctx.Request().Body).Decode(&users)

	var ctxa = ctx.Request().Context()

	u , e  := cr.CrudUseCase.Update(ctxa  , users) 

	if  e != nil{
		return ctx.JSON(http.StatusOK , "Error")
	}

	return ctx.JSON(http.StatusOK , u)
}