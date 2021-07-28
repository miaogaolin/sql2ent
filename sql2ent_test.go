package sql2ent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	sql := `
	CREATE TABLE e_email_logs (
		id int(10) UNSIGNED NOT NULL,
		source varchar(10) NOT NULL DEFAULT '' COMMENT '来源',
		site_name varchar(50) NOT NULL DEFAULT '' COMMENT '站点名称',
		domain varchar(100) NOT NULL DEFAULT '' COMMENT '域名',
		sender varchar(255) NOT NULL COMMENT '发件人',
		receiver varchar(255) NOT NULL COMMENT '接收人',
		template_type_id int(10) UNSIGNED NOT NULL COMMENT '模板类型id',
		template_name varchar(50) NOT NULL COMMENT '模板名称',
		template_version_id int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '发布模板id，外键',
		language_id int(10) UNSIGNED NOT NULL COMMENT '语言id，外键',
		info varchar(512) NOT NULL DEFAULT '' COMMENT '邮件返回信息',
		message_id varchar(50) NOT NULL COMMENT '邮件返回消息id',
		request varchar(512) NOT NULL DEFAULT '' COMMENT '请求数据',
		callback_status tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '回调状态，0未发送，1失败，1成功',
		callback_count tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '回调已重试次数',
		create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
		update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='邮件日志记录' ROW_FORMAT=DYNAMIC;
`
	_, err := Parse(sql)
	assert.Nil(t, err)
}
