# MySQL

MySQL 是业界常用的关系型数据库，本文介绍 Go 语言如何操作 MySQL 数据库。

## 连接

Go 语言中的`database/sql`标准库提供了保证 SQL 或类 SQL 数据库的范用接口，并不提供具体的数据库驱动，使用`database/sql`标准库时必须注入（至少）一个数据库驱动。

我们常用的数据库基本都有完整的第三方实现，例如：[MySQL 驱动](https://github.com/go-sql-driver/mysql)

### 下载依赖

```go
go get -u github.com/go-sql-driver/mysql
```

### 使用 MySQL 驱动

```go
func open(driverName, dataSourceName string) (*DB, error)
```

Open 打开一个 dirverName 指定的数据库，dataSourceName 指定数据源，一般至少包括数据库文件名和其它连接必要的信息。

```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //匿名导入，进初始化
)

func main() {
	//DSN：data Sourse Name
	dsn := "wancheng:wancheng@tcp(127.0.0.1:3306)/video_server"
	dbConn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("database connect success.")
	defer fmt.Println("database connect off")
	defer dbConn.Close() //注意这行代码要写在上面err判断的下面
}
```

**思考题**： 为什么上面代码中的`defer db.Close()`语句不应该写在`if err != nil`的前面呢？

**答案**：`defer`语句注册时，所有变量都必须存在且正确。所以要先对 err 做判断，确认 dbConn 变量存在且正确后再注册 defer 语句。如果 dbConn 错误，defer 语句还会执行而导致 panic。

### 初始化连接

Open 函数可能只是验证其参数格式是否正确，实际上并不创建于数据库的连接。如果要检查数据源的名称是否真实有效，应该调用 Ping 方法。

```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //匿名导入，进初始化
)

var (
	dbConn *sql.DB
)

func initDB() (err error) {
	//DSN：data Sourse Name
	dsn := "wancheng:wancheng@tcp(127.0.0.1:3306)/video_server"
	// 不会校验账号密码是否正确
	//注意：这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量 
	dbConn, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Database connect failed, err:", err)
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = dbConn.Ping()
	if err != nil {
		fmt.Println("Database doesn't exist.")
		return err
	}
	fmt.Println("database open success.")
	return nil
}

func main() {
	err = initDB()
	if err != nil {
		fmt.Println("init db failed,err:", err)
		return
	}
	fmt.Println("init db success")
	defer fmt.Println("database connect off")
	defer dbConn.Close() //注意这行代码要写在上面err判断的下面
}
```

其中`sql.DB`是表示连接的数据库对象（结构体实例），它保存了连接数据库相关的所有信息。它内部维护着一个具有零到多个底层连接的连接池，它可以安全地被多个 goroutine 同时使用。

###  SetConnMaxLifetime

```go
func (db *DB) SetConnMaxLifetime(duration time.Time)
```

设置连接最长存活时间。

### SetMaxOpenConns

```go
func (db *DB) SetOpenConns(n int)
```

`SetMaxOpenConns`设置与数据库建立连接的最大数目。 如果 n 大于 0 且小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制。 如果 n<=0，不会限制最大开启连接数，默认为 0（无限制）。

### SetMaxIdleConns

```go
func (db *DB) SetMaxIdleConns(n int)
```

SetMaxIdleConns 设置连接池中的最大闲置连接数。 如果 n 大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制。如果 n<=0，不会保持闲置连接。

## CURD

### 建库建表

我们先在 MySQL 中创建一个名为`sql_test`的数据库：

```go
CREATE DATABASE sql_test
```

进入该数据库：

```go
use sql_test
```

执行一下命令创建一张用户测试的数据表：

```go
CREATE TABLE `sql_test`.`users` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(20) NULL,
  `age` INT(11) NULL DEFAULT 0,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;
```

### 查询

为了方便查询，我们事先定义好一个结构体来存储 users 表的数据。

```go
type user struct{
	id int
	age int
	name string
}
```

#### 单行查询

单行查询`db.QueryRow()`执行一次查询，并期望返回最多一行结构（即 ROW）。QueryRow 总是返回非 nil 的值，直到返回值的 Scan 方法被调用时，才会返回被延迟的错误（如：未找到结果）。

```go
func (db *DB) QueryRow(query string, args ...interface{}) *Row
```

代码示例如下：

```go
func queryRowDemo() {
	sqlStr := "select id, name, age from users where id=?"
	var u user
	err := dbConn.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%v name:%v age:%d\n", u.id, u.name, u.age)
}
```

#### 多行查询

多行查询`db.Query()`执行一次查询，返回多行结果（即 Rows），一般用于执行 select 命令，参数 args 表示 query 中的占位参数。

```go
func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
```

代码示例如下：

```go
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
```

#### 插入数据

插入、更新和删除操作都使用`Exec`方法。

```go
func (db *DB) Exec(query string, args ...interface{}) (Result, error)
```

Exec 执行一次命令（包括查询、删除、更新、插入等），返回的 Result 是对已执行的 SQL 命令的总结。参数 args 表示 query 中的占位参数。

代码示例：

```go
func insertRowDemo() {
	sqlStr := "insert into users(name, age) values (?,?)"
	ret, err := dbConn.Exec(sqlStr, "雅婷宝贝仙女", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() //新插入的数据的in
	if err != nil {
		fmt.Printf("get lastinsert ID failed,err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d\n", theID)
}
```

#### 更新数据

代码示例如下：

```go
func updateRowDemo() {
	sqlStr := "update users set age = ? where id = ?"
	ret, err := dbConn.Exec(sqlStr, 39, 2)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows: %d\n", n)
}
```

#### 删除数据

代码示例如下：

```go
func deleteRowDemo() {
	sqlStr := "delete from users where id = ?"
	ret, err := dbConn.Exec(sqlStr, 2)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows: %d\n", n)
}
```

## MySQL 预处理

### 什么是预处理

普通 SQL 语句执行过程：

1、客户端对 SQL 语句进行占位符替换的到完整的 SQL 语句。

2、客户端发送完整的 SQL 语句到 MySQL 服务端。

3、MySQL 服务端执行完整的 SQL 语句并将结果返回客户端。

预处理执行过程：

1、把 SQL 语句分成两部分，命令部分与数据部分。

2、先把命令部分发送给 MySQL 服务端，MySQL 服务端进行 SQL 预处理。

3、然后把数据部分发送给 MySQL 服务端，MySQL 服务端对 SQL 语句进行占位符替换。

4、MySQL 服务端执行完整的 SQL 语句并将结果返回给客户端。

### 为什么要预处理

1、优化 MySQL 服务器重复执行 SQL 的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。

2、避免 SQL 注入问题。

### Go 实现 MySQL 预处理

`database/sql`中使用下面的`Prepare`方法来实现预处理操作：

```go
func (db *DB) Prepare(query string) (*Stmt, error)
```

`Prepare`方法会先将 sql 语句发送给 MySQL 服务端，返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。

查询操作的预处理示例代码如下：

```go
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
```

插入、更新和删除操作的预处理十分类似，这里以插入操作的预处理为例：

```go
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
```

### SQL 注入问题

**我们任何时候都不应该自己拼接 SQL 语句！**

这里我们演示一个自行拼接 SQL 语句的示例，编写一个根据 name 字段查询 user 表的函数如下：

```go
// sql注入示例
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u user
	err := db.QueryRow(sqlStr).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", u)
}
```

此时以下输入字符串都可以引发 SQL 注入问题：

```go
sqlInjectDemo("xxx' or 1=1#")
sqlInjectDemo("xxx' union select * from user #")
sqlInjectDemo("xxx' and (select count(*) from user) <10 #")
```

**补充：**不同的数据库中，SQL 语句使用的占位符语法不尽相同。

| 数据库     | 占位符语法   |
| :--------- | :----------- |
| MySQL      | `?`          |
| PostgreSQL | `$1`, `$2`等 |
| SQLite     | `?` 和`$1`   |
| Oracle     | `:name`      |

## Go 实现 MySQL 事务

### 什么是事务

一个最小的不可再分的工作单元，通常一个事务对应一个完整的业务（例如银行转账，该业务就是一个最小的工作单元），同时这个完整的业务需要执行多次的 DML（insert、update、delete）语句共同联合完成。A 转账给 B，这里面就需要执行两次 update 操作。

在 MySQL 中只有使用了 `Innodb `数据库引擎的数据库或表才支持事务。事务处理了可以用来维护数据库的完整性，保证成批的 SQL 语句要么全部执行，要么全部不执行。

### 事务的 ACID

事务通常必须满足 4 个条件（ACID）：原子性（atomicity，也称为不可分割性）、一致性（consistency）、隔离性（isolation，也称独立性）、持久性（durability）。

|  条件  | 解释                                                                                                                                                                                                                                                            |
| :----: | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 原子性 | 一个事务（transaction）中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。事务在执行过程中发生错误，会被回滚（Rollback）到事务开始前的状态，就像这个事务从来没有执行过一样。                                                                   |
| 一致性 | 在事务开始之前和事务结束以后，数据库的完整性没有被破坏。这表示写入的资料必须完全符合所有的预设规则，这包含资料的精确度、串联性以及后续数据库可以自发性地完成预定的工作。                                                                                        |
| 隔离性 | 数据库允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致。事务隔离分为不同级别，包括读未提交（Read uncommitted）、读提交（read committed）、可重复读（repeatable read）和串行化（Serializable）。 |
| 持久性 | 事务处理结束后，对数据的修改就是永久的，即便系统故障也不会丢失。                                                                                                                                                                                                |

### 事务相关方法

Go 语言中使用以下三个方法实现 MySQL 中的事务操作。

开始事务

```go
func (db *DB) Begin() (*Tx, error)
```

提交事务：

```go
func (tx *Tx) Commit() error
```

回滚事务：

```go
func (tx *Tx) Rollback() error
```

### 事务示例

下面的代码演示了一个简单的事务操作，该事物操作能够确保两次更新操作要么同时成功要么同时失败，不会存在中间状态。

```go
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

	//一项更新
	sqlStr1 := "update users set age = 18 where id = ?"
	ret1, err := tx.Exec(sqlStr1, 1)
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
	sqlStr2 := "update users set age = 16 where id = ?"
	ret2, err := tx.Exec(sqlStr2, 2)
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
```

# sqlx

在项目中我们通常会使用`database/sql`连接 MySQL 数据库，本文借助使用`sqlx`实现批量插入数据的例子，介绍了`sqlx`中可能被你忽视了的`sqlx.In`和`DB.NameExe c`方法。

## sqlx 介绍

在项目中我们通常可能会使用`database/sql`连接 MySQL 数据库，`sqlx`可以认为是 Go 语言内置`database/sql`的超集，它在优秀的内置`database/sql`基础上提供了一组扩展，这些扩展中除了大家常用来查询的` Get(dest interface{}, ...)) error`和`Select(dest interface{}, ...)) error`外还有很多其他强大的功能。

## 安装 sqlx

```go
go get github.com/jmoiron/sqlx
```

## 基本使用

### 连接数据库

```go
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
```

### 查询

查询单行数据，代码示例如下：

```go
package main

import (
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
	sqlStr := "select id, age, name from users where id = ?"
	var u user
	err = dbConn.Get(&u, sqlStr, 3)
	if err != nil {
		fmt.Printf("db get failed, err: %v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
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

	//单行查询
	queryRowDemo()
}
```

查询多行数据，代码示例如下：

```go
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
```

### 插入、更新和删除

sqlx 中的 exec 方法与原生 sql 中的 exec 使用基本一致：

```go
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

func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 6)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
```

### NamedExec

`DB.NamedExec`方法用来绑定 SQL 语句和结构体或 map 中的同名字段。

```go
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
```

### NamedQuery

```go
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
```

### 事务操作

对于事务操作，我们可以使用`sqlx`中提供的`db.Beginx()`和`tx.Exec()`方法。示例代码如下：

```go
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
```

## sqlx.In

`sqlx.In`是`sqlx`提供的一个非常方便的函数。

### sqlx.In的批量插入示例

#### 表结构

```
CREATE TABLE `sql_test`.`testuser` (  
`id` INT NOT NULL,  
`name` VARCHAR(45) NULL,  
`age` INT NULL,  PRIMARY KEY (`id`)
)
ENGINE = InnoDB 
DEFAULT CHARACTER SET = utf8mb4;
```

#### 结构体

定义一个`user`结构体，字段通过tag与数据库中user表的列一致：

```go
type User struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}
```

#### bindvars（绑定变量）

查询占位符`?`在内部称为***bindvars（查询占位符）***,它非常重要。你应该始终使用它们向数据库发送值，因为它们可以防止SQL注入攻击。`database/sql`不尝试对查询文本进行任何验证；它与编码的参数一起按原样发送到服务器。除非驱动程序实现一个特殊的接口，否则在执行之前，查询是在服务器上准备的。因此`bindvars`是特定于数据库的：

- MySQL中使用`?`
- PostgreSQL使用枚举的`$1`、`$2`等bindvar语法
- SQLite中`?`和`$1`的语法都支持
- Oracle中使用`:name`的语法

`bindvars`的一个常见误解是，它们用来在sql语句中插入值。它们其实仅用于参数化，不允许更改SQL语句的结构。例如，使用`bindvars`尝试参数化列或表名将不起作用：

```go
// ？不能用来插入表名（做SQL语句中表名的占位符）
db.Query("SELECT * FROM ?", "mytable")
 
// ？也不能用来插入列名（做SQL语句中列名的占位符）
db.Query("SELECT ?, ? FROM people", "name", "location")
```

#### 自己拼接语句实现批量插入

比较笨，但是很好理解。就是有多少个User就拼接多少个`(?,?)`：

```go
//BatchInsertUsers自行构造批量插入的语句
func BatchInsertUsers(users []*User) error {
    
}//BatchInsertUsers 自行实现批量插入，笨方法
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

func(){
    users := make([]*Testuser, 0, 3) //注意不能合并长度和容量的声明，否则会报panic，无效的内存地址或空指针
	var u1 = &Testuser{Age: 18, Name: "moqqll"}
	var u2 = &Testuser{Age: 20, Name: "小王子"}
	users = append(users, u1, u2)
	BatchInsertUsers(users)
}
```

#### 使用sqlx.In实现批量插入

首先我们需要给我们的结构体实现`driver.Valuer`接口：

```go
//Value ...
func (u Testuser) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}
```

使用`sqlx.In`实现批量插入代码如下：

```go
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
```

#### 使用NamedExec实现批量插入

**注意** ：该功能目前有人已经推了[#285 PR](https://github.com/jmoiron/sqlx/pull/285)，但是作者还没有发`release`，所以想要使用下面的方法实现批量插入需要暂时使用`master`分支的代码：

```
go get github.com/jmoiron/sqlx@master
```

使用`NamedExec`实现批量插入的代码如下：

```go
// BatchInsertUsers3 使用NamedExec实现批量插入
func BatchInsertUsers3(users []*User) error {
	_, err := DB.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
	return err
}
```

### sqlx.In的查询示例

关于`sqlx.In`这里再补充一个用法，在`sqlx`查询语句中实现In查询和FIND_IN_SET函数。即实现`SELECT * FROM user WHERE id in (3, 2, 1);`和`SELECT * FROM user WHERE id in (3, 2, 1) ORDER BY FIND_IN_SET(id, '3,2,1')`

#### in查询

查询id在给定id集合中的数据。

```go
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


func main(){
    //根据给定的id集合查询数据
	users := make([]Testuser, 0, 3)
	var id1 = 1
	var id2 = 2
	var id4 = 4
	ids := []int{id1, id2, id4}
	users, _ = queryByIDs(ids)
	fmt.Println(users)
}
```

#### in查询和FIND_IN_SET函数

```go
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
```

**注意**：当然，在这个例子里面你也可以先使用`IN`查询，然后通过代码按给定的ids对查询结果进行排序。











