package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

func main() {
	server := struct {
		Address string
	}{
		Address: "localhost:6379",
	}

	redisPasswd := "" // 如果需要密码，请在这里设置

	rdb := redis.NewClient(&redis.Options{
		Addr:     server.Address,
		Password: redisPasswd,
	})
	defer rdb.Close()

	limiter := redis_rate.NewLimiter(rdb)

	// 设置每秒允许的请求数
	rate := 1
	l := redis_rate.PerSecond(rate)

	ctx := context.Background()
	key := "my_key"

	for i := 0; i < 10; i++ {
		res, err := limiter.Allow(ctx, key, l)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		// fmt.Println("allowed", res.Allowed, "remaining", res.Remaining)
		if res.Allowed > 0 {
			fmt.Printf("Request %d allowed, remaining: %d\n", i+1, res.Remaining)
		} else {
			fmt.Printf("Request %d not allowed, retry after: %v\n", i+1, res.RetryAfter)
		}

		time.Sleep(200 * time.Millisecond)
	}
}

/*
官方文档：https://pkg.go.dev/github.com/go-redis/redis_rate/v9#section-readme

golang.org/x/time/rate和go-redis/redis_rate都是优秀的限流库，但它们在使用场景和实现方式上有所不同。
golang.org/x/time/rate是一个纯Go的限流库，它使用令牌桶算法进行限流。这个库非常适合在单个进程中进行限流，因为它的所有操作都在内存中完成，不需要任何外部依赖。然而，如果你需要在分布式系统中进行限流，或者需要在多个进程或服务之间共享限流状态，那么这个库可能就不太适合了。

go-redis/redis_rate是一个基于Redis的限流库，它使用滑动窗口算法进行限流。这个库非常适合在分布式系统中进行限流，因为它的状态存储在Redis中，可以在多个进程或服务之间共享。然而，由于它依赖于Redis，所以在使用它之前，你需要确保你的系统中已经安装了Redis，并且可以正常运行。

在性能方面，由于golang.org/x/time/rate的所有操作都在内存中完成，所以它的性能通常会比go-redis/redis_rate更好。然而，go-redis/redis_rate的性能也非常优秀，对于大多数应用来说，它的性能已经足够了。

总的来说，选择哪个库主要取决于你的使用场景。如果你需要在单个进程中进行限流，那么golang.org/x/time/rate可能是更好的选择。如果你需要在分布式系统中进行限流，那么go-redis/redis_rate可能是更好的选择。


令牌桶算法(Token Bucket)

令牌桶算法是一个使用令牌作为权利单位的限流算法。在这个算法中，系统会以一定的速率向令牌桶中添加令牌。当一个请求到来时，系统会尝试从令牌桶中取出一个令牌。如果令牌桶中有足够的令牌，那么请求就被允许执行；否则，请求就被拒绝。

令牌桶算法的优点是它可以处理突发流量。如果令牌桶中有足够的令牌，那么突发的请求就可以立即被处理，而不需要等待。然而，如果令牌桶中的令牌被消耗完了，那么新的请求就需要等待新的令牌被添加到令牌桶中。

滑动窗口算法(Sliding Window)

滑动窗口算法是一个使用时间窗口作为限流单位的算法。在这个算法中，系统会记录在过去的一段时间内（例如，过去的一秒或一分钟）处理的请求的数量。当一个新的请求到来时，系统会检查在当前的滑动窗口中已经处理的请求的数量。如果这个数量已经达到了限流的阈值，那么新的请求就被拒绝；否则，新的请求就被允许执行。

滑动窗口算法的优点是它可以更精确地控制请求的速率，因为它考虑了请求的时间分布。然而，滑动窗口算法不能很好地处理突发流量，因为它只关注过去的一段时间内的请求的数量，而不关注未来的请求。

区别

令牌桶算法和滑动窗口算法的主要区别在于它们处理请求的方式。令牌桶算法关注的是请求的数量，而滑动窗口算法关注的是请求的时间分布。因此，令牌桶算法更适合处理突发流量，而滑动窗口算法更适合精确地控制请求的速率。

在限流方面的应用

在限流方面，令牌桶算法和滑动窗口算法都有广泛的应用。例如，网络路由器和防火墙通常使用令牌桶算法来控制数据包的发送速率。而Web应用和API通常使用滑动窗口算法来限制用户的请求速率，以防止系统被过度使用或者遭受DDoS攻击。
*/
