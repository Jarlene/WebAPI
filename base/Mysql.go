package base

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"strings"
	"sync"
)

type MysqlConf struct {
	engine *xorm.Engine
	table string
	cols []string
	cmd string
	cond string
}


var mysqlOnce sync.Once
var mysqlConf *MysqlConf = nil

func NewMysqlConf(datasource string) (*MysqlConf, error)  {
	var err error
	mysqlOnce.Do(func() {
		engine, e := xorm.NewEngine("mysql", datasource)
		err = e
		if e == nil {
			mysqlConf = new(MysqlConf)
			mysqlConf.engine = engine
		}
	})
	return mysqlConf, err
}


func (this *MysqlConf) Table(t string) *MysqlConf {
	this.table = t
	return this
}

func (this *MysqlConf) Commend(cmd string) *MysqlConf {
	this.cmd = cmd
	return this
}

func (this *MysqlConf) Where(cond string) *MysqlConf {
	this.cond = cond
	return this
}

func (this *MysqlConf) Args() *MysqlConf {

	return this
}

func (this *MysqlConf) Cols(cols ...string) *MysqlConf  {
	this.cols = cols
	return this
}

func (this *MysqlConf) Send() (interface{}, error)  {
	var sql string
	if this.cmd == "select" {
		sql = this.cmd + " " + strings.Join(this.cols, ",") + " from " + this.table
		return this.engine.Query(sql)
	} else if this.cmd == "update" {

	} else if this.cmd == "delete" {

	} else if this.cmd == "insert" {

	} else  {

	}
	return nil, nil
}

func (this *MysqlConf) Sql(sql string) (interface{}, error) {
	res, err:= this.engine.Query(sql)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	return res, nil
}

func (this *MysqlConf) Close()  {
	this.engine.Close()
	Default().Remove("mysql")
}

func (this *MysqlConf) Insert(sql string) (interface{}, error) {
	res, err:= this.engine.Exec(sql)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return id, nil
}