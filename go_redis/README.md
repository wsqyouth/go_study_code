检查Redis是否安装成功可以执行以下命令：
```
redis-cli ping
sudo netstat -tlnp | grep 6379  
```
如果Redis已经成功安装，则会返回“PONG”作为响应。
要重启Redis，可以执行以下命令：
1. 使用systemctl命令重启Redis：
```
sudo systemctl restart redis
```
2. 或者使用service命令重启Redis：
```
sudo service redis restart
```