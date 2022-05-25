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
	generators[VALUES] = _values
	generators[SELECT] = _select
	generators[LIMIT] = _limit
	generators[WHERE] = _where
	generators[ORDERBY] = _orderBy
}

func genBindVars(num int) string {
	var vars []string
	for i := 0; i < num; i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ", ")
}

// INSERT INTO $tableName ($fields)
func _insert(values ...interface{}) (string, []interface{}) {
	tableName := values[0]
	// .(TYPE) 限制访问的数据类型
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("INSERT INTO %s (%v)", tableName, fields), []interface{}{}
}

// 组装 VALUES 及其值
// VALUES ($v1), ($v2), ...
func _values(values ...interface{}) (string, []interface{}) {
	var bindStr string
	var vars []interface{}
	var sql strings.Builder
	sql.WriteString("VALUES ")
	valuesLen := len(values)
	for i, value := range values {
		v := value.([]interface{})
		if bindStr == "" {
			bindStr = genBindVars(len(v))
		}
		sql.WriteString(fmt.Sprintf("(%v)", bindStr))
		if i+1 != valuesLen {
			sql.WriteString(", ")
		}
		vars = append(vars, v...)
	}
	return sql.String(), vars
}

// SELECT $fields FROM $tableName
func _select(values ...interface{}) (string, []interface{}) {
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("SELECT %v FROM %s", fields, tableName), nil
}

// LIMIT $num
func _limit(values ...interface{}) (string, []interface{}) {
	return "LIMIT ?", values
}

// WHERE $desc
func _where(values ...interface{}) (string, []interface{}) {
	desc, vars := values[0], values[1:]
	return fmt.Sprintf("SELECT %s", desc), vars
}

// ORDER BY %s
func _orderBy(values ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("ORDER BY %s", values[0]), []interface{}{}
}
