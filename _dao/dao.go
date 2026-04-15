package _dao

import (
	"fmt"
	"github.com/junyang7/go-common/_cmd"
	"github.com/junyang7/go-common/_conf"
	"github.com/junyang7/go-common/_directory"
	"github.com/junyang7/go-common/_file"
	"github.com/junyang7/go-common/_json"
	"github.com/junyang7/go-common/_name"
	"github.com/junyang7/go-common/_sql"
	"github.com/junyang7/go-common/_string"
	"strings"
)

const Tpl = `package dao
import "github.com/junyang7/go-common/_sql"
func @1@() *_sql.Sql {
	return _sql.New().Business("@2@").Table("@3@")
}`

func Build(useDatabasePrefix bool, root string, dbName string, tbName string) {
	if !_directory.Exists(root) {
		_directory.Create(root)
	}
	prefix := ""
	if useDatabasePrefix {
		prefix = dbName + "_"
	}
	tpl := Tpl
	tpl = _string.ReplaceAll(tpl, "@1@", _name.UpperCamelCase(prefix+tbName))
	tpl = _string.ReplaceAll(tpl, "@2@", dbName)
	tpl = _string.ReplaceAll(tpl, "@3@", tbName)
	path := root + "/" + strings.ToLower(prefix+tbName) + ".go"
	_file.Write(path, tpl)
}
func BuildByAuto(useDatabasePrefix bool) {
	root := _directory.Current() + "/dao"
	if !_directory.Exists(root) {
		_directory.Create(root)
	}
	raw := _conf.Get("sql").Value()
	var clusterDict map[string]*_sql.Cluster
	_json.Decode(_json.Encode(raw), &clusterDict)
	if len(clusterDict) == 0 {
		return
	}
	for dbName, ms := range clusterDict {
		machine := &_sql.Machine{}
		if len(ms.Master) > 0 {
			machine = ms.Master[0]
		}
		if len(ms.Slaver) > 0 {
			machine = ms.Slaver[0]
		}
		sql := fmt.Sprintf("SELECT `table_name` as `table` FROM `information_schema`.`tables` where `table_schema` = '%s'", machine.Database)
		tableList := _sql.New().Business(dbName).Sql(sql).Query()
		for _, table := range tableList {
			tbName := table["table"]
			Build(useDatabasePrefix, root, dbName, tbName)
		}
	}
	cmd := fmt.Sprintf("cd %s && go fmt ./...", root)
	_cmd.ExecuteByStd(cmd)
}
