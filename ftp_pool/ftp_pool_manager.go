package main

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

type FtpPoolManager struct {
	PoolCapacity int
	clientMap    map[string]*FtpPoolClient
	mu           sync.Mutex
}

var _ftpPoolManager *FtpPoolManager

// Init initialize the FTP connection pool
func Init() error {
	_ftpPoolManager = &FtpPoolManager{
		clientMap:    make(map[string]*FtpPoolClient),
		PoolCapacity: 5,
	}

	return nil
}

// Get ftp pool get the FTP connection pool manager
func Get() *FtpPoolManager {
	return _ftpPoolManager
}

// GetFtpClient according to user attribute get the corresponding connection pool
func (fp *FtpPoolManager) GetFtpClient(serverAddress, username, password string) (*FtpPoolClient, error) {
	fp.mu.Lock()
	defer fp.mu.Unlock()

	clientKey := fmt.Sprintf("%s-%s", serverAddress, username)
	ftpPoolClient, ok := fp.clientMap[clientKey]
	// 测试时打印属性观察
	fmt.Printf("find_map map :%v, clientKey:%v,exist:%v\n", fp.clientMap, clientKey, ok)
	if ok {
		return ftpPoolClient, nil
	}
	// if ftpPoolClient, ok := fp.clientMap[clientKey]; ok {
	// 	return ftpPoolClient, nil
	// }

	ftpPoolClient, err := NewFtpPoolClient(fp.PoolCapacity, serverAddress, username, password)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// 测试时打印属性观察
	availableConns, activeConns := ftpPoolClient.pool.ReportStats()
	fmt.Printf("new_ftp_pool AvailConnections: %v,activeConns:%v,Capacity:%v\n", availableConns, activeConns, fp.PoolCapacity)

	fp.clientMap[clientKey] = ftpPoolClient
	return ftpPoolClient, nil
}

// Close ftp pool close the FTP connection pool
func (fp *FtpPoolManager) Close() {
	for _, ftpPoolClient := range fp.clientMap {
		if ftpPoolClient != nil {
			ftpPoolClient.Close()
		}
	}
}

func main() {
	Init()
	ftppoolmanager := Get()

	server := "localhost:2100"
	username := "jovy"
	password := "123456"
	ftppoolclient, err := ftppoolmanager.GetFtpClient(server, username, password)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ftppoolclient.StoreFile("test.txt", []byte("Hello, World!"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ftppoolclient)
	ftppoolmanager.Close()
}
