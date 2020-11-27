package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql" //匿名导入，进初始化
	"github.com/jmoiron/sqlx"
)

type user struct {
	//结构体字段首字母得大写
	//否则err = dbConn.Get(&u, sqlStr, 3)会报错
	ID   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

//Testuser ...
type Testuser struct {
	ID   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

var (
	dbConn *sqlx.DB
)

//Value ...
func (u Testuser) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

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
	err := dbConn.Get(&u, sqlStr, 6)
	if err != nil {
		fmt.Printf("db get failed, err: %v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

func queryMultiRowDemo() {
	sqlStr := "select id, name, age from users where id > ?"
	var users []user
	err := dbConn.Select(&users, sqlStr, 0)
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

func namedQuery() {
	sqlStr := "select * from users where name = :name"

	// rows, err := dbConn.NamedQuery(sqlStr, map[string]interface{}{
	// 	"name": "谭万铖最帅",
	// })
	// if err != nil {
	// 	fmt.Printf("dbConn.NamedQuery failed, err:%v", err)
	// 	return
	// }
	// defer rows.Close()

	u := user{
		Name: "谭万铖最帅",
	}
	rows, err := dbConn.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("dbConn.NamedQuery failed, err:%v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}

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
	sqlStr1 := "Update users set age=0 where id=?"
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
	sqlStr2 := "Update users set age=20 where id=?"
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

//BatchInsertUsers 自行实现批量插入，笨方法
func BatchInsertUsers(users []*Testuser) error {
	//存放(?,?)的slice
	valueStrings := make([]string, 0, len(users))
	//存放values的slice
	valueArgs := make([]interface{}, 0, len(users)*2)

	//遍历users准备相关数据
	for _, u := range users {
		//此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?,?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	//自行拼接要执行的具体语句
	sqlStr := fmt.Sprintf("insert into testuser (name,age) values %s", strings.Join(valueStrings, ","))
	fmt.Println(sqlStr)
	_, err := dbConn.Exec(sqlStr, valueArgs...)
	return err
}

//BatchInsertUsers2 sqlx.In批量插入
func BatchInsertUsers2(users []interface{}) error {
	fmt.Println(users)
	query, args, _ := sqlx.In(
		"insert into testuser(name,age) values (?),(?)",
		users...,
	)
	fmt.Println(query)
	fmt.Println(args)
	_, err := dbConn.Exec(query, args...)
	return err
}

//根据给定ID查询
func queryByIDs(ids []int) (users []Testuser, err error) {
	//动态填充id
	query, args, err := sqlx.In("select id,name,age from testuser where id in(?)", ids)
	if err != nil {
		return
	}
	//sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = dbConn.Rebind(query)
	err = dbConn.Select(&users, query, args...)
	return
}

//QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int) (users []Testuser, err error) {
	//动态填充id
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In("select id,name,age from testuser where id in (?) order by find_in_set(id,?)", ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}
	query = dbConn.Rebind(query)
	err = dbConn.Select(&users, query, args...)
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("db init failed, err:%v\n", err)
		return
	}
	fmt.Println("main.go: db init success!")
	fmt.Println()
	defer fmt.Println("main.go: db connect off")
	defer fmt.Println()
	defer dbConn.Close()

	//根据给定的id集合查询数据
	users := make([]Testuser, 0, 3)
	users2 := make([]Testuser, 0, 3)
	var id1 = 1
	var id2 = 2
	var id4 = 4
	ids := []int{id1, id2, id4}
	ids2 := []int{id4, id1, id2}
	users, _ = queryByIDs(ids)
	users2, _ = QueryAndOrderByIDs(ids2)
	fmt.Println(users)
	fmt.Println(users2)

	// users := make([]interface{}, 0, 3) //注意不能合并长度和容量的声明，否则会报panic，无效的内存地址或空指针
	// var u1 = Testuser{Age: 98, Name: "仙女1"}
	// var u2 = Testuser{Age: 99, Name: "仙女2"}
	// users = append(users, u1, u2)
	// // BatchInsertUsers(users)
	// fmt.Println(users)
	// BatchInsertUsers2(users)

	//namedQuery
	// namedQuery()

	//transactionDemo2
	// transactionDemo2()

	// //insertUserDemo
	// insertUserDemo()

	// //插入数据
	// insertRowDemo()

	// //多行查询
	// queryMultiRowDemo()

	// //单行查询
	// queryRowDemo()
}
