package main

import (
	"context"
	"errors"
	"fmt"
)

func main() {
	userObj := userChecker{1001, "zhangsan", 100}
	checkers := []Checker{NewUserChecker(userObj.id, userObj.name, userObj.age)}
	processors := []Processor{NewUserProcessor(userObj.id, userObj.name, userObj.age)}
	if err := NewPreProcessor(checkers, processors).Do(context.Background()); err != nil {
		fmt.Printf("process falied. err:%v", err)
		return
	}
	fmt.Println("Congratulations. countdown: %+V", userObj)
}

type userChecker struct {
	id   uint64
	name string
	age  uint64
}

func NewUserChecker(
	id uint64,
	name string,
	age uint64,
) *userChecker {
	return &userChecker{
		id:   id,
		name: name,
		age:  age,
	}
}

// Check 实现校验接口
func (c *userChecker) Check(ctx context.Context) error {
	if c.id < 10 {
		return errors.New("id must lt 10")
	}
	if c.name == "" {
		return errors.New("name must not empty")
	}
	if c.age < 18 {
		return errors.New("age must lt 18")
	}
	return nil
}

type userProcessor struct {
	id   uint64
	name string
	age  uint64
}

func NewUserProcessor(
	id uint64,
	name string,
	age uint64,
) *userProcessor {
	return &userProcessor{
		id:   id,
		name: name,
		age:  age,
	}
}

// Process 处理接口
func (c *userProcessor) Process(ctx context.Context) error {
	c.id = c.id + 1
	c.name = "Super" + c.name
	c.age = c.age + 10
	return nil
}
