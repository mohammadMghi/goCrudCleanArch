package domain

import "context"

type CrudUsecase interface{
	Create(ctx context.Context  ,user []User  )( u []User ,e error)
	Remove(ctx context.Context,user []User ) ( u []User ,e error)
	Read(ctx context.Context,user []User )( u []User ,e error)
	Update(ctx context.Context,user []User )( u []User ,e error)
}

