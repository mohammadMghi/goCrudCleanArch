package domain

import "context"

type crud struct{

}

type CrudRepository interface{
    Create(ctx context.Context , user User  )(ruser User ,e error)
	Remove(user User)(  error)
	Read(ctx context.Context , id int64)
	Update(ctx context.Context , user User)(error)
}