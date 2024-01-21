package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
)

func main() {
	cfg := replication.BinlogSyncerConfig{
		//ServerID: 100, Flavor: "mysql", Host: "localhost", Port: 3306, User: "user", Password: "password",
		ServerID: 100, Flavor: "mysql", Host: "localhost", Port: 3306, User: "coopers", Password: "2019Youth",
	}

	syncer := replication.NewBinlogSyncer(cfg)
	streamer, err := syncer.StartSync(mysql.Position{})
	if err != nil {
		log.Fatalf("StartSync failed: %v", err)
	}

	for {
		ev, err := streamer.GetEvent(context.Background())
		if err != nil {
			log.Fatalf("GetEvent failed: %v", err)
		}

		// 处理事件
		switch e := ev.Event.(type) {
		case *replication.RowsEvent:
			// 处理写入、删除、更新事件
			// fmt.Printf("RowsEvent: Schema: %s, Table: %s, Rows: %v\n", e.Table.Schema, e.Table.Table, e.Rows)
			handleEvent(ev)
		case *replication.GTIDEvent:
			// 处理GTID事件
			uuid := string(e.SID)
			gno := e.GNO
			fmt.Printf("GTIDEvent: UUID: %s, GNO: %d\n", uuid, gno)

		}

	}
}

func handleEvent(ev *replication.BinlogEvent) error {
	switch ev.Header.EventType {
	case replication.WRITE_ROWS_EVENTv2, replication.WRITE_ROWS_EVENTv1:
		rowsEvent := ev.Event.(*replication.RowsEvent)
		// 处理写入事件
		fmt.Printf("WRITE_ROWS_EVENT: Schema: %s, Table: %s, Rows: %v\n", rowsEvent.Table.Schema, rowsEvent.Table.Table, rowsEvent.Rows)
	case replication.DELETE_ROWS_EVENTv2, replication.DELETE_ROWS_EVENTv1:
		rowsEvent := ev.Event.(*replication.RowsEvent)
		// 处理删除事件
		fmt.Printf("DELETE_ROWS_EVENT: Schema: %s, Table: %s, Rows: %v\n", rowsEvent.Table.Schema, rowsEvent.Table.Table, rowsEvent.Rows)
	case replication.UPDATE_ROWS_EVENTv2, replication.UPDATE_ROWS_EVENTv1:
		rowsEvent := ev.Event.(*replication.RowsEvent)
		// 处理更新事件
		fmt.Printf("UPDATE_ROWS_EVENT: Schema: %s, Table: %s, Rows: %v\n", rowsEvent.Table.Schema, rowsEvent.Table.Table, rowsEvent.Rows)
	}
	return nil
}
