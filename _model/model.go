package _model

import (
	"fmt"
	"github.com/junyang7/go-common/_cmd"
	"github.com/junyang7/go-common/_directory"
	"github.com/junyang7/go-common/_file"
	"github.com/junyang7/go-common/_sql"
	"github.com/junyang7/go-common/_string"
	"os"
	"regexp"
	"strings"
)

const TplModel = `package model
type @1@ struct {
@2@
}`

func Build() {
	root := _directory.Current() + "/model"
	pattern := regexp.MustCompile("`(\\w+)`\\s+(\\w+)")
	poolDict := _sql.GetPoolDict()
	for uk, poolList := range poolDict {
		dbName := strings.TrimSuffix(strings.TrimSuffix(uk, ".master"), ".slaver")
		pool := poolList[0]
		tableList := _sql.New().Pool(pool).Sql("show tables").Query()
		for _, table := range tableList {
			tbName := table["Tables_in_runtime"]
			res := _sql.New().Pool(pool).Sql(fmt.Sprintf("show create table %s", tbName)).Query()
			matchedList := pattern.FindAllStringSubmatch(res[0]["Create Table"], -1)
			fieldList := []string{}
			for _, matched := range matchedList {
				fieldName := matched[1]
				fieldType := matched[2]
				fieldTypeGo := "string"
				switch fieldType {
				case "tinyint", "smallint", "mediumint", "int", "integer", "bigint":
					fieldTypeGo = "int64"
					break
				case "float", "double":
					fieldTypeGo = "float64"
					break
				}
				fieldList = append(fieldList, fmt.Sprintf("%s %s `json:\"%s\" sql:\"%s\"`", _string.ToUpperCamelCase(fieldName), fieldTypeGo, fieldName, fieldName))
			}
			field := strings.Join(fieldList, "\n")
			tplModel := TplModel
			tplModel = _string.ReplaceAll(tplModel, "@1@", _string.ToUpperCamelCase(dbName+"_"+tbName))
			tplModel = _string.ReplaceAll(tplModel, "@2@", field)
			path := root + "/" + dbName + "_" + tbName + ".go"
			_file.Write(path, tplModel, os.ModePerm)
		}
	}
	cmd := fmt.Sprintf("cd %s && go fmt ./...", root)
	_cmd.ExecuteByStd(cmd)
}
