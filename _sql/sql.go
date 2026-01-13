package _sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/junyang7/go-common/_as"
	"github.com/junyang7/go-common/_conf"
	"github.com/junyang7/go-common/_dict"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_json"
	"github.com/junyang7/go-common/_list"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
	"reflect"
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
type Cluster struct {
	Master []*Machine `json:"master"`
	Slaver []*Machine `json:"slaver"`
}

var poolDict = map[string][]*sql.DB{}
var m = sync.RWMutex{}

func getPoolDictName(business string, master bool) string {
	poolDictName := business + "."
	if master {
		poolDictName += "master"
	} else {
		poolDictName += "slaver"
	}
	return poolDictName
}
func getDsn(machine *Machine) string {
	var dsn string
	switch machine.Driver {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", machine.Username, machine.Password, machine.Host, machine.Port, machine.Database)
	case "sqlite3":
		dsn = machine.Database
	default:
		_interceptor.Insure(false).Message(`driver is not support now`).Do()
	}
	return dsn
}
func open(machine *Machine) (pool *sql.DB) {
	dsn := getDsn(machine)
	pool, err := sql.Open(machine.Driver, dsn)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	if err := pool.Ping(); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	pool.SetMaxOpenConns(10)
	pool.SetConnMaxIdleTime(1 * time.Hour)
	pool.SetConnMaxLifetime(1 * time.Hour)
	return pool
}
func Load() {
	raw := _conf.Get("sql").Value()
	var clusterDict map[string]*Cluster
	_json.Decode(_json.Encode(raw), &clusterDict)
	if len(clusterDict) == 0 {
		return
	}
	for business, ms := range clusterDict {
		for _, machine := range ms.Master {
			poolDictName := getPoolDictName(business, true)
			poolDict[poolDictName] = append(poolDict[poolDictName], open(machine))
		}
		for _, machine := range ms.Slaver {
			poolDictName := getPoolDictName(business, false)
			poolDict[poolDictName] = append(poolDict[poolDictName], open(machine))
		}
	}
}
func GetPoolDict() map[string][]*sql.DB {
	return poolDict
}

type Sql struct {
	business          string                   // 业务标识，用来区分资源
	masterMachineList []*Machine               // 主库机器配置列表
	slaverMachineList []*Machine               // 从库机器配置列表
	table             string                   // 表名
	sql               string                   // 最终执行的SQL语句
	parameter         []interface{}            // 最终执行的SQL语句配套的参数列表
	master            bool                     // 是否使用主库
	where             string                   // where条件
	field             string                   // 字段
	index             string                   // 索引
	offset            int                      // 偏移量
	limit             int                      // 限制数
	order             string                   // 排序
	group             string                   // 分组
	having            string                   // having条件
	rowList           []map[string]interface{} // 多行表数据
	tx                *sql.Tx                  // 事务
	pool              *sql.DB                  // 数据库连接池
	bind              interface{}              // 从这里解析数据到rowList或者从结果集中解析数据到这里
	ignore            []string                 // bind时忽略的字段
}

