package union

import (
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"dtstack.com/dtstack/easymatrix/matrix/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func getConn(){
	user := "root"
	password := "dtstack"
	host:= "172.16.10.37"
	port := 3306
	dbname := "dtagent_test"
	log.ConfigureLogger("/tmp/matrix",100,3,1)
	db,_ := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local&parseTime=true", user, password, host, port, dbname))
	model.MYSQLDB = db
	err := Build()
	if err != nil{
		fmt.Printf("err: %v \n",err)
	}
}
//func Test1(t *testing.T) {
//	getConn()
//	tbscs,err := UnionT4T7.Select(58,"dtstack-system")
//	if err != nil{
//		fmt.Println("err",err.Error())
//		return
//	}
//	for _, sc := range tbscs{
//		fmt.Printf("sc %+v \n",sc)
//	}
//}
