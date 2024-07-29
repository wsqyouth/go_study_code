参考: https://segmentfault.com/a/1190000040755506

该文章中没有标明使用 host 网络链接模式，导致登录哨兵容器验证失败
```
redis-cli -h 127.0.0.1 -p 26377
SENTINEL masters

预期输出: "master-redis"
```

## 后续
哨兵模式与选举
