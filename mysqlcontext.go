package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"	
)


type MySqlContext struct{
	currentContext *gorm.DB
	connectionString string
}

func NewMySqlContext(db_conn string) *MySqlContext {
	return &MySqlContext{connectionString: db_conn}
}

// ConnectDB connects to a MySQL database
func (_ctx *MySqlContext) connectDB(db_conn string) (*gorm.DB, error) {
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", db_conn)
	if err != nil {
		return nil, err
	}	
	return db, nil
} 

func (_ctx *MySqlContext) GetContext() (*gorm.DB, error) {
	//type err
	if(_ctx.currentContext == nil){		
		ctx, err := _ctx.connectDB(_ctx.connectionString)
		if err != nil {
			return nil, err
		}
		_ctx.currentContext = ctx
	}
	
	return _ctx.currentContext, nil
}