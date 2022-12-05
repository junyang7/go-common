package _sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
	"github.com/junyang7/go-common/src/_map"
	"github.com/junyang7/go-common/src/_slice"
	"math/rand"
	"strings"
	"time"
)

type Connection struct {
	Driver    string `json:"driver"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	Database  string `json:"database"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Charset   string `json:"charset"`
	Collation string `json:"collation"`
}

type Group struct {
	Count      int           `json:"count"`
	Connection []*Connection `json:"connection"`
}

type Cluster struct {
	Master *Group `json:"master"`
	Slaver *Group `json:"slaver"`
}

type Conf struct {
	Count   int        `json:"count"`
	Cluster []*Cluster `json:"cluster"`
}

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
	rowList       []map[string]interface{}
	placeholder   string
	sql           string
	tx            *sql.Tx
	dsn           string
}

var sqlConf = map[string]*Conf{}
var sqlPool = map[string]*sql.DB{}

func Initialize(conf map[string]*Conf) {
	sqlConf = conf
	for _, cluster := range sqlConf {
		for _, group := range cluster.Cluster {
			for _, connection := range group.Master.Connection {
				openAndSaveToPool(connection)
			}
			for _, connection := range group.Slaver.Connection {
				openAndSaveToPool(connection)
			}
		}
	}
}

func getDsn(connection *Connection) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", connection.Username, connection.Password, connection.Host, connection.Port, connection.Database)
}
func openAndSaveToPool(connection *Connection) {
	dsn := getDsn(connection)
	if _, ok := sqlPool[dsn]; ok {
		return
	}
	p, err := sql.Open(connection.Driver, dsn)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrSqlOpen).
		Message(err).
		Do()
	p.SetMaxOpenConns(50)
	p.SetConnMaxIdleTime(50)
	p.SetConnMaxLifetime(1 * time.Hour)
	sqlPool[dsn] = p
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
func (this *Sql) getRowList() []map[string]interface{} {
	if 0 == len(this.rowList) {
		return []map[string]interface{}{}
	}
	return this.rowList
}
func (this *Sql) getWhere() string {
	if "" == this.where {
		return ""
	}
	return " WHERE " + this.where
}
func (this *Sql) getPool() *sql.DB {
	if "" == this.dsn {
		if !this.shard {
			this.databaseIndex = 0
		}
		database, ok := sqlConf[this.baseDatabase]
		_interceptor.Insure(ok).
			CodeMessage(_codeMessage.ErrMapKeyNotExists).
			Data(map[string]interface{}{"database": this.baseDatabase}).
			Do()
		_interceptor.Insure(this.databaseIndex < database.Count).
			CodeMessage(_codeMessage.ErrSliceOutOfIndex).
			Data(map[string]interface{}{"database": this.baseDatabase, "index": this.databaseIndex}).
			Do()
		cluster := database.Cluster[this.databaseIndex]
		var group *Group
		if this.master {
			group = cluster.Master
		} else {
			group = cluster.Slaver
		}
		r := 0
		if count := group.Count; count > 1 {
			rand.Seed(time.Now().Unix())
			r = rand.Intn(group.Count - 1)
		}
		this.dsn = getDsn(group.Connection[r])
	}
	return sqlPool[this.dsn]
}

