package usecase

import (
	"context"

	"example.com/go-demo-1/domain"
)
type crudUseCase struct{
	crudRepo domain.CrudRepository
}


func NewCrudUsecase (curdRepo domain.CrudRepository) domain.CrudRepository{
	return &crudUseCase{
		crudRepo: curdRepo,
	}
}


func (m *crudUseCase)Create(ctx context.Context , user domain.User  )(ruser domain.User ,e error){
	ruser,e = m.crudRepo.Create(ctx , user)
	return 
}
func (m *crudUseCase)Remove(user domain.User)(e  error){
	 e = m.crudRepo.Remove( user)
	return 
}
func (m *crudUseCase) Read(ctx context.Context , id int64)(user domain.User ,e error ){
	user,e = m.crudRepo.Read(ctx,  id)
	return
}
func (m *crudUseCase)Update(ctx context.Context , user domain.User)(re int64 ,e error){
	re,e = m.crudRepo.Update(ctx,  user)
	return 
}