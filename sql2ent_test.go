package sql2ent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	sql := `
CREATE TABLE users (
  id bigint(20) unsigned NOT NULL,
  other_id bigint(20) unsigned NOT NULL,
  enum_column enum('a','b','c','d') DEFAULT NULL,
  int_column int(10) DEFAULT '0',
   create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`

	_, err := Parse(sql)
	assert.Nil(t, err)
}
