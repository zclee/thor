package api

import (
	"github.com/vechain/thor/logdb"
)

//LogInterface for query logs
type LogInterface struct {
	ldb *logdb.LogDB
}

//NewLogInterface new LogInterface
func NewLogInterface(ldb *logdb.LogDB) *LogInterface {
	return &LogInterface{
		ldb,
	}
}

//Filter query logs with option
func (li *LogInterface) Filter(option *logdb.FilterOption) ([]*logdb.Log, error) {
	logs, err := li.ldb.Filter(option)
	if err != nil {
		return nil, err
	}
	return logs, nil
}