package zenoorm

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestEngine(t *testing.T) {
	engine, _ := NewEngine("sqlite3", "zeno.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	t.Logf("Exec success, %d affected\n", count)
}
