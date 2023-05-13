package mysql

import (
	"context"
	"database/sql"

	"example.com/go-demo-1/domain"
)

type MysqlCrudRepository struct{
	mysqlDb *sql.DB
}


func NewMySqlCrudRepository(db *sql.DB) domain.CrudRepository{

	return &MysqlCrudRepository{
		mysqlDb:db,
	}
}
 

func (m *MysqlCrudRepository)Create(ctx context.Context  ){
	
}
func (m *MysqlCrudRepository)Remove(ctx context.Context){

}
func (m *MysqlCrudRepository) Read(ctx context.Context){

}
func (m *MysqlCrudRepository)Update(ctx context.Context){

}