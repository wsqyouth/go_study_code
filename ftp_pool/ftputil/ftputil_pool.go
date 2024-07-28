package ftputil

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/jlaffaye/ftp"
	"github.com/pkg/errors"
)

var (
	ErrPoolClosed = errors.New("pool is closed")
	ErrPoolEmpty  = errors.New("pool is empty")
)

// FtpConnectionPool is a pool of FTP connections
type FtpConnectionPool struct {
	conns       chan *ftp.ServerConn
	mu          sync.Mutex
	activeConns int32 // current number of active connections
	capacity    int32 // maximum number of connections
	closed      bool  // pool is closed or not

	server   string
	username string
	password string
}

// NewFtpConnectionPool creates a new FTP connection pool
func NewFtpConnectionPool(capacity int, server, username, password string) (*FtpConnectionPool, error) {
	if capacity <= 0 {
		return nil, errors.New("invalid capacity settings")
	}

	pool := &FtpConnectionPool{
		conns:    make(chan *ftp.ServerConn, capacity),
		capacity: int32(capacity),
		server:   server,
		username: username,
		password: password,
	}

	for i := 0; i < capacity; i++ {
		conn, err := pool.NewFtpConnection()
		if err != nil {
			return nil, err
		}
		pool.conns <- conn
		pool.incrementActiveConnections()
	}
	return pool, nil
}

func (p *FtpConnectionPool) incrementActiveConnections() {
	atomic.AddInt32(&p.activeConns, 1)
}

func (p *FtpConnectionPool) decrementActiveConnections() {
	atomic.AddInt32(&p.activeConns, -1)
}

// NewFtpConnection creates FTP connection
func (p *FtpConnectionPool) NewFtpConnection() (*ftp.ServerConn, error) {
	newConn, err := ftp.Dial(p.server, ftp.DialWithDisabledUTF8(true))
	if err != nil {
		return nil, fmt.Errorf("failed to dial FTP server: %w", err)
	}

	err = newConn.Login(p.username, p.password)
	if err != nil {
		if closeErr := newConn.Quit(); closeErr != nil {
			return nil, fmt.Errorf("failed to close connection: %w", closeErr)
		}
		return nil, fmt.Errorf("failed to login to FTP server: %w", err)
	}

	return newConn, nil
}

// GetFtpConnection Get an FTP connection from the pool, if invalid open a new one
func (p *FtpConnectionPool) GetFtpConnection() (*ftp.ServerConn, error) {
	if p.conns == nil {
		return nil, ErrPoolClosed
	}

	conn := <-p.conns
	if p.checkFtpConnectionValid(conn) {
		return conn, nil
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	if err := p.QuitFtpConnection(conn); err != nil {
		return nil, err
	}
	p.decrementActiveConnections()
	if p.activeConns < p.capacity {
		newConn, err := p.NewFtpConnection()
		if err != nil {
			return nil, err
		}
		p.incrementActiveConnections()
		return newConn, nil
	}
	return nil, errors.New("unable to retrieve or create FTP connection")
}

// PutFtpConnection Put an FTP connection back to the pool, if full close the connection
func (p *FtpConnectionPool) PutFtpConnection(conn *ftp.ServerConn) error {
	if p.conns == nil || p.closed {
		if quitErr := conn.Quit(); quitErr != nil {
			return quitErr
		}
		return nil
	}

	if len(p.conns) < cap(p.conns) {
		p.conns <- conn
		return nil
	}
	if quitErr := conn.Quit(); quitErr != nil {
		return errors.New("failed to close connection")
	}

	return nil
}

func (p *FtpConnectionPool) Close() error {
	if p.closed {
		return nil
	}

	p.closed = true
	close(p.conns)

	var err error
	for conn := range p.conns {
		if quitErr := conn.Quit(); quitErr != nil {
			err = quitErr
		}
		fmt.Println("ftpPoolClient conn Close", err)
	}

	return err
}

func (p *FtpConnectionPool) checkFtpConnectionValid(conn *ftp.ServerConn) bool {
	if conn == nil {
		return false
	}
	return conn.NoOp() == nil
}

// QuitFtpConnection Close the existing connection if it's not nil
func (p *FtpConnectionPool) QuitFtpConnection(conn *ftp.ServerConn) error {
	if conn == nil {
		return nil
	}
	if err := conn.Quit(); err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}
	return nil
}

// ReportStats return the number of available and active connections
func (p *FtpConnectionPool) ReportStats() (availableConns int32, activeConns int32) {
	return int32(len(p.conns)), p.activeConns
}
