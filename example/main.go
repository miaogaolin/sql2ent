package main

import (
	"fmt"

	sql2ent "github.com/miaogaolin/sql2ent/parser"
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
