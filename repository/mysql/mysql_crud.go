package mysql

import (
	"context"
	"database/sql"

	"example.com/go-demo-1/domain"
 
)

type mysqlCrudRepository struct{
	mysqlDb *sql.DB
}


func NewMySqlCrudRepository(db *sql.DB) domain.CrudRepository{

	return &mysqlCrudRepository{
		mysqlDb:db,
	}
}
 

func (m mysqlCrudRepository)get(ctx context.Context , query string,args ...interface{})(user domain.User , err error){
	stmt ,err := m.mysqlDb.PrepareContext(ctx , query)

	if err != nil{

	}

	row := stmt.QueryRowContext(ctx , args...)

	user = domain.User{}

	err = row.Scan(
		&user.ID,
		&user.Phone,
		&user.Name,
		&user.Password,
	)

	return
}

 


func (m *mysqlCrudRepository)Create(ctx context.Context , user domain.User  )(uuser domain.User ,e error){

	
	query := "INSERT INTO `teacher` (`phone`, `name`, `password`) VALUES (?, ?, ?)"
	insertResult , err := m.mysqlDb.ExecContext(context.Background() , query , user.Phone , user.Name , user.Password)

	if err != nil {

	}
	selectQuery := `SELECT id, phone, name, password FROM users WHERE id=?`

	id , err:=insertResult.LastInsertId()
	user , err = m.get(ctx , selectQuery , id)
 

	user , e =m.get(ctx , query , id)
 

	return user , err
}


func (m *mysqlCrudRepository)Remove(user domain.User)(e  error){
	_, err := m.mysqlDb.Exec("DELETE FROM users WHERE id = ?", user.ID)
 
	return err
}


func (m *mysqlCrudRepository) Read(ctx context.Context , id int64)(u domain.User , e error){
	query := `SELECT id, name, created_at, updated_at FROM author WHERE id=?`
 

	u , e =m.get(ctx , query , id)

	return
}


func (m *mysqlCrudRepository)Update(ctx context.Context , user domain.User)(re int64 ,e error){
	result, err := m.mysqlDb.Exec("update product set phone = ?, name = ?, password = ? WHERE id = ?", user.Phone, user.Name, user.Password , user.ID)
 
 
	re , e  =result.LastInsertId() 
	
 
	return re, err
}