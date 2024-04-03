package datasource

import "context"

// GetData 从数据源获取数据
func GetData(ctx context.Context) (string, error) {
	return "srcdata", nil
}

// DataSource 数据源实现
type DataSource struct {
}

// NewDataSource 获取数据
func NewDataSource(ctx context.Context) *DataSource {
	return &DataSource{}
}

func (ds *DataSource) GetData(ctx context.Context) (string, error) {
	return "srcdata", nil
}
