package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/time/rate"
)

// objConverter 转换器接口
type objConverter interface {
	Convert(ctx context.Context) (err error)
}

type Change struct {
	converter  objConverter
	changeFunc func(ctx context.Context) error
}

func applyChanges(ctx context.Context, changes ...*Change) error {
	for _, each := range changes {
		if err := each.converter.Convert(ctx); err != nil {
			return errors.Wrap(err, "converter.Convert error")
		}
		mpLimiter, ok := GetLimiter(ctx)
		if ok {
			beforeLimiterTime := time.Now()
			err := mpLimiter.Wait(ctx)
			if err != nil {
				return err
			}
			fmt.Printf("mpLimiter wait: %s\n", time.Now().Sub(beforeLimiterTime).String())
		}

		if err := each.changeFunc(ctx); err != nil {
			return errors.Wrap(err, "changeFunc error")
		}
	}
	return nil
}

// FirstLevelConverter
type FirstLevelConverter struct {
	srcData string
	dstData *string
}

func newFirstLevelConverter(srcData string, dstData *string) *FirstLevelConverter {
	return &FirstLevelConverter{
		srcData: srcData,
		dstData: dstData,
	}
}

func (c *FirstLevelConverter) Convert(ctx context.Context) (err error) {
	srcData, dstData := c.srcData, c.dstData
	if srcData == "" {
		return errors.New("srcData is empty")
	}
	*dstData = "FirstLevelConverter has processed: \nTime:" + srcData + time.Now().String()
	return nil
}

// generateFirstLevelChange 产生change
func generateFirstLevelChange(ctx context.Context, srcData string, dstData *string) (change *Change, err error) {
	// 创建一个FirstLevelConverter实例
	converter := newFirstLevelConverter(srcData, dstData)

	// 创建一个Change实例
	change = &Change{
		converter: converter,
		changeFunc: func(ctx context.Context) error {
			fmt.Println("firstLevel change...")
			return nil
		},
	}
	return change, nil
}

type mpLimiterKey struct{}

func GetLimiter(ctx context.Context) (*rate.Limiter, bool) {
	limiter, ok := ctx.Value(mpLimiterKey{}).(*rate.Limiter)
	return limiter, ok
}

func main() {
	// 创建一个带有限速器的上下文
	ctx := context.WithValue(context.Background(), mpLimiterKey{}, rate.NewLimiter(rate.Every(time.Second), 1))

	// 示例数据
	srcData := "example data"
	var dstData string
	change, err := generateFirstLevelChange(ctx, srcData, &dstData)
	if err != nil {
		fmt.Printf("Generating level change error: %v\n", err)
	}
	// 应用更改
	err = applyChanges(ctx, change)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success! Converted data: %s\n", dstData)
	}
}

/*
包括一个简单的main函数来演示如何使用applyChanges函数。这个代码的优点在于它将业务逻辑与框架解耦，使得代码更易于维护和扩展。
具体业务实现和框架解耦：通过使用objConverter接口和Change结构体，我们可以将具体的业务逻辑（如FirstLevelConverter）与框架代码（如applyChanges）分离。这使得我们可以在不修改框架代码的情况下添加新的业务逻辑。
支持可扩展：由于业务逻辑和框架代码是分离的，我们可以轻松地为不同的业务场景创建新的转换器实现。只需实现objConverter接口并将其传递给applyChanges函数即可。
灵活的限速器：通过使用context和rate.Limiter，我们可以为不同的业务场景提供灵活的限速设置。这有助于防止过载和提高系统的稳定性。
*/
