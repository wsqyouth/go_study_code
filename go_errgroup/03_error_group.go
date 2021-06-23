package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

func checkGoroutineErr(errCtx context.Context) error {
	select {
	case <-errCtx.Done():
		return errCtx.Err()
	default:
		return nil
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, errCtx := errgroup.WithContext(ctx)

	for i := 0; i < 3; i++ {
		index := i
		group.Go(func() error {
			fmt.Println("index=", index)
			if index == 0 {
				fmt.Println("index == 0, end!")
			} else if index == 1 {
				fmt.Println("index == 1, start...")
				cancel()
				fmt.Println("inde == 1, has error!")
			} else if index == 2 {
				fmt.Println("index == 2, start...")
				time.Sleep(time.Second * 3)
				if err := checkGoroutineErr(errCtx); err != nil {
					return err
				}
				fmt.Println("index == 2, has done!")
			}
			return nil
		})
	}

	err := group.Wait()
	if err != nil {
		fmt.Println("Get error: ", err)
	} else {
		fmt.Println("All Done!")
	}
}
