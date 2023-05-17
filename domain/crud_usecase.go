package domain

import "context"

type CrudUsecase interface{
    Create(ctx context.Context , user User  )(uuser User ,e error)
	Remove(user User)( e  error)
	Read(ctx context.Context , id int64)(user User ,e error )
	Update(ctx context.Context , user User)(re int64 ,e error)
}

