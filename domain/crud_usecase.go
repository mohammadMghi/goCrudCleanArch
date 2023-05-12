package domain

import "context"

type crudUsecase interface{
	Create(ctx context.Context  )
	Remove(ctx context.Context)
	Read(ctx context.Context)
	Update(ctx context.Context)
}

