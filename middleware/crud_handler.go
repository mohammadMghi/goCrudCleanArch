package middleware

import (
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
	
	e.GET("/remove" ,handler.remove)
	e.GET("/update" , handler.update)
	e.GET("/read" ,handler.read )
	e.GET("/create" , handler.create)
}

func (cr *CrudHandler)update( u []User ,e error)  {
	var test:=c.QueryParam("test")
	var ctx = c.Request().Context()
	e := cr.CrudUseCase.Update(ctx) 
	if  e != nil{

	}
}

func (cr CrudHandler)remove( u []User ,e error)  {
	user ,e :=cr.CrudUseCase.Remove()
	if e != nil{

	}

	c.Response().Handler().Set(`X-Crusor`)
}
func (cr CrudHandler)read( u []User ,e error) {

}

func (cr CrudHandler)create( u []User ,e error) {

}