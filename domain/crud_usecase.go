package domain

import "context"

type CrudUsecase interface{
	Create(ctx context.Context , user Users )(users Users ,   e error)


	Remove(ctx context.Context  ) (users Users, e error)


	Read(ctx context.Context  )( users Users , e error)

	
	Update(ctx context.Context , u Users )(users Users,  e error)
}

