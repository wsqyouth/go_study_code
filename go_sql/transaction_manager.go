package main

import (
	"fmt"
	"sync"
	"time"
)

type Transaction struct {
	ID         uint64
	Operation  string
	Connection string
	StartTime  int64
	Deadline   time.Time
	mutex      sync.Mutex
}

type TransactionManager struct {
	transactions map[uint64]*Transaction
	dbConnection string
	uuid         uint64
	mutex        sync.RWMutex
}

func (tm *TransactionManager) AddTransaction(id uint64, connection string, operation string, deadline time.Time) (*Transaction, error) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	trans := &Transaction{
		ID:         id,
		Connection: connection,
		Operation:  operation,
		StartTime:  time.Now().Unix(),
		Deadline:   deadline,
	}

	tm.transactions[id] = trans

	return trans, nil
}

func (tm *TransactionManager) GetTransaction(id uint64) (*Transaction, error) {
	tm.mutex.RLock()
	defer tm.mutex.RUnlock()

	trans, ok := tm.transactions[id]
	if !ok {
		return nil, fmt.Errorf("transaction not found")
	}

	return trans, nil
}

func main() {
	manager := &TransactionManager{
		transactions: make(map[uint64]*Transaction),
		dbConnection: "mockDB",
	}

	transID := uint64(1)
	connection := "mockDB"
	operation := "mockOperation"
	deadline := time.Now().Add(1 * time.Hour)

	trans, err := manager.AddTransaction(transID, connection, operation, deadline)
	if err != nil {
		fmt.Println("Error adding transaction:", err)
		return
	}

	fmt.Println("Added transaction:", trans)

	retrievedTrans, err := manager.GetTransaction(transID)
	if err != nil {
		fmt.Println("Error retrieving transaction:", err)
		return
	}

	fmt.Println("Retrieved transaction:", retrievedTrans)
}

/*
设计的好处主要有以下几点：
事务管理：通过将事务信息（如事务ID、数据库连接、事务开始时间等）封装在transaction结构中，可以方便地管理和操作事务。
并发控制：使用互斥锁可以保证在并发环境下对transaction的安全操作。
资源复用：通过transactionManager，可以复用数据库连接和事务资源，避免频繁地创建和销毁资源，提高性能。
要实现多个库实现事务的隔离，可以为每个数据库连接创建一个独立的transaction，并在transactionManager中管理这些transaction。
		每个transaction都有自己的事务ID和数据库连接，因此可以独立地进行事务操作，实现事务的隔离。

对数据库的事务操作可以通过transaction的方法来实现。例如，可以在transaction中添加Commit和Rollback方法，分别用于提交和回滚事务。
transID是事务的唯一标识，通常需要存储在持久化的存储系统中，如数据库或分布式缓存中。在事务开始时生成transID，并在事务结束后删除。

这个设计还有一些可以优化的点：
错误处理：在当前的设计中，如果开始事务或添加transaction失败，没有进行错误处理。可以添加错误处理逻辑，如重试或回滚操作。
超时处理：可以添加事务超时处理逻辑，如果事务在规定的时间内没有完成，可以自动回滚事务。
连接池管理：可以添加连接池管理逻辑，如连接的获取、释放、健康检查等。
分布式事务：如果需要支持分布式事务，可以考虑使用两阶段提交（2PC）或三阶段提交（3PC）等协议。
*/
