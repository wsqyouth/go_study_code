package filter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func echoHandle(ctx context.Context, req interface{}, rsp interface{}) error {
	preq := req.(*string)
	prsp := rsp.(*string)
	*prsp = *preq
	rsp = prsp

	return nil
}

func TestNoopFilter(t *testing.T) {

	req := "echo"
	rsp := ""
	err := NoopFilter(context.Background(), &req, &rsp, echoHandle)
	assert.Nil(t, err)
	assert.Equal(t, rsp, req)
}

func TestFilterChain_Handle(t *testing.T) {
	req := "echo"
	rsp := ""
	//noopFilter
	{
		fc := FilterChain{}
		err := fc.Handle(context.Background(), &req, &rsp, echoHandle)
		assert.Nil(t, err)
		assert.Equal(t, rsp, req)
	}

	//oneFilter
	{
		fc := FilterChain{NoopFilter}
		err := fc.Handle(context.Background(), &req, &rsp, echoHandle)
		assert.Nil(t, err)
		assert.Equal(t, rsp, req)
	}

	// multiFilter
	{
		fc := FilterChain{NoopFilter, NoopFilter, NoopFilter}
		err := fc.Handle(context.Background(), &req, &rsp, echoHandle)
		assert.Nil(t, err)
		assert.Equal(t, rsp, req)
	}
}

func TestGetClient(t *testing.T) {
	Register("noop", NoopFilter, NoopFilter)
	f := GetClient("noop")
	assert.NotNil(t, f)
}

func TestGetServer(t *testing.T) {
	Register("noop", NoopFilter, NoopFilter)
	f := GetServer("noop")
	assert.NotNil(t, f)
}
