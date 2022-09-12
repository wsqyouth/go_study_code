package main

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

type ReqParam struct {
	id   int64
	name string
}

// 使用反射进行校验
func checkParam(req ReqParam) error {
	var paramMap = map[interface{}]string{
		req.id:   "id",
		req.name: "name",
	}
	for paramValue, paramName := range paramMap {
		if err := validatValEmpty(paramValue, paramName); err != nil {
			return err
		}
	}
	return nil
}

func validatValEmpty(val interface{}, name string) error {
	if reflect.ValueOf(val).IsZero() {
		return errors.Errorf("val is empty when name: %v", name)
	}
	return nil
}

// 简单校验
func checkParamSample(req ReqParam) error {
	if req.id == 0 || req.name == "" {
		return errors.Errorf("val is empty when req: %v", req)
	}
	return nil
}
func main() {
	var req ReqParam
	req.id = 23
	//req.name = "hello"
	if err := checkParamSample(req); err != nil {
		fmt.Printf("checkParamSample err: %v. req: %+v\n", err, req)
	}
	if err := checkParam(req); err != nil {
		fmt.Printf("checkParam err: %v. req: %+v\n", err, req)
	}
}
