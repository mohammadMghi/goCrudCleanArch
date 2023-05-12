package domain

import "context"

type crud struct{

}

type CrudRepository interface{
	Create(ctx context.Context  )
	Remove(ctx context.Context)
	Read(ctx context.Context)
	Update(ctx context.Context)
}