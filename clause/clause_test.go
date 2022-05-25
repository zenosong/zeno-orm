package clause

import (
	"reflect"
	"testing"
)

func testSelect(t *testing.T) {
	var clause Clause
	clause.Set(LIMIT, 3)
	clause.Set(WHERE, "Name = ?", "Tom")
	clause.Set(SELECT, "User", []string{"*"})
	clause.Set(ORDERBY, "Age ASC")
	sql, vars := clause.Build(SELECT, WHERE, ORDERBY, LIMIT)
	t.Log(sql, vars)

	if sql != "SELECT * FROM User SELECT Name = ? ORDER BY Age ASC LIMIT ?" {
		t.Fatal("failed to Build SQL")
	}

	if !reflect.DeepEqual(vars, []interface{}{"Tom", 3}) {
		t.Fatal("failed to Build sqlVars")
	}
}

func TestClause_Build(t *testing.T) {
	t.Run("select", func(t *testing.T) {
		testSelect(t)
	})
}
