
limiter := NewLimiter(10, 1);
这里有两个参数：
第一个参数是 r Limit。代表每秒可以向 Token 桶中产生多少 token。Limit 实际上是 float64 的别名。
第二个参数是 b int。b 代表 Token 桶的容量大小。
那么，对于以上例子来说，其构造出的限流器含义为，其令牌桶大小为 1, 以每秒 10 个 Token 的速率向桶中放置 Token。


Wait/WaitN
func (lim *Limiter) Wait(ctx context.Context) (err error)
func (lim *Limiter) WaitN(ctx context.Context, n int) (err error)
Wait 实际上就是 WaitN(ctx,1)。

当使用 Wait 方法消费 Token 时，如果此时桶内 Token 数组不足 (小于 N)，那么 Wait 方法将会阻塞一段时间，直至 Token 满足条件。如果充足则直接返回。

这里可以看到，Wait 方法有一个 context 参数。
我们可以设置 context 的 Deadline 或者 Timeout，来决定此次 Wait 的最长时间。


参考：
https://pkg.go.dev/golang.org/x/time/rate#Limiter.Limit
https://www.cyhone.com/articles/usage-of-golang-rate/
