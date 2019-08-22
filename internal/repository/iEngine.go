package repository

type IEngine interface {
	//Exec(sqlorArgs ...interface{}) (sql.Result, error)
	//Get(bean interface{}) (bool, error)
	//Find(beans interface{}, condiBeans ...interface{}) error
	//SQL(query interface{}, args ...interface{}) *xorm.Session
	Insert(beans ...interface{}) (int64, error)
	Exist(bean ...interface{}) (bool, error)
	//Update(bean interface{}, condiBeans ...interface{}) (int64, error)
}
