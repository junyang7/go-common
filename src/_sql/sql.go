package _sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_sql/_conf"
	"math/rand"
	"strings"
	"time"
)

type Sql struct {
	baseDatabase  string
	baseTable     string
	shard         bool
	databaseIndex int
	tableIndex    int
	master        bool
	index         string
	field         string
	where         string
	group         string
	order         string
	offset        int
	limit         int
	parameter     []interface{}
	row           map[string]interface{}
	placeholder   string
	sql           string
	tx            *sql.Tx
	machine       *_conf.Machine
}

func New(baseDatabase string, baseTable string) *Sql {
	return &Sql{
		baseDatabase: baseDatabase,
		baseTable:    baseTable,
	}
}
func (this *Sql) DatabaseIndex(databaseIndex int) *Sql {
	this.databaseIndex = databaseIndex
	return this
}
func (this *Sql) TableIndex(tableIndex int) *Sql {
	this.tableIndex = tableIndex
	return this
}
func (this *Sql) Shard() *Sql {
	this.shard = true
	return this
}
func (this *Sql) Master() *Sql {
	this.master = true
	return this
}
func (this *Sql) Index(index string) *Sql {
	this.index = index
	return this
}
func (this *Sql) Field(field string) *Sql {
	this.field = field
	return this
}
func (this *Sql) Where(where string) *Sql {
	this.where = where
	return this
}
func (this *Sql) Order(order string) *Sql {
	this.order = order
	return this
}
func (this *Sql) Offset(offset int) *Sql {
	this.offset = offset
	return this
}
func (this *Sql) Group(group string) *Sql {
	this.group = group
	return this
}
func (this *Sql) Limit(limit int) *Sql {
	this.limit = limit
	return this
}
func (this *Sql) Parameter(parameter ...interface{}) *Sql {
	this.parameter = parameter
	return this
}