func New() *Sql {
	return &Sql{
		parameter: []interface{}{},
		rowList:   []map[string]interface{}{},
	}
}
func (this *Sql) Pool(pool *sql.DB) *Sql {
	this.pool = pool
	return this
}
func (this *Sql) Business(business string) *Sql {
	this.business = business
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
func (this *Sql) Query() []map[string]string {
	return this.query()
}
func (this *Sql) Machine(machine *Machine) *Sql {
	this.masterMachineList = append(this.masterMachineList, machine)
	this.slaverMachineList = append(this.slaverMachineList, machine)
	return this
}
func (this *Sql) MasterMachine(machine *Machine) *Sql {
	this.masterMachineList = append(this.masterMachineList, machine)
	return this
}
func (this *Sql) SlaverMachine(machine *Machine) *Sql {
	this.slaverMachineList = append(this.slaverMachineList, machine)
	return this
}
func (this *Sql) Master(master bool) *Sql {
	this.master = master
	return this
}
func (this *Sql) Table(table string) *Sql {
	this.table = table
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
func (this *Sql) Having(having string) *Sql {
	this.having = having
	return this
}
func (this *Sql) RowList(rowList []map[string]interface{}) *Sql {
	this.rowList = rowList
	return this
}
func (this *Sql) Row(row map[string]interface{}) *Sql {
	this.rowList = append(this.rowList, row)
	return this
}
func (this *Sql) Bind(bind interface{}) *Sql {
	this.bind = bind
	return this
}
func (this *Sql) Ignore(ignore ...string) *Sql {
	this.ignore = ignore
	return this
}
func (this *Sql) Tx(tx *sql.Tx) *Sql {
	this.tx = tx
	return this
}
func (this *Sql) Begin() *sql.Tx {
	if nil == this.tx {
		this.Master(true)
		tx, err := this.getPool().Begin()
		if nil != err {
			_interceptor.Insure(false).Message(err).Do()
		}
		this.Tx(tx)
	}
	return this.tx
}
func (this *Sql) Commit() {
	if err := this.tx.Commit(); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}
func (this *Sql) Rollback() {
	if err := this.tx.Rollback(); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}
func (this *Sql) Transaction(group func(tx *sql.Tx)) {
	this.Begin()
	defer func() {
		if err := recover(); nil != err {
			this.Rollback()
			_interceptor.Insure(false).Message(err).Do()
		}
	}()
	group(this.tx)
	this.Commit()
}
func (this *Sql) AddList() int64 {
	this.Master(true)
	this.buildRowList()
	this.buildAddList()
	res := this.execute()
	lastInsertId, err := res.LastInsertId()
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return lastInsertId
}
func (this *Sql) Add() int64 {
	this.Master(true)
	this.buildRowList()
	this.buildAddList()
	res := this.execute()
	lastInsertId, err := res.LastInsertId()
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return lastInsertId
}
func (this *Sql) Del() int64 {
	this.Master(true)
	this.buildDel()
	res := this.execute()
	effectedRowCount, err := res.RowsAffected()
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return effectedRowCount
}
func (this *Sql) Set() int64 {
	this.Master(true)
	this.buildRowList()
	this.buildSet()
	res := this.execute()
	effectedRowCount, err := res.RowsAffected()
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return effectedRowCount
}
func (this *Sql) GetList() (o []map[string]string) {
	this.buildGetList()
	o = this.query()
	if nil != this.bind {
		bind := reflect.ValueOf(this.bind)
		if bind.Kind() == reflect.Ptr && bind.Elem().Kind() == reflect.Slice && bind.Elem().Type().Elem().Kind() == reflect.Struct {
			for _, row := range o {
				v := reflect.New(bind.Elem().Type().Elem()).Elem()
				t := v.Type()
				for i := 0; i < v.NumField(); i++ {
					f := v.Field(i)
					n := t.Field(i).Tag.Get("sql")
					if !_list.In(n, this.ignore) {
						switch f.Kind() {
						case reflect.String:
							f.SetString(row[n])
							break
						case reflect.Bool:
							f.SetBool(_as.Bool(row[n]))
							break
						case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
							f.SetInt(_as.Int64(row[n]))
							break
						case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
							f.SetUint(_as.Uint64(row[n]))
							break
						case reflect.Float32, reflect.Float64:
							f.SetFloat(_as.Float64(row[n]))
							break
						default:
							_interceptor.Insure(false).Message("数据类型不支持").Do()
						}
					}
				}
				bind.Elem().Set(reflect.Append(bind.Elem(), v))
			}
		}
	}
	return o
}
func (this *Sql) Get() (o map[string]string) {
	this.buildGetList()
	res := this.query()
	if len(res) > 0 {
		o = res[0]
	} else {
		o = map[string]string{}
	}
	if nil != this.bind {
		bind := reflect.ValueOf(this.bind)
		if bind.Kind() == reflect.Ptr && bind.Elem().Kind() == reflect.Struct {
			v := bind.Elem()
			t := v.Type()
			if len(o) > 0 {
				row := o
				for i := 0; i < v.NumField(); i++ {
					f := v.Field(i)
					n := t.Field(i).Tag.Get("sql")
					if !_list.In(n, this.ignore) {
						switch f.Kind() {
						case reflect.String:
							f.SetString(row[n])
							break
						case reflect.Bool:
							f.SetBool(_as.Bool(row[n]))
							break
						case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
							f.SetInt(_as.Int64(row[n]))
							break
						case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
							f.SetUint(_as.Uint64(row[n]))
							break
						case reflect.Float32, reflect.Float64:
							f.SetFloat(_as.Float64(row[n]))
							break
						default:
							_interceptor.Insure(false).Message("数据类型不支持").Do()
						}
					}
				}
			}
		}
	}
	return o
}
func (this *Sql) Count() (o int64) {
	this.buildCount()
	res := this.query()
	if len(res) > 0 {
		o = _as.Int64(res[0]["c"])
	} else {
		o = 0
	}
	if nil != this.bind {
		if this.bind == nil {
			return
		}
		bind := reflect.ValueOf(this.bind)
		if bind.Kind() != reflect.Ptr {
			return
		}
		elem := bind.Elem()
		val := reflect.ValueOf(o)
		if val.Type().AssignableTo(elem.Type()) {
			elem.Set(val)
		} else if val.Type().ConvertibleTo(elem.Type()) {
			elem.Set(val.Convert(elem.Type()))
		}
	}
	return o
}
func (this *Sql) Execute() int64 {
	this.Master(true)
	this.buildRowList()
	sql := this.getSql()
	res := this.execute()
	if strings.HasPrefix(strings.ToLower(sql), "insert") {
		lastInsertId, err := res.LastInsertId()
		if nil != err {
			_interceptor.Insure(false).Message(err).Do()
		}
		return lastInsertId
	}
	effectedRowCount, err := res.RowsAffected()
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return effectedRowCount
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
	return "`" + strings.Trim(strings.TrimSpace(this.table), "`") + "`"
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
		fieldList := strings.Split(this.field, ",")
		for k, v := range fieldList {
			fieldList[k] = "`" + strings.Trim(strings.TrimSpace(v), "`") + "`"
		}
		return strings.Join(fieldList, ",")
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
func (this *Sql) getHaving() string {
	if "" == this.having {
		return ""
	} else {
		return " HAVING " + this.having
	}
}
func (this *Sql) getRowList() []map[string]interface{} {
	return this.rowList
}
func (this *Sql) getMasterMachineList() []*Machine {
	return this.masterMachineList
}
func (this *Sql) getSlaverMachineList() []*Machine {
	return this.slaverMachineList
}
func (this *Sql) getBusiness() string {
	return this.business
}
func (this *Sql) getPool() *sql.DB {
	if nil != this.pool {
		return this.pool
	}
	business := this.getBusiness()
	master := this.getMaster()
	poolDictName := getPoolDictName(business, master)
	var poolList []*sql.DB
	var ok bool
	m.RLock()
	poolList, ok = poolDict[poolDictName]
	m.RUnlock()
	if ok {
		r := rand.Intn(len(poolList))
		this.pool = poolList[r]
		return this.pool
	}
	m.Lock()
	defer m.Unlock()
	poolList, ok = poolDict[poolDictName]
	if ok {
		r := rand.Intn(len(poolList))
		this.pool = poolList[r]
		return this.pool
	}
	var machineList []*Machine
	if master {
		machineList = this.getMasterMachineList()
	} else {
		machineList = this.getSlaverMachineList()
	}
	if len(machineList) > 0 {
		r := rand.Intn(len(machineList))
		machine := machineList[r]
		this.pool = open(machine)
		poolDict[poolDictName] = append(poolDict[poolDictName], open(machine))
		return this.pool
	}
	_interceptor.Insure(false).Message("[sql]没有找到相关配置").Do()
	return nil
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
	for k, v := range fieldList {
		fieldList[k] = "`" + strings.Trim(strings.TrimSpace(v), "`") + "`"
	}
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
		templateList = append(templateList, "`"+strings.Trim(field, "`")+"`"+" = ?")
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
	having := this.getHaving()
	if "" != having {
		sql += having
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
	having := this.getHaving()
	sql := fmt.Sprintf("SELECT COUNT(*) AS c FROM %s%s%s", table, where, having)
	this.Sql(sql)
	return this
}
func (this *Sql) buildRowList() *Sql {
	if this.bind == nil {
		return this
	}
	bind := reflect.ValueOf(this.bind)
	if bind.Kind() != reflect.Ptr {
		newPtr := reflect.New(bind.Type())
		newPtr.Elem().Set(bind)
		bind = newPtr
	}
	elem := bind.Elem()
	if elem.Kind() == reflect.Slice || elem.Kind() == reflect.Array {
		for i, j := 0, elem.Len(); i < j; i++ {
			item := elem.Index(i)
			if item.Kind() == reflect.Ptr {
				if !item.IsNil() {
					item = item.Elem()
				} else {
					continue
				}
			}
			if item.Kind() == reflect.Struct {
				row := make(map[string]interface{})
				for k := 0; k < item.NumField(); k++ {
					f := item.Field(k)
					n := item.Type().Field(k).Tag.Get("sql")
					if n != "" && !_list.In(n, this.ignore) {
						row[n] = f.Interface()
					}
				}
				this.Row(row)
			}
		}
	} else if elem.Kind() == reflect.Struct {
		row := make(map[string]interface{})
		for k := 0; k < elem.NumField(); k++ {
			f := elem.Field(k)
			n := elem.Type().Field(k).Tag.Get("sql")
			if n != "" && !_list.In(n, this.ignore) {
				row[n] = f.Interface()
			}
		}
		this.Row(row)
	}
	return this
}
func (this *Sql) query() []map[string]string {
	var rowList *sql.Rows
	var err error
	if nil != this.tx {
		rowList, err = this.tx.Query(this.getSql(), this.getParameter()...)
		if nil != err {
			_interceptor.Insure(false).Message(err).Do()
		}
	} else {
		rowList, err = this.getPool().Query(this.getSql(), this.getParameter()...)
		if nil != err {
			_interceptor.Insure(false).Message(err).Do()
		}
	}
	fieldList, err := rowList.Columns()
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	dest := make([]interface{}, len(fieldList))
	for i, _ := range dest {
		dest[i] = new(sql.RawBytes)
	}
	res := make([]map[string]string, 0)
	for rowList.Next() {
		if err := rowList.Scan(dest...); nil != err {
			_interceptor.Insure(false).Message(err).Do()
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
			_interceptor.Insure(false).Message(err).Do()
		}
		return res
	}
	res, err := this.getPool().Exec(this.getSql(), this.getParameter()...)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
