package _dao

import (
	"fmt"
	"github.com/junyang7/go-common/_cmd"
	"github.com/junyang7/go-common/_directory"
	"github.com/junyang7/go-common/_file"
	"github.com/junyang7/go-common/_sql"
	"github.com/junyang7/go-common/_string"
	"os"
	"strings"
)

const TplDao = `package dao
import "github.com/junyang7/go-common/_sql"
func @1@() *_sql.Sql {
	return _sql.New().Business("@2@").Table("@3@")
}`

func Build() {
	root := _directory.Current() + "/dao"
	poolDict := _sql.GetPoolDict()
	for uk, poolList := range poolDict {
		dbName := strings.TrimSuffix(strings.TrimSuffix(uk, ".master"), ".slaver")
		pool := poolList[0]
		tableList := _sql.New().Pool(pool).Sql(fmt.Sprintf("SELECT `table_name` as `table` FROM `information_schema`.`tables` where `table_schema` = '%s'", dbName)).Query()
		for _, table := range tableList {
			tbName := table["table"]
			tplDao := TplDao
			tplDao = _string.ReplaceAll(tplDao, "@1@", _string.ToUpperCamelCase(dbName+"_"+tbName))
			tplDao = _string.ReplaceAll(tplDao, "@2@", dbName)
			tplDao = _string.ReplaceAll(tplDao, "@3@", tbName)
			path := root + "/" + dbName + "_" + tbName + ".go"
			_file.Write(path, tplDao, os.ModePerm)
		}
	}
	cmd := fmt.Sprintf("cd %s && go fmt ./...", root)
	_cmd.ExecuteByStd(cmd)
}
