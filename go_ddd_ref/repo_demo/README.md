
对user.go进行生产mock文件
```
mockgen -destination=repository/mock/user_mock.go -package=mock_repository repo_demo/repository IUserRepo
tree ./ -L 3   
```
