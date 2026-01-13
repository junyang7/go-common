package _table

import (
	"fmt"
	"github.com/junyang7/go-common/_cmd"
	"github.com/junyang7/go-common/_directory"
	"github.com/junyang7/go-common/_file"
	"github.com/junyang7/go-common/_name"
	"github.com/junyang7/go-common/_sql"
	"github.com/junyang7/go-common/_string"
	"os"
	"regexp"
	"strings"
)

const Tpl = `package table
type @1@ struct {
@2@
}`

func Build(root string, dbName string, tbName string, showCreateTable string) {
	fieldList := []string{}
	matchedList := regexp.MustCompile("`(\\w+)`\\s+(\\w+)").FindAllStringSubmatch(showCreateTable, -1)
	for _, matched := range matchedList {
		fieldName := matched[1]
		fieldType := "string"
		switch matched[2] {
		case "tinyint", "smallint", "mediumint", "int", "integer", "bigint":
			fieldType = "int64"
			break
		case "float", "double":
			fieldType = "float64"
			break
		}
		fieldList = append(fieldList, fmt.Sprintf("%s %s `json:\"%s\" sql:\"%s\"`", _name.UpperCamelCase(fieldName), fieldType, fieldName, fieldName))
	}
	field := strings.Join(fieldList, "\n")
	tpl := Tpl
	tpl = _string.ReplaceAll(tpl, "@1@", _name.UpperCamelCase(dbName+"_"+tbName))
	tpl = _string.ReplaceAll(tpl, "@2@", field)
	path := root + "/" + dbName + "_" + tbName + ".go"
	_file.Write(path, tpl, os.ModePerm)
}
func BuildByAuto() {
	root := _directory.Current() + "/table"
	poolDict := _sql.GetPoolDict()
	for uk, poolList := range poolDict {
		dbName := strings.TrimSuffix(strings.TrimSuffix(uk, ".master"), ".slaver")
		pool := poolList[0]
		tableList := _sql.New().Pool(pool).Sql(fmt.Sprintf("SELECT `table_name` as `table` FROM `information_schema`.`tables` where `table_schema` = '%s'", dbName)).Query()
		for _, table := range tableList {
			tbName := table["table"]
			res := _sql.New().Pool(pool).Sql(fmt.Sprintf("show create table %s", tbName)).Query()
			showCreateTable := res[0]["Create Table"]
			Build(root, dbName, tbName, showCreateTable)
		}
	}
	cmd := fmt.Sprintf("cd %s && go fmt ./...", root)
	_cmd.ExecuteByStd(cmd)
}
