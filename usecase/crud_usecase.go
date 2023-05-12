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


func (m *crudUseCase)Create(ctx context.Context  ){
	
}
func (m *crudUseCase)Remove(ctx context.Context){

}
func (m *crudUseCase) Read(ctx context.Context){

}
func (m *crudUseCase)Update(ctx context.Context){

}