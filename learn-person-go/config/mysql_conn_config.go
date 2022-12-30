package config

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"sync"
)

var (
	PoolUnInvaildSizeError = errors.New("pool size is unvaild")
	PoolIsClosedError      = errors.New("pool had closed")
)

// 连接池定义
type Pool struct {
	sync.Mutex                // 保证连接池线程安全
	Size       int            // 连接池连接数量
	ConnChan   chan io.Closer // 存储连接的管道
	IsClose    bool
	ctx        context.Context
}

func init() {
	NewConnSize(100)
}

//创建链接
func NewConnSize(size int) (*Pool, error) {
	if size <= 0 {
		return nil, PoolUnInvaildSizeError
	}
	return &Pool{
		ConnChan: make(chan io.Closer, size),
		ctx:      context.Background(),
	}, nil
}

//获取链接
func (pool *Pool) GetConnFormPool() (io.Closer, error) {
	if pool.IsClose == true {
		return nil, PoolIsClosedError
	}
	select {
	case conn, ok := <-pool.ConnChan:
		if !ok {
			return nil, PoolIsClosedError
		}
		fmt.Println("获取链接：", conn)
		return conn, nil
	default:
		return pool.getNewConn(pool.ctx)
	}
}

func (pool *Pool) getNewConn(ctx context.Context) (io.Closer, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_test?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal("数据库连接失败", err)
		return nil, err
	}
	conn, _ := db.Conn(ctx)
	select {
	case pool.ConnChan <- conn:
		fmt.Println("连接放入连接池")
	default:
		fmt.Println("连接池满了，连接丢弃")
		conn.Close()
	}
	return conn, nil
}

// 释放连接
func (pool *Pool) ReleaseConn(conn io.Closer) error {
	pool.Lock()
	defer pool.Unlock()

	if pool.IsClose == true {
		return PoolIsClosedError
	}

	select {
	case pool.ConnChan <- conn:
		fmt.Println("连接已放回", conn)
	default:
		fmt.Println("连接池满了，连接丢弃")
		conn.Close()
	}
	return nil
}

// 关闭连接池
func (pool *Pool) ClosePool() error {
	pool.Lock()
	defer pool.Unlock()

	if pool.IsClose == true {
		return PoolIsClosedError
	}

	pool.IsClose = true
	close(pool.ConnChan)

	for conn := range pool.ConnChan {
		conn.Close()
	}
	return nil
}
