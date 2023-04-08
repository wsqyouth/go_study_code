### 一、操作步骤：
    1. 启动代理服务器
    go run main.go
    2. 分别启动3个缓存服务器
    $ go run server/main.go -p 8080
    start server: 8080

    $ go run server/main.go -p 8081
    start server: 8081

    $ go run server/main.go -p 8082
    start server: 8082
    3. 通过代理服务器获取key
    curl -i http://localhost:18888/key?key=1234

### 二、vscode断点调试及单元测试
打开 .vscode/settings.json

    {
        "version": "0.2.0",
        "configurations": [
            {
                "name": "Launch",
                "type": "go",
                "request": "launch",
                "mode": "auto",
                "program": "${fileDirname}",  //调试程序的路径（绝对路径） ${fileDirname} 调试当前文件所在目录下的所有文件
                "env": {},　
                "editor.fontSize": 14,
                "args": []
            }
        ],
        "tabnine.experimentalAutoImports": true,
        "go.inferGopath": false
    }

### 三、总结思考
#### 缘起：
    最近在学习redis，联想到pulsar也使用一致性哈希算法将特定的top分给固定的消费者处理,于是找到这样的一个代码进行研究。

#### 核心要点：
    * 在移除或者添加一个服务器时，能够尽可能小地改变已存在的服务请求与处理请求服务器之间的映射关系。
    * 对象key映射到服务器，从缓存对象key的位置开始，沿顺时针方向遇到的第一个服务器，便是当前对象将要缓存到的服务器；
    * 扩缩容场景:只有少部分对象需要重新分配
    * 数据偏斜&服务器性能平衡问题:因为节点分布不均匀而造成数据倾斜问题,通过引入虚拟节点来解决负载不均衡
    * 将每台物理服务器虚拟为一组虚拟服务器，将虚拟服务器放置到哈希环上，如果要确定对象的服务器，需先确定对象的虚拟服务器，再由虚拟服务器确定物理服务器；
    * 分配的虚拟节点个数越多，映射在hash环上才会越趋于均匀，节点太少的话很难看出效果；
    * 引入虚拟节点的同时也增加了新的问题，要做虚拟节点和真实节点间的映射，对象key->虚拟节点->实际节点之间的转换；
#### 代码概括:
    代理服务器是一个类似注册中心的角色,缓存服务器启动后注册，同时用户的请求到来之后根据策略选择具体的特定服务器进行业务处理。
    该代码将map,lock,切片/环使用的很巧妙,同时添加了http服务器进行实现
#### 参考：
    https://jasonkayzk.github.io/2022/02/12/%E4%B8%80%E8%87%B4%E6%80%A7Hash%E7%AE%97%E6%B3%95%E6%80%BB%E7%BB%93%E4%B8%8E%E5%BA%94%E7%94%A8/