package domain

import "context"

type CrudUsecase interface{
	Create(ctx context.Context  )( u Users,e error)
	Remove(ctx context.Context  ) ( u Users ,e error)
	Read(ctx context.Context  )( u Users,e error)
	Update(ctx context.Context )( u Users ,e error)
}

