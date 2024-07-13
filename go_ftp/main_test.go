package main

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

func BenchmarkFTPClient_StoreFileWithMaxConnections(b *testing.B) {
	// Assume NewFTPConnectionPool has been called elsewhere to initialize the pool
	// with a maxConns value of 4. For example:
	pool, err := NewFTPConnectionPool("localhost:2100", "jovy", "123456", 5)
	if err != nil {
		log.Fatalf("Failed to initialize FTP connection pool: %v", err)
	}
	defer pool.Close()

	var wg sync.WaitGroup
	buffer := []byte("test data for benchmarking")

	b.ResetTimer()

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// Use the connection pool to store the file
			err := pool.StoreFileWithPool(fmt.Sprintf("file_%d.txt", i), buffer)
			if err != nil {
				b.Errorf("Failed to store file: %v", err)
			}
		}(i)
	}

	wg.Wait()
}