func (this *Sql) getTable() string {
	if this.shard {
		return this.baseTable + "_" + _as.String(this.tableIndex)
	}
	return this.baseTable
}
func (this *Sql) getField() string {
	if "" == this.field {
		return "*"
	}
	return this.field
}
func (this *Sql) getIndex() string {
	if "" == this.index {
		return this.index
	}
	return " FORCE INDEX (" + this.index + ")"
}
func (this *Sql) getOrder() string {
	if "" == this.order {
		return ""
	}
	return " ORDER BY " + this.order
}
func (this *Sql) getParameter() []interface{} {
	if nil == this.parameter {
		return []interface{}{}
	}
	return this.parameter
}
func (this *Sql) getGroup() string {
	if "" == this.group {
		return ""
	}
	return " GROUP BY " + this.group
}
func (this *Sql) getRow() map[string]interface{} {
	if nil == this.row {
		return map[string]interface{}{}
	}
	return this.row
}
func (this *Sql) getWhere() string {
	if "" == this.where {
		return ""
	}
	return " WHERE " + this.where
}
func (this *Sql) getPool() *sql.DB {
	if nil == this.machine {
		cluster := _conf.Conf.Database[this.baseDatabase].Cluster[_as.String(this.databaseIndex)]
		if this.master {
			this.machine = cluster.Master.Machine[0]
		} else {
			r := 0
			if count := cluster.Slaver.Count; count > 1 {
				rand.Seed(time.Now().Unix())
				r = rand.Intn(cluster.Slaver.Count - 1)
			}
			this.machine = cluster.Slaver.Machine[r]
		}
	}
	if nil == this.machine.Pool {
		pool, err := sql.Open(this.machine.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", this.machine.Username, this.machine.Password, this.machine.Host, this.machine.Port, this.machine.Database))
		if nil != err {
			panic(err)
		}
		this.machine.Pool = pool
	}
	return this.machine.Pool
}

func (this *Sql) buildAdd() *Sql {
	row := this.getRow()
	fieldList := make([]string, 0, len(row))
	for field, _ := range row {
		fieldList = append(fieldList, field)
	}
	this.field = strings.Join(fieldList, ",")
	this.parameter = make([]interface{}, 0, len(row))
	placeholderList := make([]string, 0, len(row))
	for _, field := range fieldList {
		this.parameter = append(this.parameter, row[field])
		placeholderList = append(placeholderList, "?")
	}
	this.placeholder = strings.Join(placeholderList, ",")
	this.sql = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", this.getTable(), this.getField(), this.placeholder)
	return this
}
func (this *Sql) buildCount() *Sql {
	this.sql = fmt.Sprintf("SELECT COUNT(*) AS c FROM %s%s", this.getTable(), this.getWhere())
	return this
}
func (this *Sql) buildDel() *Sql {
	this.sql = fmt.Sprintf("DELETE FROM %s%s", this.getTable(), this.getWhere())
	return this
}
func (this *Sql) buildExists() *Sql {
	this.sql = fmt.Sprintf("SELECT COUNT(*) AS c FROM %s%s LIMIT 1", this.getTable(), this.getWhere())
	return this
}
func (this *Sql) buildGetList() *Sql {
	this.sql = "SELECT %s FROM %s"
	index := this.getIndex()
	if index != "" {
		this.sql += index
	}
	where := this.getWhere()
	if where != "" {
		this.sql += where
	}
	group := this.getGroup()
	if group != "" {
		this.sql += group
	}
	order := this.getOrder()
	if order != "" {
		this.sql += order
	}
	limit := this.limit
	if limit > 0 {
		offset := this.offset
		this.sql += " LIMIT " + _as.String(offset) + "," + _as.String(limit)
	}
	this.sql = fmt.Sprintf(this.sql, this.getField(), this.getTable())
	return this
}
func (this *Sql) buildSet() *Sql {
	row := this.getRow()
	placeholderList := make([]string, 0, len(row))
	parameterList := make([]interface{}, 0, len(row))
	for field, value := range row {
		placeholderList = append(placeholderList, field+" = ?")
		parameterList = append(parameterList, value)
	}
	this.placeholder = strings.Join(placeholderList, ",")
	this.parameter = append(parameterList, this.parameter...)
	this.sql = fmt.Sprintf("UPDATE %s SET %s%s", this.getTable(), this.placeholder, this.getWhere())
	return this
}

func (this *Sql) execute() sql.Result {
	if nil != this.tx {
		res, err := this.tx.Exec(this.sql, this.getParameter()...)
		if err != nil {
			this.Rollback()
			panic(err)
		}
		return res
	}
	res, err := this.getPool().Exec(this.sql, this.getParameter()...)
	if nil != err {
		panic(err)
	}
	return res
}
func (this *Sql) query() []map[string]string {

	var rowList *sql.Rows
	var err error

	if nil != this.tx {
		rowList, err = this.tx.Query(this.sql, this.getParameter()...)
		if err != nil {
			this.Rollback()
			panic(err)
		}
	} else {
		rowList, err = this.getPool().Query(this.sql, this.getParameter()...)
		if nil != err {
			panic(err)
		}
	}

	defer func() {
		err := rowList.Close()
		if nil != err {
			panic(err)
		}
	}()

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
		err := rowList.Scan(dest...)
		if nil != err {
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

func (this *Sql) Add(row map[string]interface{}) int64 {
	this.master = true
	this.row = row
	this.buildAdd()
	res := this.execute()
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return lastInsertId
}
func (this *Sql) Del() int64 {
	this.master = true
	this.buildDel()
	res := this.execute()
	effectedRowCount, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return effectedRowCount
}
func (this *Sql) Set(row map[string]interface{}) int64 {
	this.master = true
	this.row = row
	this.buildSet()
	res := this.execute()
	effectedRowCount, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return effectedRowCount
}
func (this *Sql) Get() map[string]string {
	this.limit = 1
	this.buildGetList()
	res := this.query()
	if len(res) > 0 {
		return res[0]
	} else {
		return map[string]string{}
	}
}
func (this *Sql) GetList() []map[string]string {
	this.buildGetList()
	return this.query()
}
func (this *Sql) Count() int64 {
	this.buildCount()
	return _as.Int64(this.query()[0]["c"])
}
func (this *Sql) Exists() bool {
	this.buildExists()
	res := this.query()[0]["c"]
	return _as.Int64(res) > 0
}
func (this *Sql) Execute(sql string, add bool) int64 {
	this.master = true
	this.sql = sql
	res := this.execute()
	if add {
		lastInsertId, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		return lastInsertId
	}
	effectedRowCount, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	return effectedRowCount
}
func (this *Sql) Query(sql string) []map[string]string {
	this.sql = sql
	return this.query()
}
func (this *Sql) Begin() *Sql {
	tx, err := this.getPool().Begin()
	if nil != err {
		panic(err)
	}
	this.tx = tx
	return this
}
func (this *Sql) Commit() *Sql {
	if err := this.tx.Commit(); nil != err {
		this.Rollback()
		panic(err)
	}
	return this
}
func (this *Sql) Rollback() *Sql {
	if err := this.tx.Rollback(); nil != err {
		panic(err)
	}
	return this
}
func (this *Sql) Tx(rd *Sql) *Sql {
	this.tx = rd.tx
	return this
}
