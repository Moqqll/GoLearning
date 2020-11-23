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

**答案**：`defer`语句注册时，所有变量都必须存在且正确。所以要先对err做判断，确认dbConn变量存在且正确后再注册defer语句。如果dbConn错误，defer语句还会执行而导致panic。

### 初始化连接

Open函数可能只是验证其参数格式是否正确，实际上并不创建于数据库的连接。如果要检查数据源的名称是否真实有效，应该调用Ping方法。

```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //匿名导入，进初始化
)

var (
	dbConn *sql.DB
	err    error
)

func initDB() error {
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

其中`sql.DB`是表示连接的数据库对象（结构体实例），它保存了连接数据库相关的所有信息。它内部维护着一个具有零到多个底层连接的连接池，它可以安全地被多个goroutine同时使用。

### SetMaxOpenConns

```go
func (db *DB) SetOpenConns(n int)
```

`SetMaxOpenConns`设置与数据库建立连接的最大数目。 如果n大于0且小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制。 如果n<=0，不会限制最大开启连接数，默认为0（无限制）。

### SetMaxIdleConns

```go
func (db *DB) SetMaxIdleConns(n int)
```

SetMaxIdleConns设置连接池中的最大闲置连接数。 如果n大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制。如果n<=0，不会保持闲置连接。

## CURD

### 建库建表

我们先在MySQL中创建一个名为`sql_test`的数据库：

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

为了方便查询，我们事先定义好一个结构体来存储users表的数据。

```go
type user struct{
	id int
	age int
	name string
}
```

#### 单行查询

单行查询`db.QueryRow()`执行一次查询，并期望返回最多一行结构（即ROW）。QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误（如：未找到结果）。

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

多行查询`db.Query()`执行一次查询，返回多行结果（即Rows），一般用于执行select命令，参数args表示query中的占位参数。

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

Exec执行一次命令（包括查询、删除、更新、插入等），返回的Result是对已执行的SQL命令的总结。参数args表示query中的占位参数。

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

## MySQL预处理

### 什么是预处理

普通SQL语句执行过程：

1、客户端对SQL语句进行占位符替换的到完整的SQL语句。

2、客户端发送完整的SQL语句到MySQL服务端。

3、MySQL服务端执行完整的SQL语句并将结果返回客户端。

预处理执行过程：

1、把SQL语句分成两部分，命令部分与数据部分。

2、先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。

3、然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。

4、MySQL服务端执行完整的SQL语句并将结果返回给客户端。

### 为什么要预处理

1、优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。

2、避免SQL注入问题。

### Go实现MySQL预处理

`database/sql`中使用下面的`Prepare`方法来实现预处理操作：

```go
func (db *DB) Prepare(query string) (*Stmt, error)
```

`Prepare`方法会先将sql语句发送给MySQL服务端，返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。

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

### SQL注入问题

**我们任何时候都不应该自己拼接SQL语句！**

这里我们演示一个自行拼接SQL语句的示例，编写一个根据name字段查询user表的函数如下：

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

此时以下输入字符串都可以引发SQL注入问题：

```go
sqlInjectDemo("xxx' or 1=1#")
sqlInjectDemo("xxx' union select * from user #")
sqlInjectDemo("xxx' and (select count(*) from user) <10 #")
```

**补充：**不同的数据库中，SQL语句使用的占位符语法不尽相同。

| 数据库     | 占位符语法   |
| :--------- | :----------- |
| MySQL      | `?`          |
| PostgreSQL | `$1`, `$2`等 |
| SQLite     | `?` 和`$1`   |
| Oracle     | `:name`      |

## Go实现MySQL事务

### 什么是事务

一个最小的不可再分的工作单元，通常一个事务对应一个完整的业务（例如银行转账，该业务就是一个最小的工作单元），同时这个完整的业务需要执行多次的DML（insert、update、delete）语句共同联合完成。A转账给B，这里面就需要执行两次update操作。

在MySQL中只有使用了Innodb数据库引擎的数据库或表才支持事务。事务处理了可以用来维护数据库的完整性，保证成批的SQL语句要么全部执行，要么全部不执行。

### 事务的ACID

事务通常必须满足4个条件（ACID）：原子性（atomicity，也称为不可分割性）、一致性（consistency）、隔离性（isolation，也称独立性）、持久性（durability）。

|  条件  | 解释                                                         |
| :----: | :----------------------------------------------------------- |
| 原子性 | 一个事务（transaction）中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。事务在执行过程中发生错误，会被回滚（Rollback）到事务开始前的状态，就像这个事务从来没有执行过一样。 |
| 一致性 | 在事务开始之前和事务结束以后，数据库的完整性没有被破坏。这表示写入的资料必须完全符合所有的预设规则，这包含资料的精确度、串联性以及后续数据库可以自发性地完成预定的工作。 |
| 隔离性 | 数据库允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致。事务隔离分为不同级别，包括读未提交（Read uncommitted）、读提交（read committed）、可重复读（repeatable read）和串行化（Serializable）。 |
| 持久性 | 事务处理结束后，对数据的修改就是永久的，即便系统故障也不会丢失。 |

### 事务相关方法

Go语言中使用以下三个方法实现MySQL中的事务操作。 

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

在项目中我们通常会使用`database/sql`连接MySQL数据库，本文借助使用`sqlx`实现批量插入数据的例子，介绍了`sqlx`中可能被你忽视了的`sqlx.In`和`DB.NameExec`方法。

## sqlx介绍

在项目中我们通常可能会使用`database/sql`连接MySQL数据库，`sqlx`可以认为是Go语言内置`database/sql`的超集，它在优秀的内置`database/sql`基础上提供了一组扩展，这些扩展中除了大家长用来查询的`	Get(dest interface{}, ...)) error`和`Select(dest interface{}, ...)) error`外还有很多其他强大的功能。

## 安装sqlx

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

sqlx中的exec方法与原生sql中的exec使用基本一致：

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

`DB.NamedExec`方法用来绑定SQL语句和结构体或map中的同名字段。

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







### 事务操作

对于事务操作，我们可以使用`sqlx`中提供的`db.Beginx()`和`tx.Exec()`方法。示例代码如下：

```

```















