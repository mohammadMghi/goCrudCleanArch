package usecase

import (
	"context"

	"example.com/go-demo-1/domain"
)
type crudUseCase struct{
	crudRepo domain.CrudRepository
}


func (c crudUseCase)NewCrudUsecase (curdRepo domain.CrudRepository) domain.CrudRepository{
	return &crudUseCase{
		crudRepo: curdRepo,
	}
}


func (m *crudUseCase)Create(ctx context.Context , user domain.User  )(ruser domain.User ,e error){
	
}
func (m *crudUseCase)Remove(user domain.User)(  error){

}
func (m *crudUseCase) Read(ctx context.Context , id int64){

}
func (m *crudUseCase)Update(ctx context.Context , user domain.User)(error){

}