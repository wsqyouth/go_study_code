package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

var n1 int64

func addSyncByAtomic(delta int64) int64 {
	return atomic.AddInt64(&n1, delta)
}

func readSyncByAtomic() int64 {
	return atomic.LoadInt64(&n1)
}

var n2 int64
var rwmu sync.RWMutex

func addSyncByRWMutex(delta int64) {
	rwmu.Lock()
	n2 += delta
	rwmu.Unlock()
}

func readSyncByRWMutex() int64 {
	var n int64
	rwmu.RLock()
	n = n2
	rwmu.RUnlock()
	return n
}

func BenchmarkAddSyncByAtomic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			addSyncByAtomic(1)
		}
	})
}

func BenchmarkReadSyncByAtomic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			readSyncByAtomic()
		}
	})
}

func BenchmarkAddSyncByRWMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			addSyncByRWMutex(1)
		}
	})
}

func BenchmarkReadSyncByRWMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			readSyncByRWMutex()
		}
	})
}

// $go test -bench . go-atomic-package-1_test.go -cpu 2
// 读写锁的性能随着并发量增大的变化情况与前面讲解sync.RWMutex时的一致；
// 利用原子操作的无锁并发写的性能随着并发量增大几乎保持恒定；
// 利用原子操作的无锁并发读的性能随着并发量增大有持续提升的趋势，并且性能约为读锁的200倍。
