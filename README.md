### mysql_proxy

用于代理真实 mysql 服务器,并打印传输到真实 mysql 服务的 sql 语句

### 配置文件

代理服务器配置

```yaml 
ProxyMysql:
  Host: 127.0.0.1
  Port: 4406
  User: root
  Pwd: 123456
```

真实服务器配置

```yaml
RealMysql:
  Host: 127.0.0.1
  Port: 3306
  User: root
  Pwd: 123456
  Db: xxx
```

### 客户端连接

客户端连接到 `127.0.0.1:4406` 就可以像连接 `127.0.0.1:3306` 一样使用真实服务器

### 用途

使用 mysql 代理,可以让你使用第三方软件连接数据库时,无需看第三方软件的源码是怎么实现的,而只需要看到别人的 SQL 是如何运行的