# BackendGoECOReWi

## 🌈平台简介
[ ] 上传计划文件
[ ] 计划文件名解析与计划内容解析

## 🏭在线体验

## 开发流程

```shell
# 初始化模块
#go mod init <项目模块名称>

# 依赖关系处理 ,根据go.mod文件
go mod tidy
# 运行项目
go run main.go
```

## 数据库操作
```mysql
mysql -u <用户名> -p<密码>
drop DATABASE ECOReWi;
CREATE DATABASE IF NOT EXISTS ECOReWi DEFAULT CHARSET utf8 COLLATE utf8_general_ci;
use ECOReWi;
```