package main

import "sync"

type TableList struct {
	tableList        []*Table
	tableListCounter int
	tableMutex       sync.Mutex
}

func NewTableList() *TableList {
	tableListCounter := 0
	tableList := make([]*Table, 0)
	for i := 0; i < tableNumber; i++ {
		tableList = append(tableList, NewTable(tableListCounter, 0, nil))
		tableListCounter++
	}
	return &TableList{tableList, tableListCounter, sync.Mutex{}} //memory adress
}

func (tl *TableList) start() {
	for _, table := range tl.tableList {
		go table.waitCustomers()
	}
}

func (tl *TableList) deliver(delivery *Delivery, now int64) {
	tl.tableMutex.Lock()
	defer tl.tableMutex.Unlock()
	tl.tableList[delivery.TableId].deliver(delivery, now)
}

func (tl *TableList) serveTable(waiter *Waiter) *Order {
	tl.tableMutex.Lock()
	defer tl.tableMutex.Unlock()

	for _, table := range tl.tableList {
		if table.status == 1 {
			return table.serve(waiter)
		}
	}
	return nil
}
