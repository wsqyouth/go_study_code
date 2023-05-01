package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	testTransactionID = "transaction-123"
	testDSN           = "coopers:2019Youth@tcp(localhost:3306)/sql_test"
)

type DBPool interface {
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
	dbs   DBPool
	uuid  UUIDGenerator
	mutex sync.RWMutex
}

func NewSessionPool(dbs DBPool, uuid UUIDGenerator) *SessionPool {
	return &SessionPool{
		m:    make(map[uint64]*session),
		dbs:  dbs,
		uuid: uuid,
	}
}

// AddTransaction Create a new session with the transaction
func (sp *SessionPool) AddTransaction() (uint64, error) {
	sp.mutex.Lock()
	defer sp.mutex.Unlock()

	db := sp.dbs.Get()
	tx, err := db.Begin()
	if err != nil {
		return 0, errors.Wrap(err, "failed to start transaction")
	}
	transID := sp.uuid.Generate()
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
		return nil, errors.New(fmt.Sprintf("transaction not found:%v", transID))
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

//execOperation  Execute database operations using the transaction
func execOperation(s *session) error {
	_, err := s.tx.ExecContext(context.Background(), "INSERT INTO user (name, age) VALUES (?, ?)", "wsqshiwu4", 28)
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

func testScenario(transactionID uint64, sp *SessionPool) {
	var s *session
	var err error

	if transactionID != 0 {
		// Use the existing session
		s, err = sp.GetTransaction(transactionID)
		if err != nil {
			fmt.Println("Error getting transaction:", err)
			return
		}
		fmt.Println("get existed trans: ", transactionID)
	} else {
		// Create a new session and put it into the session pool
		transactionID, err = sp.AddTransaction()
		if err != nil {
			fmt.Println("Error adding transaction:", err)
			return
		}
		s, err = sp.GetTransaction(transactionID)
		if err != nil {
			fmt.Println("Error getting transaction:", err)
			return
		}
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
}

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

	// Check if a transaction ID is provided as a command-line argument
	var transactionID uint64
	if len(os.Args) > 1 {
		transactionID, err = strconv.ParseUint(os.Args[1], 10, 64)
		if err != nil {
			fmt.Println("Error converting transaction ID:", err)
			return
		}
	}

	// Test case 1: No transaction ID provided
	go testScenario(0, sp)

	// Test case 2: Existing transaction ID provided
	// If a transaction ID is provided as a command-line argument, use it.
	// Otherwise, use a default transaction ID (e.g., 1).
	if transactionID != 0 {
		go testScenario(transactionID, sp)
	} else {
		go testScenario(1, sp)
	}
	// Keep the main process running
	for {
		time.Sleep(1 * time.Second)
	}
}

/*
总结：
1. 想模拟带事务和不带事务id两种请求下的场景,没有测试成功,整体思路是没问题的
2. 待改进的地方: sessionPool的初始化和事务提交主要功能没问题,但何时进行回收需要考虑
3. 针对单个连接:第一次和最后一次完成事务与session pool交互,中间携带事务id进行数据处理
*/
