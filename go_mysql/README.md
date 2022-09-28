 go test -bench=. -benchmem
 test  -gcflags=-l -run Test_mysqlDBPool_Get
 go test  -gcflags=-l -run Test_mysqlDBPool_Open
