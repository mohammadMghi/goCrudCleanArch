package middleware

import "example.com/go-demo-1/domain"



type ResponseError struct{
	Message string `json:message`
}


type CrudHandler struct{
	CrudUseCase domain.CrudUsecase
}

func NewCrudHalder(){

}