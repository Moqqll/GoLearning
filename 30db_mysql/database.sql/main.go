package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //匿名导入，进行初始化init()
)

type user struct {
	id   int
	age  int
	name string
}

var (
	dbConn *sql.DB
)

func initDB() (err error) {
	//DSN：data Sourse Name
	dsn := "wancheng:wancheng@tcp(127.0.0.1:3306)/sql_test"
	// 不会校验账号密码是否正确
	// 注意：这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量
	dbConn, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("initDB: Database open failed, err:", err)
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = dbConn.Ping()
	if err != nil {
		fmt.Println("initDB: Database doesn't exist.")
		return err
	}
	fmt.Println("initDB: database connect success.")
	return nil
}

//MySQL增删改查
func queryRowDemo() {
	sqlStr := "select id, name, age from users where id=?"
	var u user
	err := dbConn.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age) //sacn函数带有close。
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

func queryMultiRowDemo() {
	sqlStr := "select id, name, age from users where id > ?"
	rows, err := dbConn.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close() //非常重要：关闭rows释放持有的数据库连接

	//循环读取
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

func insertRowDemo() {
	sqlStr := "insert into users(name, age) values (?,?)"
	ret, err := dbConn.Exec(sqlStr, "雅婷宝贝仙女", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	var theID int64
	theID, err = ret.LastInsertId() //新插入的数据的in
	if err != nil {
		fmt.Printf("get lastinsert ID failed,err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d\n", theID)
}

func updateRowDemo() {
	sqlStr := "update users set age = ? where id = ?"
	ret, err := dbConn.Exec(sqlStr, 39, 2)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	var n int64
	n, err = ret.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows: %d\n", n)
}

func deleteRowDemo() {
	sqlStr := "delete from users where id = ?"
	ret, err := dbConn.Exec(sqlStr, 2)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	var n int64
	n, err = ret.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows: %d\n", n)
}

//MySQL预处理
func prepareQueryDemo() {
	sqlStr := "select id, name, age from users where id >?"
	stmtQuery, err := dbConn.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmtQuery.Close()
	rows, err := stmtQuery.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("sacn failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

func prepareInsertDemo() {
	sqlStr := "insert into users(name,age) values(?,?)"
	stmtIn, err := dbConn.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmtIn.Close()

	//一次预处理，多次插入，提升性能
	_, err = stmtIn.Exec("彩虹仙女", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmtIn.Exec("沙河娜扎", 28)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	fmt.Println("insert success.")
}

//MySQL事务
func transactionDemo() {
	tx, err := dbConn.Begin() //开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin transaction failed, err:%v\n", err)
		return
	}

	// // SQL预处理
	// sqlStr := "update users set age = ? where id = ?"
	// stmtIn, err := dbConn.Prepare(sqlStr)
	// if err != nil {
	// 	fmt.Printf("prepare failed, err:%v\n", err)
	// 	return
	// }
	// defer stmtIn.Close()

	//使用事务预处理
	sqlStr := "update users set age = ? where id = ?"
	stmtIn, err := tx.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmtIn.Close()

	//一项更新
	// sqlStr1 := "update users set age = 18 where id = ?"
	ret1, err := stmtIn.Exec(0, 1)
	if err != nil {
		tx.Rollback() //回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	//二项更新
	// sqlStr2 := "update users set age = 16 where id = ?"
	ret2, err := stmtIn.Exec(18, 3)
	if err != nil {
		tx.Rollback() //回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec ret2.RowsAffected() failed, err:%v\n", err)
		return
	}

	//事务提交处理
	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit()
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦！")
		return
	}
	fmt.Println("exec trans success!")
}

func main() {
	//init DB
	err := initDB()
	if err != nil {
		fmt.Println("init db failed,err:", err)
		return
	}
	fmt.Println("main.go: init db success")
	fmt.Println()

	defer fmt.Println("main.go: database connect off")
	defer fmt.Println()
	defer dbConn.Close() //注意这行代码要写在上面err判断的下面

	//prepareInsert
	// prepareInsertDemo()

	//transaction
	transactionDemo()

	//queryRow
	// queryRowDemo()

	//queryMultiRow
	// queryMultiRowDemo()

	//insertRow
	// insertRowDemo()

	//updateRow
	// updateRowDemo()

	//deleteRowD
	// deleteRowDemo()

	//prepareQuery
	// prepareQueryDemo()
}
