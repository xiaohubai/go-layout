package ants

import (
	"log"
	"sync"

	ants "github.com/panjf2000/ants/v2"
)

var _pool *ants.Pool
var _poolOnce sync.Once

const PoolSize = 16

func init() {
	if _, err := DefaultPool(); err != nil {
		panic(err)
	}
}

// DefaultPool 协程池
func DefaultPool() (p *ants.Pool, err error) {
	if _pool != nil {
		return _pool, nil
	}
	_poolOnce.Do(func() {
		p, err = ants.NewPool(PoolSize)
	})
	if err != nil {
		return
	}
	_pool = p
	return _pool, nil
}

func NewPool(size int, opts ...ants.Option) (*ants.Pool, error) {
	return ants.NewPool(size, opts...)
}

// Go submits a task to pool.
func Go(task func()) {
	if err := _pool.Submit(task); err != nil {
		log.Fatalln(err)
	}
}

// Submit submits a task to pool.
func Submit(task func()) error {
	return _pool.Submit(task)
}

// Running returns the number of the currently running goroutines.
func Running() int {
	return _pool.Running()
}

// Cap returns the capacity of this default pool.
func Cap() int {
	return _pool.Cap()
}

// Free returns the available goroutines to work.
func Free() int {
	return _pool.Free()
}

// Release Closes the default pool.
func Release() {
	_pool.Release()
}
