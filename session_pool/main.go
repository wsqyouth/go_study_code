package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"sync"
	"time"
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
		return s, nil
	}
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
	// Remove the session from the pool
	delete(sessionPool.m, s.transactionID)
	return nil
}

func main() {
	// Open a database connection
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Create a session pool
	sessionPool := NewSessionPool(db)
	// Start a new transaction using a session from the pool
	s, err := sessionPool.GetSession(context.Background(), "transaction-123")
	if err != nil {
		panic(err)
	}
	// Execute database operations using the transaction
	_, err = s.tx.ExecContext(context.Background(), "INSERT INTO users (name, email) VALUES (?, ?)", "John Doe", "johndoe@example.com")
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
}
