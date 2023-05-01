package main

import (
	"context"
	"database/sql"
	"errors"
	// "fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(*sql.Tx), args.Error(1)
}

type MockTx struct {
	mock.Mock
}

func (m *MockTx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	mockArgs := m.Called(ctx, query, args)
	return mockArgs.Get(0).(sql.Result), mockArgs.Error(1)
}

func (m *MockTx) Commit() error {
	return m.Called().Error(0)
}

func (m *MockTx) Rollback() error {
	return m.Called().Error(0)
}

func TestSessionPool(t *testing.T) {
	mockDB := new(MockDB)
	mockTx := new(MockTx)

	mockDB.On("BeginTx", mock.Anything, mock.Anything).Return(mockTx, nil)
	sessionPool := NewSessionPool(mockDB)

	s, err := sessionPool.GetSession(context.Background(), testTransactionID)
	if err != nil {
		t.Fatal(err)
	}

	mockTx.On("ExecContext", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	_, err = s.tx.ExecContext(context.Background(), "INSERT INTO users (name, email) VALUES (?, ?)", "John Doe", "johndoe@example.com")
	if err != nil {
		t.Fatal(err)
	}

	mockTx.On("Commit").Return(nil)
	err = s.Commit()
	if err != nil {
		t.Fatal(err)
	}

	sessionPool.ReleaseSession(testTransactionID)

	s, err = sessionPool.GetSession(context.Background(), testTransactionID)
	if err != nil {
		t.Fatal(err)
	}

	mockTx.On("ExecContext", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("execContext error"))
	_, err = s.tx.ExecContext(context.Background(), "INSERT INTO users (name, email) VALUES (?, ?)", "John Doe", "johndoe@example.com")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	mockTx.On("Rollback").Return(nil)
	err = s.Rollback()
	if err != nil {
		t.Fatal(err)
	}
}
