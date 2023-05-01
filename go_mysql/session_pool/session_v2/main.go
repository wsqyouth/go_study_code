package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"os"
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

type dbPool interface {
	Get() *sql.DB
}

type UUIDGenerator interface {
	Generate() uint64
}

type SimpleDBPool struct {
	db *sql.DB
}

func (p *SimpleDBPool) Get() *sql.DB {
	return p.db
}

type SimpleUUIDGenerator struct {
	counter uint64
}

func (sug *SimpleUUIDGenerator) Generate() uint64 {
	sug.counter++
	return sug.counter
}

type session struct {
	transactionID uint64
	tx            *sql.Tx
	db            *sql.DB
	txBeginTime   int64
	mutex         sync.Mutex
}

type SessionPool struct {
	m     map[uint64]*session
	dbs   dbPool
	uuid  UUIDGenerator
	mutex sync.RWMutex
}

func NewSessionPool(dbs dbPool, uuid UUIDGenerator) *SessionPool {
	return &SessionPool{
		m:    make(map[uint64]*session),
		dbs:  dbs,
		uuid: uuid,
	}
}

func (sp *SessionPool) AddTransaction() (uint64, error) {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()

	db := sp.dbs.Get()
	tx, err := db.Begin()
	if err != nil {
		return 0, errors.Wrap(err, "failed to start transaction")
	}
	transID := sp.uuid.Generate()
	// Create a new session with the transaction
	sp.m[transID] = &session{
		transactionID: transID,
		tx:            tx,
		db:            db,
		txBeginTime:   time.Now().Unix(),
		mutex:         sync.Mutex{},
	}

	return transID, nil
}

func (sp *SessionPool) RemoveTransaction(transactionID uint64) error {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()

	s, ok := sp.m[transactionID]
	if !ok {
		return errors.New("transaction not found")
	}

	s.tx.Rollback()
	delete(sp.m, transactionID)
	return nil
}

func (sp *SessionPool) GetTransaction(transID uint64) (*session, error) {
	sp.mutex.RLock()
	defer sp.mutex.RUnlock()
	s, ok := sp.m[transID]
	if !ok {
		return nil, errors.New("transaction not found")
	}
	return s, nil
}

func (sp *SessionPool) CloseDBs() (err error) {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()

	for _, s := range sp.m {
		if err = s.tx.Rollback(); err != nil {
			return errors.Wrap(err, "failed to rollback transaction")
		}
	}
	sp.dbs.Get().Close()
	return nil
}

func (s *session) transCommit() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if err := s.tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	fmt.Println("Commit the transaction")
	return nil
}

func (s *session) transRollback() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if err := s.tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to Rollback transaction")
	}
	fmt.Println("Rollback the transaction")
	return nil
}

func execOperation(s *session) error {
	// Execute database operations using the transaction
	_, err := s.tx.ExecContext(context.Background(), "INSERT INTO user (name, age) VALUES (?, ?)", "wsq1", 27)
	if err != nil {
		// Rollback the transaction on error
		if err := s.transRollback(); err != nil {
			return errors.Wrap(err, "failed to rollback transaction")
		}
		return errors.Wrap(err, "failed to execute operation")
	}
	// Commit the transaction
	if err := s.transCommit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}
	return nil
}

// main 测试事务新建场景
func main() {
	// Open a database connection
	db, err := sql.Open("mysql", testDSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dbPool := &SimpleDBPool{db: db}
	uuid := &SimpleUUIDGenerator{}
	sp := NewSessionPool(dbPool, uuid)

	transactionID, err := sp.AddTransaction()
	if err != nil {
		fmt.Println("Error adding transaction:", err)
		return
	}

	s, err := sp.GetTransaction(transactionID)
	if err != nil {
		fmt.Println("Error getting transaction:", err)
		return
	}
	fmt.Println("Transaction started successfully:", transactionID)

	// Execute database operations using the transaction
	err = execOperation(s)
	if err != nil {
		fmt.Println("Error executing operation:", err)
		return
	}
	fmt.Println("Transaction execOperation successfully")

	err = sp.RemoveTransaction(transactionID)
	if err != nil {
		fmt.Println("Error removing transaction:", err)
		return
	}
	fmt.Println("Transaction ended successfully")
	sp.CloseDBs()
}

/*
总结：
1. 对代码进行进一步抽离,实现一个事务的初始化、创建和销毁功能
*/
