package main

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //匿名导入，进初始化
	"github.com/jmoiron/sqlx"
)

type user struct {
	//结构体字段首字母得大写
	//否则err = dbConn.Get(&u, sqlStr, 3)会报错
	ID   int
	Age  int
	Name string
}

var (
	dbConn *sqlx.DB
	err    error
)

func initDB() (err error) {
	dsn := "wancheng:wancheng@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	//也可以使用MustConnect，连接不成功就panic
	dbConn, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	dbConn.SetMaxOpenConns(20)
	dbConn.SetMaxIdleConns(10)
	return
}

func queryRowDemo() {
	sqlStr := "select id, name, age from users where id = ?"
	var u user
	err = dbConn.Get(&u, sqlStr, 3)
	if err != nil {
		fmt.Printf("db get failed, err: %v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

func queryMultiRowDemo() {
	sqlStr := "select id, name, age from users where id > ?"
	var users []user
	err = dbConn.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err: %v\n", err)
		return
	}
	fmt.Printf("users: %#v\n", users)
}

func insertRowDemo() {
	sqlStr := "insert into users(name, age) values(?,?)"
	ret, err := dbConn.Exec(sqlStr, "沙河小王子", 19)
	if err != nil {
		fmt.Printf("insert failed, err: %v\n", err)
		return
	}
	theID, err := ret.LastInsertId() //新插入数据的id
	if err != nil {
		fmt.Printf("get LastInsertId failed, err: %v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

func insertUserDemo() {
	sqlStr := "insert into users (name,age) values (:name,:age)"
	_, err := dbConn.NamedExec(sqlStr, map[string]interface{}{
		"name": "谭万铖最帅",
		"age":  18,
	})
	if err != nil {
		fmt.Printf("insert user failed, err: %v\n", err)
		return
	}
	return
}

// func namedQuery() {

// }

func transactionDemo2() (err error) {
	tx, err := dbConn.Beginx() //开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err: %v\n", err)
		return err
	}

	//注册  事务提交
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) //re-throw panic after Rollback
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback()
		} else {
			err = tx.Commit() //err is nil, if Commit returns error, update err.
			fmt.Println("commit")
		}
	}()

	//first insert
	sqlStr1 := "Update users set age=20 where id=?"
	ret, err := tx.Exec(sqlStr1, 6)
	if err != nil {
		return err
	}
	n, err := ret.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}

	//second insert
	sqlStr2 := "Update users set age=50 where id=?"
	ret2, err := tx.Exec(sqlStr2, 5)
	if err != nil {
		return err
	}
	n, err = ret2.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	return err
}

func main() {
	err = initDB()
	if err != nil {
		fmt.Printf("db init failed, err:%v\n", err)
		return
	}
	fmt.Println("main.go: db init success!")
	fmt.Println()
	defer fmt.Println("main.go: db connect off")
	defer fmt.Println()
	defer dbConn.Close()

	//transactionDemo2
	transactionDemo2()

	// //insertUserDemo
	// insertUserDemo()

	// //插入数据
	// insertRowDemo()

	// //多行查询
	// queryMultiRowDemo()

	// //单行查询
	// queryRowDemo()
}
