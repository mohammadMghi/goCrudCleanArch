package middleware

import (
	"encoding/json"
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
	
	e.DELETE("/remove" ,handler.remove)
	e.POST("/update" , handler.update)
	e.GET("/read" ,handler.read )
	e.PUT("/create" , handler.create)
}

func (cr *CrudHandler)Update(ctx echo.Context )( u Users ,e error) {
	
 
	var users domain.Users  
	
	json.NewDecoder(c.Request().Body).Decode(&users)

	var ctx = c.Request().Context()
	esd  , sus := cr.CrudUseCase.Update(ctx echo.Context , users.UsersArray) 
	if  e != nil{

	}
}

func (cr CrudHandler)Remove(ctx context.Context  ) ( u Users ,e error) {
	user ,e :=cr.CrudUseCase.Remove()
	if e != nil{

	}

	c.Response().Handler().Set(`X-Crusor`)
}
func (cr CrudHandler)Read(ctx context.Context  )( u Users,e error) {

}

func (cr CrudHandler)Create(ctx context.Context  )( u Users,e error) {

}