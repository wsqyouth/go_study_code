在MySQL中，你可以使用以下命令来完成你的需求：

添加新的账号密码：
CREATE USER 'newuser'@'localhost' IDENTIFIED BY 'password';
这里的newuser是你要创建的用户名，localhost表示这个用户只能从本地登录，password是用户的密码。

添加增删改查权限：
GRANT SELECT, INSERT, UPDATE, DELETE ON database_name.* TO 'newuser'@'localhost';
这里的database_name是你要授权的数据库名。这个命令会给newuser在database_name数据库上的SELECT, INSERT, UPDATE, DELETE权限。

检查对应账号的权限：
SHOW GRANTS FOR 'newuser'@'localhost';
这个命令会显示newuser的所有权限。

删除账号密码：
DROP USER 'newuser'@'localhost';
这个命令会删除newuser用户。

注意：在执行这些命令之前，你需要先登录到MySQL。另外，这些命令都需要管理员权限才能执行。在执行完GRANT命令之后，你可能需要运行FLUSH PRIVILEGES;命令来立即生效。

---

centos下mysql检查及重启
```
systemctl status mysql //检查MySQL服务状态
sudo systemctl restart mysql //重启MySQL服务：
```