package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	// we have to import the driver, but don't use it in our code
	// so we use the `_` symbol
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int
	Name string
	Age  int
}

type Dao interface {
	InsertUser(ctx context.Context, user User) error
	SelectUser(ctx context.Context, name string) ([]*User, error)
}

type UserDao struct {
	db *sql.DB
}

func (d *UserDao) InsertUser(ctx context.Context, user User) error {
	_, err := d.db.ExecContext(ctx, "INSERT INTO user (name, age) VALUES (?, ?)", user.Name, user.Age)
	return err
}

func (d *UserDao) SelectUser(ctx context.Context, name string) ([]*User, error) {
	rows, err := d.db.QueryContext(ctx, "SELECT id, name, age FROM user WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

type Transactor interface {
	Transact(ctx context.Context, txFunc func(ctx context.Context) error) error
}

type DBTransactor struct {
	db *sql.DB
}

func (t *DBTransactor) Transact(ctx context.Context, txFunc func(ctx context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = txFunc(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func main() {
	// 链接数据库
	dsn := "coopers:2019Youth@tcp(localhost:3306)/sql_test?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	ctx := context.Background()

	userDao := &UserDao{db: db}
	transactor := &DBTransactor{db: db}
	// 事务插入
	fmt.Println("Transaction example:")
	err = transactor.Transact(ctx, func(ctx context.Context) error {
		return userDao.InsertUser(ctx, User{Name: "Alice", Age: 30})
	})
	if err != nil {
		fmt.Println("Transaction failed:", err)
	} else {
		fmt.Println("Transaction succeeded.")
	}
	// 非事务查询
	users, err := userDao.SelectUser(ctx, "Alice")
	if err != nil {
		fmt.Println("Transaction failed:", err)
	}
	fmt.Println("Alice's users found:")
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}
