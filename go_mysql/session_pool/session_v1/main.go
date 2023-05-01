package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"sync"
	"time"

	// we have to import the driver, but don't use it in our code
	// so we use the `_` symbol
	_ "github.com/go-sql-driver/mysql"
)

const (
	testTransactionID = "transaction-123"
	testDSN           = "coopers:2019Youth@tcp(localhost:3306)/sql_test"
)

type session struct {
	transactionID string
	tx            *sql.Tx
	db            *sql.DB
	txBeginTime   int64
	mutex         sync.Mutex
}

type SessionPool struct {
	mutex sync.Mutex
	m     map[string]*session
	db    *sql.DB
}

func NewSessionPool(db *sql.DB) *SessionPool {
	return &SessionPool{
		m:  make(map[string]*session),
		db: db,
	}
}

func (p *SessionPool) GetSession(ctx context.Context, transID string) (*session, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	// Check if a session with the given transaction ID already exists
	if s, ok := p.m[transID]; ok {
		fmt.Println("Transaction use exist")
		return s, nil
	}
	fmt.Println("Transaction new created: ", transID)
	// Start a new transaction
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to start transaction")
	}
	// Create a new session with the transaction
	s := &session{
		transactionID: transID,
		tx:            tx,
		db:            p.db,
		txBeginTime:   time.Now().Unix(),
		mutex:         sync.Mutex{},
	}
	p.m[transID] = s
	return s, nil
}

func (p *SessionPool) ReleaseSession(transID string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	fmt.Println("Transaction freed:", transID)
	delete(p.m, transID)
}

func (s *session) Rollback() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// Rollback the transaction
	if err := s.tx.Rollback(); err != nil {
		return errors.Wrap(err, "failed to rollback transaction")
	}
	return nil
}

func (s *session) Commit() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// Commit the transaction
	if err := s.tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}

func main() {
	// Open a database connection
	db, err := sql.Open("mysql", testDSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Create a session pool
	sessionPool := NewSessionPool(db)
	// Start a new transaction using a session from the pool
	s, err := sessionPool.GetSession(context.Background(), testTransactionID)
	if err != nil {
		panic(err)
	}
	// Execute database operations using the transaction
	_, err = s.tx.ExecContext(context.Background(), "INSERT INTO user (name, age) VALUES (?, ?)", "hww", 27)
	if err != nil {
		// Rollback the transaction on error
		if err := s.Rollback(); err != nil {
			panic(err)
		}
		panic(err)
	}
	// Commit the transaction
	if err := s.Commit(); err != nil {
		panic(err)
	}
	fmt.Println("Transaction completed successfully")
	time.Sleep(time.Second * 10)
	// Release the session
	sessionPool.ReleaseSession(testTransactionID)
}

/*
总结：
1. 使用事务及sesson_pool固定值测试
*/
