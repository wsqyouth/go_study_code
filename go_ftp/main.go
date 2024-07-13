package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/jlaffaye/ftp"
)

type FTPConnectionPool struct {
	conns    chan *ftp.ServerConn
	maxConns int
}

func NewFTPConnectionPool(server, username, password string, maxConns int) (*FTPConnectionPool, error) {
	pool := &FTPConnectionPool{
		conns:    make(chan *ftp.ServerConn, maxConns),
		maxConns: maxConns,
	}

	for i := 0; i < maxConns; i++ {
		conn, err := ftp.Dial(server, ftp.DialWithTimeout(5*time.Second))
		if err != nil {
			return nil, err
		}
		err = conn.Login(username, password)
		if err != nil {
			return nil, err
		}
		pool.conns <- conn
	}

	return pool, nil
}

func (p *FTPConnectionPool) GetConnection() (*ftp.ServerConn, error) {
	return <-p.conns, nil
}

func (p *FTPConnectionPool) ReleaseConnection(conn *ftp.ServerConn) {
	p.conns <- conn
}

func (p *FTPConnectionPool) Close() {
	close(p.conns)
	for conn := range p.conns {
		_ = conn.Quit()
	}
}

func (p *FTPConnectionPool) StoreFileWithPool(remotePath string, buffer []byte) error {
	conn, err := p.GetConnection()
	if err != nil {
		return err
	}
	defer p.ReleaseConnection(conn)

	data := bytes.NewBuffer(buffer)
	err = conn.Stor(remotePath, data)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	return nil
}

func main() {
	fmt.Println("hello world")
}
