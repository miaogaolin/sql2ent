# sql2ent
该项目提供 sql 语句转化为 `entgo schema` 代码的命令行工具, 以提高工作效率。

[前往学习entgo](https://entgo.io)

## 功能

### 已完成
1. 使用命令行批量转化
2. 支持 mysql

### 计划
1. 支持更多的数据库，例如：MariaDB、SQLite、PostgreSQL。
2. 读取数据库，批量生成 `schema` 文件。
3. 等等...

## 快速开始

### 第一步：安装 `sql2ent`
```shell
# Go 1.15 或更低版本
go get -u github.com/miaogaolin/excel-proc

# Go 1.16 或更高版本
go install github.com/miaogaolin/excel-proc@latest
```

### 第二步：运行命令
```shell
sql2ent mysql ddl -src "./sql/*.sql" -dir "./ent/schema"
```
说明：
* -src: 输入 sql 路径，可模糊匹配
* -dir: 输出目录，默认 `./ent/schema`



## 参与开源

1. 点击 Fork
2. 提交自己的代码到 Fork 的仓库中
3. Pull Request 将自己的代码合并