package mysql

import (
	"context"
	"database/sql"

	"example.com/go-demo-1/domain"
	"github.com/go-ginger/dl/query"
)

type MysqlCrudRepository struct{
	mysqlDb *sql.DB
}


func NewMySqlCrudRepository(db *sql.DB) domain.CrudRepository{

	return &MysqlCrudRepository{
		mysqlDb:db,
	}
}
 

func (m MysqlCrudRepository)get(ctx context.Context , query string,args ...interface{})(user domain.User , err error){
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

func (m MysqlCrudRepository)insert(ctx context.Context , user domain.User)(u domain.User ,err error){
	
	query := "INSERT INTO `teacher` (`phone`, `name`, `password`) VALUES (?, ?, ?)"
	insertResult , err := m.mysqlDb.ExecContext(context.Background() , query , user.Phone , user.Name , user.Password)

	if err != nil {

	}

	query = `SELECT id, phone, name, password FROM users WHERE id=?`

	id , err:=insertResult.LastInsertId()
	user , err = m.get(ctx , query , id)

	if err != nil {

	}

	return user , err
}

func (m *MysqlCrudRepository)Create(ctx context.Context , id int64  )(domain.User , error){

}
func (m *MysqlCrudRepository)Remove(ctx context.Context)(domain.User , error){

}
func (m *MysqlCrudRepository) Read(ctx context.Context , id int64)(domain.User , error){
	query := "" 
	return m.get(ctx , query , id)
}
func (m *MysqlCrudRepository)Update(ctx context.Context)(domain.User , error){

}