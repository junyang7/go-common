package _sql

import (
	"database/sql"
	"fmt"
	"git.ziji.fun/junyang/go-common/_as"
	"git.ziji.fun/junyang/go-common/_dict"
	"git.ziji.fun/junyang/go-common/_is"
	"git.ziji.fun/junyang/go-common/_list"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"sync"
	"time"
)

type Machine struct {
	Driver    string `json:"driver"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	Database  string `json:"database"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Charset   string `json:"charset"`
	Collation string `json:"collation"`
}
type Master struct {
	Count       int        `json:"count"`
	MachineList []*Machine `json:"machine_list"`
}
type Slaver struct {
	Count       int        `json:"count"`
	MachineList []*Machine `json:"machine_list"`
}
type Cluster struct {
	Master *Master
	Slaver *Slaver
}
type Business struct {
	Count       int        `json:"count"`
	ClusterList []*Cluster `json:"cluster_list"`
}
type Sql struct {
	machineMaster *Machine
	machineSlaver *Machine
	driver        string
	table         string
	sql           string
	parameter     []interface{}
	master        bool
	where         string
	field         string
	index         string
	offset        int
	limit         int
	order         string
	group         string
	rowList       []map[string]interface{}
	transaction   bool
	tx            *sql.Tx
	dsn           string
}

var pool = map[string]*sql.DB{}
var m = sync.RWMutex{}