func (this *Sql) buildAddList() *Sql {
	rowList := this.getRowList()
	row := rowList[0]
	fieldList := _map.KeyList(row)
	this.field = _slice.Implode(fieldList, ",")
	this.placeholder = strings.TrimRight(strings.Repeat("("+strings.TrimRight(strings.Repeat("?, ", len(row)), " ,")+"), ", len(rowList)), " ,")
	for _, row := range rowList {
		for _, field := range fieldList {
			this.parameter = append(this.parameter, row[field])
		}
	}
	this.sql = fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;", this.getTable(), this.getField(), this.placeholder)
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
			_interceptor.Insure(false).
				CodeMessage(_codeMessage.ErrSqlTxExec).
				Message(err).
				Do()
		}
		return res
	}
	res, err := this.getPool().Exec(this.sql, this.getParameter()...)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrSqlDBExec).
		Message(err).
		Do()
	return res
}
func (this *Sql) query() []map[string]string {

	var rowList *sql.Rows
	var err error

	if nil != this.tx {
		rowList, err = this.tx.Query(this.sql, this.getParameter()...)
		if err != nil {
			this.Rollback()
			_interceptor.Insure(false).
				CodeMessage(_codeMessage.ErrSqlTxQuery).
				Message(err).
				Do()
		}
	} else {
		rowList, err = this.getPool().Query(this.sql, this.getParameter()...)
		_interceptor.Insure(nil == err).
			CodeMessage(_codeMessage.ErrSqlDBQuery).
			Message(err).
			Do()
	}

	defer func() {
		err := rowList.Close()
		_interceptor.Insure(nil == err).
			CodeMessage(_codeMessage.ErrSqlRowsClose).
			Message(err).
			Do()
	}()

	fieldList, err := rowList.Columns()
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrSqlRowsColumns).
		Message(err).
		Do()
	dest := make([]interface{}, len(fieldList))
	for i, _ := range dest {
		dest[i] = new(sql.RawBytes)
	}
	res := make([]map[string]string, 0)
	for rowList.Next() {
		err := rowList.Scan(dest...)
		_interceptor.Insure(nil == err).
			CodeMessage(_codeMessage.ErrSqlRowsScan).
			Message(err).
			Do()
		row := make(map[string]string)
		for i, value := range dest {
			row[fieldList[i]] = string(*(value.(interface{}).(*sql.RawBytes)))
		}
		res = append(res, row)
	}
	return res

}

func (this *Sql) Add(row map[string]interface{}) int {
	this.master = true
	this.rowList = []map[string]interface{}{row}
	this.buildAddList()
	res := this.execute()
	lastInsertId, err := res.LastInsertId()
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrSqlResultLastInsertId).
		Message(err).
		Do()
	return _as.Int(lastInsertId)
}
func (this *Sql) AddList(rowList []map[string]interface{}) int {
	this.master = true
	this.rowList = rowList
	this.buildAddList()
	res := this.execute()
	lastInsertId, err := res.LastInsertId()
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrSqlResultLastInsertId).
		Message(err).
		Do()
	return _as.Int(lastInsertId)
}
func (this *Sql) Del() int {
	this.master = true
	this.buildDel()
	res := this.execute()
	effectedRowCount, err := res.RowsAffected()
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrSqlResultRowsAffected).
		Message(err).
		Do()
	return _as.Int(effectedRowCount)
}
func (this *Sql) Set(row map[string]interface{}) int {
	this.master = true
	this.row = row
	this.buildSet()
	res := this.execute()
	effectedRowCount, err := res.RowsAffected()
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrSqlResultRowsAffected).
		Message(err).
		Do()
	return _as.Int(effectedRowCount)
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
func (this *Sql) Count() int {
	this.buildCount()
	return _as.Int(this.query()[0]["c"])
}
func (this *Sql) Exists() bool {
	this.buildExists()
	res := this.query()[0]["c"]
	return _as.Int64(res) > 0
}
func (this *Sql) Execute(sql string, add bool) int {
	this.master = true
	this.sql = sql
	res := this.execute()
	if add {
		lastInsertId, err := res.LastInsertId()
		_interceptor.Insure(nil == err).
			CodeMessage(_codeMessage.ErrSqlResultLastInsertId).
			Message(err).
			Do()
		return _as.Int(lastInsertId)
	}
	effectedRowCount, err := res.RowsAffected()
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrSqlResultRowsAffected).
		Message(err).
		Do()
	return _as.Int(effectedRowCount)
}
func (this *Sql) Query(sql string) []map[string]string {
	this.sql = sql
	return this.query()
}
func (this *Sql) Begin() *Sql {
	tx, err := this.getPool().Begin()
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrSqlDBBegin).
		Message(err).
		Do()
	this.tx = tx
	return this
}
func (this *Sql) Commit() *Sql {
	if err := this.tx.Commit(); nil != err {
		this.Rollback()
		_interceptor.Insure(false).
			CodeMessage(_codeMessage.ErrSqlTxCommit).
			Message(err).
			Do()
	}
	return this
}
func (this *Sql) Rollback() *Sql {
	err := this.tx.Rollback()
	_interceptor.Insure(false).
		CodeMessage(_codeMessage.ErrSqlTxRollback).
		Message(err).
		Do()
	return this
}
func (this *Sql) Tx(rd *Sql) *Sql {
	this.tx = rd.tx
	return this
}
