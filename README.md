golang语言实现mysql数据库的增删改查

golang语言对于数据库的操作，只是实现了驱动的接口，具体的驱动得自己实现，
本项目基于"github.com/go-sql-driver/mysql" 实现的mysql的增删改查操作


func connect() (*sql.DB, error) {     //加载驱动
	db, err := sql.Open("mysql", "root:123456@/goweb")
	checkErr(err)
	return db, err
}