func New() *Sql {
	return &Sql{
		parameter: []interface{}{},
		rowList:   []map[string]interface{}{},
	}
}
func (this *Sql) Machine(machine *Machine) *Sql {
	this.machineMaster = machine
	this.machineSlaver = machine
	return this
}
func (this *Sql) MachineMaster(machineMaster *Machine) *Sql {
	this.machineMaster = machineMaster
	return this
}
func (this *Sql) MachineSlaver(machineSlaver *Machine) *Sql {
	this.machineSlaver = machineSlaver
	return this
}
func (this *Sql) Driver(driver string) *Sql {
	this.driver = driver
	return this
}
func (this *Sql) Dsn(dsn string) *Sql {
	this.dsn = dsn
	return this
}
func (this *Sql) Table(table string) *Sql {
	this.table = table
	return this
}
func (this *Sql) Sql(sql string) *Sql {
	this.sql = sql
	return this
}
func (this *Sql) Parameter(parameter ...interface{}) *Sql {
	this.parameter = parameter
	return this
}
func (this *Sql) Master(master bool) *Sql {
	this.master = master
	return this
}
func (this *Sql) Where(where string) *Sql {
	this.where = where
	return this
}
func (this *Sql) Field(field string) *Sql {
	this.field = field
	return this
}
func (this *Sql) Index(index string) *Sql {
	this.index = index
	return this
}
func (this *Sql) Offset(offset int) *Sql {
	this.offset = offset
	return this
}
func (this *Sql) Limit(limit int) *Sql {
	this.limit = limit
	return this
}
func (this *Sql) Order(order string) *Sql {
	this.order = order
	return this
}
func (this *Sql) Group(group string) *Sql {
	this.group = group
	return this
}
func (this *Sql) RowList(rowList []map[string]interface{}) *Sql {
	this.rowList = rowList
	return this
}
func (this *Sql) Row(row map[string]interface{}) *Sql {
	this.rowList = []map[string]interface{}{row}
	return this
}
func (this *Sql) Tx(tx *sql.Tx) *Sql {
	this.tx = tx
	return this
}
func (this *Sql) BeginTransaction() *sql.Tx {
	this.Master(true)
	tx, err := this.getPool().Begin()
	if nil != err {
		panic(err)
	}
	return tx
}
func (this *Sql) Commit() {
	if err := this.tx.Commit(); nil != err {
		panic(err)
	}
}
func (this *Sql) Rollback() {
	if err := this.tx.Rollback(); nil != err {
		panic(err)
	}
}
func (this *Sql) AddList() string {
	this.Master(true)
	this.buildAddList()
	res := this.execute()
	lastInsertId, err := res.LastInsertId()
	if nil != err {
		panic(err)
	}
	return _as.String(lastInsertId)
}
func (this *Sql) Add() string {
	this.Master(true)
	this.buildAddList()
	res := this.execute()
	lastInsertId, err := res.LastInsertId()
	if nil != err {
		panic(err)
	}
	return _as.String(lastInsertId)
}
func (this *Sql) Del() string {
	this.Master(true)
	this.buildDel()
	res := this.execute()
	effectedRowCount, err := res.RowsAffected()
	if nil != err {
		panic(err)
	}
	return _as.String(effectedRowCount)
}
func (this *Sql) Set() string {
	this.Master(true)
	this.buildSet()
	res := this.execute()
	effectedRowCount, err := res.RowsAffected()
	if nil != err {
		panic(err)
	}
	return _as.String(effectedRowCount)
}
func (this *Sql) GetList() []map[string]string {
	this.buildGetList()
	return this.query()
}
func (this *Sql) Get() map[string]string {
	this.buildGetList()
	res := this.query()
	if len(res) > 0 {
		return this.query()[0]
	} else {
		return map[string]string{}
	}
}
func (this *Sql) Count() string {
	this.buildCount()
	return this.query()[0]["c"]
}
func (this *Sql) Query() []map[string]string {
	return this.query()
}
func (this *Sql) Execute() string {
	this.Master(true)
	sql := this.getSql()
	res := this.execute()
	if strings.HasPrefix(strings.ToLower(sql), "insert") {
		lastInsertId, err := res.LastInsertId()
		if nil != err {
			panic(err)
		}
		return _as.String(lastInsertId)
	}
	effectedRowCount, err := res.RowsAffected()
	if nil != err {
		panic(err)
	}
	return _as.String(effectedRowCount)
}
func (this *Sql) getMachine() *Machine {
	if this.getMaster() {
		return this.machineMaster
	}
	return this.machineSlaver
}
func (this *Sql) getSql() string {
	return strings.TrimSpace(this.sql)
}
func (this *Sql) getParameter() []interface{} {
	return this.parameter
}
func (this *Sql) getMaster() bool {
	return this.master
}
func (this *Sql) getTable() string {
	return this.table
}
func (this *Sql) getWhere() string {
	if "" == this.where {
		return ""
	} else {
		return " WHERE " + this.where
	}
}
func (this *Sql) getField() string {
	if "" == this.field {
		return "*"
	} else {
		return this.field
	}
}
func (this *Sql) getIndex() string {
	if "" == this.index {
		return ""
	} else {
		return " FORCE INDEX (" + this.index + ")"
	}
}
func (this *Sql) getOffset() int {
	return this.offset
}
func (this *Sql) getLimit() string {
	if 0 == this.limit {
		return ""
	} else {
		return fmt.Sprintf(" LIMIT %d,%d", this.getOffset(), this.limit)
	}
}
func (this *Sql) getOrder() string {
	if "" == this.order {
		return ""
	} else {
		return " ORDER BY " + this.order
	}
}
func (this *Sql) getGroup() string {
	if "" == this.group {
		return ""
	} else {
		return " GROUP BY " + this.group
	}
}
func (this *Sql) getRowList() []map[string]interface{} {
	return this.rowList
}
func (this *Sql) getDriver() string {
	if _is.Empty(this.driver) {
		machine := this.getMachine()
		if _is.Empty(machine) {
			panic(`machine is empty`)
		}
		if _is.Empty(machine.Driver) {
			panic("driver is empty")
		}
		this.driver = machine.Driver
	}
	return this.driver
}
func (this *Sql) getDsn() string {
	if _is.Empty(this.dsn) {
		machine := this.getMachine()
		if _is.Empty(machine) {
			panic(`machine is empty`)
		}
		if _is.Empty(machine.Driver) {
			panic("driver is empty")
		}
		switch machine.Driver {
		case "mysql":
			this.dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", machine.Username, machine.Password, machine.Host, machine.Port, machine.Database)
			break
		case "sqlite3":
			this.dsn = machine.Database
			break
		default:
			panic("driver is not support now")
		}
	}
	return this.dsn
}
func (this *Sql) getPool() *sql.DB {
	dsn := this.getDsn()
	driver := this.getDriver()
	var db *sql.DB
	var ok bool
	m.RLock()
	db, ok = pool[dsn]
	m.RUnlock()
	if ok {
		return db
	}
	m.Lock()
	defer m.Unlock()
	db, ok = pool[dsn]
	if ok {
		return db
	}
	db, err := sql.Open(driver, dsn)
	if nil != err {
		panic(err)
	}
	if err := db.Ping(); nil != err {
		panic(err)
	}
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(1 * time.Hour)
	db.SetConnMaxLifetime(1 * time.Hour)
	pool[dsn] = db
	return db
}
func (this *Sql) buildAddList() *Sql {
	rowList := this.getRowList()
	fieldList := _dict.KeyList(rowList[0])
	template := strings.TrimRight(strings.Repeat("("+strings.TrimRight(strings.Repeat("?, ", len(rowList[0])), " ,")+"), ", len(rowList)), " ,")
	parameter := this.getParameter()
	for _, row := range rowList {
		for _, field := range fieldList {
			parameter = append(parameter, row[field])
		}
	}
	table := this.getTable()
	field := _list.Implode(",", fieldList)
	this.Sql(fmt.Sprintf("INSERT INTO %s (%s) VALUES %s", table, field, template))
	this.Parameter(parameter...)
	return this
}
func (this *Sql) buildDel() *Sql {
	table := this.getTable()
	where := this.getWhere()
	this.Sql(fmt.Sprintf("DELETE FROM %s%s", table, where))
	return this
}
func (this *Sql) buildSet() *Sql {
	rowList := this.getRowList()
	row := rowList[0]
	templateList := make([]string, 0, len(row))
	parameter := []interface{}{}
	for field, value := range row {
		templateList = append(templateList, field+" = ?")
		parameter = append(parameter, value)
	}
	template := _list.Implode(",", templateList)
	parameter = append(parameter, this.getParameter()...)
	table := this.getTable()
	where := this.getWhere()
	this.Sql(fmt.Sprintf("UPDATE %s SET %s%s", table, template, where))
	this.Parameter(parameter...)
	return this
}
func (this *Sql) buildGetList() *Sql {
	field := this.getField()
	table := this.getTable()
	sql := fmt.Sprintf("SELECT %s FROM %s", field, table)
	index := this.getIndex()
	if "" != index {
		sql += index
	}
	where := this.getWhere()
	if "" != where {
		sql += where
	}
	group := this.getGroup()
	if "" != group {
		sql += group
	}
	order := this.getOrder()
	if "" != order {
		sql += order
	}
	limit := this.getLimit()
	if "" != limit {
		sql += limit
	}
	this.Sql(sql)
	return this
}
func (this *Sql) buildCount() *Sql {
	table := this.getTable()
	where := this.getWhere()
	sql := fmt.Sprintf("SELECT COUNT(*) AS c FROM %s%s", table, where)
	this.Sql(sql)
	return this
}
func (this *Sql) query() []map[string]string {
	var rowList *sql.Rows
	var err error
	if nil != this.tx {
		rowList, err = this.tx.Query(this.getSql(), this.getParameter()...)
		if nil != err {
			panic(err)
		}
	} else {
		rowList, err = this.getPool().Query(this.getSql(), this.getParameter()...)
		if nil != err {
			panic(err)
		}
	}
	fieldList, err := rowList.Columns()
	if nil != err {
		panic(err)
	}
	dest := make([]interface{}, len(fieldList))
	for i, _ := range dest {
		dest[i] = new(sql.RawBytes)
	}
	res := make([]map[string]string, 0)
	for rowList.Next() {
		if err := rowList.Scan(dest...); nil != err {
			panic(err)
		}
		row := make(map[string]string)
		for i, value := range dest {
			row[fieldList[i]] = string(*(value.(interface{}).(*sql.RawBytes)))
		}
		res = append(res, row)
	}
	return res
}
func (this *Sql) execute() sql.Result {
	if nil != this.tx {
		res, err := this.tx.Exec(this.getSql(), this.getParameter()...)
		if nil != err {
			panic(err)
		}
		return res
	}
	res, err := this.getPool().Exec(this.getSql(), this.getParameter()...)
	if nil != err {
		panic(err)
	}
	return res
}
