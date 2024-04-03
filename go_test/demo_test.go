package main

import (
	"context"
	"go_test/datasource"
	"os"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
)

// 命名规则：Test+测试函数名
func TestIsInUint32(t *testing.T) {
	// 用于断言
	ass := assert.New(t)
	// 生成测试用例
	var tests = []struct {
		a     uint32
		array []uint32
		want  bool
	}{
		{uint32(1), []uint32{1, 2, 3}, true},
		{uint32(4), []uint32{1, 2, 3}, false},
	}
	// 遍历所有测试用例
	for _, test := range tests {
		// 断言 test.want 跟函数IsInUint32(test.a, test.array)的返回值相等，若不相等则输出"not ok"
		ass.Equal(test.want, IsInUint32(test.a, test.array), "not ok")
	}
}

// mock 函数
func TestInit(t *testing.T) {
	// 将函数os.LookupEnv,替换一个固定返回"1",true的匿名函数
	patches := gomonkey.ApplyFunc(os.LookupEnv, func(a string) (string, bool) {
		return "1", true
	})
	// patches.
	defer patches.Reset()
	Init()
}

// mock 方法
func TestStackLen(t *testing.T) {
	ass := assert.New(t)
	// 实例化一个stack对象
	var st = NewStack()
	// 将该对象的Len函数替换掉；
	// 注意，需要使用反射获取该对象的type，用字符串表明函数名称
	patcher := gomonkey.ApplyMethod(reflect.TypeOf(st), "Len", func(_ *Stack) int { return 10 })
	defer patcher.Reset()
	// 如果没有mock，st.Len()返回0，mock后返回10
	ass.Equal(10, st.Len(), "ok")
	// 如果执行失败，请检查是否关闭开启了内联函数导致
}

// mock 函数序列
// 测试函数
func TestIsTagSeq(t *testing.T) {
	ass := assert.New(t)
	// 生成一个输出序列
	outputs := []gomonkey.OutputCell{
		{Values: gomonkey.Params{true}},
		{Values: gomonkey.Params{false}},
	}
	// 将IsTag函数的输出mock成outputs
	// 即第一次调用IsTag时直接返回true，第二次则直接返回false
	patches := gomonkey.ApplyFuncSeq(IsTag, outputs)
	defer patches.Reset()
	ass.Equal(true, IsTag("123"), "not ok")
	ass.Equal(false, IsTag("123"), "not ok")
}

func TestGetDataFromDB(t *testing.T) {

	tests := []struct {
		name string

		want    string
		wantErr bool
	}{
		// 测试gomonky 包调用函数
		{
			name:    "test1",
			want:    "123",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var p = gomonkey.ApplyFunc(datasource.GetData, func(ctx context.Context) (string, error) {
				return "123", nil
			})
			defer p.Reset()

			got, err := GetDataFromDB(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataFromDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDataFromDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDataFromDBStruct(t *testing.T) {

	tests := []struct {
		name string

		want    string
		wantErr bool
	}{
		// 测试gomonky 包调用方法
		{
			name:    "test1",
			want:    "123",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockDs := datasource.NewMockDataSource(ctx)
			var p = gomonkey.ApplyFunc(datasource.GetData, func(ctx context.Context) (string, error) {
				return "123", nil
			})
			defer p.Reset()

			got, err := GetDataFromDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataFromDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDataFromDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
