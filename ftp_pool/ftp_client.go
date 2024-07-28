package main

import (
	"bytes"
	"fmt"
	"io"

	"awesomeProject/ftp_pool/ftputil"
	"go.uber.org/zap"
)

type FtpPoolClient struct {
	pool *ftputil.FtpConnectionPool
}

// NewFtpPoolClient ftp pool client
func NewFtpPoolClient(poolCapacity int, serverAddress, username, password string) (*FtpPoolClient, error) {
	pool, err := ftputil.NewFtpConnectionPool(poolCapacity, serverAddress, username, password)
	if err != nil {
		return nil, err
	}

	return &FtpPoolClient{
		pool: pool,
	}, nil
}

// StoreFile store file
func (fp *FtpPoolClient) StoreFile(path string, data []byte) error {
	if fp.pool == nil {
		return ftputil.ErrPoolEmpty
	}
	p := fp.pool

	conn, err := p.GetFtpConnection()
	if err != nil {
		return err
	}

	defer func() {
		if err := p.PutFtpConnection(conn); err != nil {
			fmt.Printf("putFtpConnection error", zap.Error(err))
		}
	}()

	buffer := bytes.NewBuffer(data)
	err = conn.Stor(path, buffer)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	// 测试时打印属性
	// availableConns, activeConns := p.ReportStats()
	// fmt.Printf("get_ftp_pool", zap.Int32("AvailConnections", availableConns),
	// 	zap.Int32("ActiveConnections", activeConns), zap.String("path", path))

	return nil
}

// ReadFile read file
func (fp *FtpPoolClient) ReadFile(path string) ([]byte, error) {
	if fp.pool == nil {
		return nil, ftputil.ErrPoolEmpty
	}

	p := fp.pool
	conn, err := p.GetFtpConnection()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := p.PutFtpConnection(conn); err != nil {
			fmt.Printf("putFtpConnection error", zap.Error(err))
		}
	}()

	reader, err := conn.Retr(path)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %w", err)
	}

	return io.ReadAll(reader)
}

// Close ftp pool close
func (fp *FtpPoolClient) Close() error {
	if fp.pool != nil {
		fmt.Println("ftpPool pool Close")
		fp.pool.Close()
	}
	return nil
}
