package _dao

import (
	"fmt"
	"github.com/junyang7/go-common/_cmd"
	"github.com/junyang7/go-common/_directory"
	"github.com/junyang7/go-common/_file"
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

func Build(root string, dbName string, tbName string) {
	tpl := Tpl
	tpl = _string.ReplaceAll(tpl, "@1@", _name.UpperCamelCase(dbName+"_"+tbName))
	tpl = _string.ReplaceAll(tpl, "@2@", dbName)
	tpl = _string.ReplaceAll(tpl, "@3@", tbName)
	path := root + "/" + strings.ToLower(dbName) + "_" + strings.ToLower(tbName) + ".go"
	_file.Write(path, tpl, 0644)
}
func BuildByAuto() {
	root := _directory.Current() + "/dao"
	poolDict := _sql.GetPoolDict()
	for uk, poolList := range poolDict {
		dbName := strings.TrimSuffix(strings.TrimSuffix(uk, ".master"), ".slaver")
		pool := poolList[0]
		tableList := _sql.New().Pool(pool).Sql(fmt.Sprintf("SELECT `table_name` as `table` FROM `information_schema`.`tables` where `table_schema` = '%s'", dbName)).Query()
		for _, table := range tableList {
			tbName := table["table"]
			Build(root, dbName, tbName)
		}
	}
	cmd := fmt.Sprintf("cd %s && go fmt ./...", root)
	_cmd.ExecuteByStd(cmd)
}
