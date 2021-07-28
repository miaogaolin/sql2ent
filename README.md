# sql2ent
该项目提供 sql 语句转化为 `entgo schema` 代码的快速工具, 以提高工作效率。

[前往学习entgo](https://entgo.io)

## 功能

### 已完成

1. 支持 mysql 的 `create sql`语句转化为`entgo schema` 代码
2. 第三包引入调用 `sql2ent.Parse(string)`

### 计划
1. 支持更多的数据库，例如：MariaDB、SQLite、PostgreSQL。
2. 提供命令行工具。
3. 读取数据库，批量生成 `schema` 文件。
等等...

## 快速开始

### 第一步：初始化 `ent`
```shell
ent init Users
```
注：如果已经初始化过就省略这一步骤


### 第二步：生成 `schema`

[前往在线工具](https://www.printlove.cn/tools/sql2ent)

拉取最新代码：
```shell
go get -u github.com/miaogaolin/sql2ent
```

拷贝以下代码：
```go
package main

import (
	"fmt"

	"github.com/miaogaolin/sql2ent"
)

func main() {
	sql := `
CREATE TABLE users (
 id int(10) unsigned NOT NULL AUTO_INCREMENT,
 name varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
 email varchar(50) NOT NULL COMMENT '邮箱',
 mobile char(11) NOT NULL DEFAULT '' COMMENT '手机号',
 status tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态， 0禁用，1启用',
 login_type tinyint(1) NOT NULL COMMENT '登录类型，0钉钉，1密码',
 user_id int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新用户id',
 is_system tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '系统最高管理员，1是，0否',
 login_ip varchar(20) NOT NULL DEFAULT '' COMMENT '登录ip',
 token varchar(255) NOT NULL DEFAULT '' COMMENT 'token',
 create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
 update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
 PRIMARY KEY (id) USING BTREE,
 UNIQUE KEY email_unique (email) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户表'
`
	res, err := sql2ent.Parse(sql)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

```
`go run main.go` 运行如上代码，将结果复制到 `ent/schema/Users.go` 文件中。

### 第三步：生成代码
```shell
go generate ./ent
```

## 参与开源

1. 点击 Fork
2. 提交自己的代码到 Fork 的仓库中
3. Pull Request 将自己的代码合并

## 联系我
![](./Wechat.jpeg)