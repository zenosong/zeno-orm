package clause

import (
	"fmt"
	"strings"
)

// 查询语句生成器
// SELECT col1, col2, ...
// FROM table_name
// WHERE [ conditions ]
// GROUP BY col1
// HAVING [ conditions ]

type generator func(values ...interface{}) (string, []interface{})

var generators map[Type]generator

func init() {
	generators = make(map[Type]generator)
	generators[INSERT] = _insert
}

func genBindVars(num int) string {
	var vars []string
	for i := 0; i < num; i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ", ")
}

func _insert(values ...interface{}) (string, []interface{}) {
	tableName := values[0]
	// .(TYPE) 限制访问的数据类型
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("INSERT INTO %s (%v)", tableName, fields), []interface{}{}
}
