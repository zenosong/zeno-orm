package zenoorm

import (
	"database/sql"
	"zenoorm/dialect"
	"zenoorm/log"
	"zenoorm/session"
)

// Engine 交互前的准备工作（比如连接/测试数据库），交互后的收尾工作（关闭连接）
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	// 测试连接是否成功
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	// 获取对应数据库的差异处理器
	dialect, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s not found", driver)
		return
	}
	e = &Engine{db: db, dialect: dialect}
	log.Info("Connect database success")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error(err)
	}
	log.Info("Close database success")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}
