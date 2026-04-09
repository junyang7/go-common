package _sql

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/junyang7/go-common/_assert"
)

type sqlBindModel struct {
	ID       int    `sql:"id,pk"`
	Name     string `sql:"name,text"`
	Keep     string `sql:"keep"`
	Skip     string `sql:"-"`
	Untagged string
	hidden   string `sql:"hidden"`
}

func TestResolveSQLTagName(t *testing.T) {
	fields := reflect.TypeOf(struct {
		ID       int    `sql:"id"`
		Name     string `sql:"name,pk"`
		Skip     string `sql:"-"`
		Literal  string `sql:"-,"`
		Untagged string
		Empty    string `sql:""`
	}{})

	name, ok := resolveSQLTagName(fields.Field(0))
	_assert.True(t, ok)
	_assert.Equal(t, name, "id")

	name, ok = resolveSQLTagName(fields.Field(1))
	_assert.True(t, ok)
	_assert.Equal(t, name, "name")

	name, ok = resolveSQLTagName(fields.Field(2))
	_assert.False(t, ok)
	_assert.Equal(t, name, "")

	name, ok = resolveSQLTagName(fields.Field(3))
	_assert.True(t, ok)
	_assert.Equal(t, name, "-")

	name, ok = resolveSQLTagName(fields.Field(4))
	_assert.False(t, ok)
	_assert.Equal(t, name, "")

	name, ok = resolveSQLTagName(fields.Field(5))
	_assert.False(t, ok)
	_assert.Equal(t, name, "")
}

func TestSqlBindRowToStruct(t *testing.T) {
	model := sqlBindModel{
		Keep:     "preserve",
		Untagged: "stay",
		hidden:   "secret",
	}
	row := map[string]string{
		"id":     "7",
		"name":   "alice",
		"-":      "skip",
		"hidden": "should_not_bind",
	}
	s := New().Ignore("name")
	s.bindRowToStruct(row, reflect.ValueOf(&model).Elem())

	_assert.Equal(t, model.ID, 7)
	_assert.Equal(t, model.Name, "")
	_assert.Equal(t, model.Keep, "preserve")
	_assert.Equal(t, model.Skip, "")
	_assert.Equal(t, model.Untagged, "stay")
	_assert.Equal(t, model.hidden, "secret")
}

func TestSqlBuildRowList(t *testing.T) {
	{
		model := sqlBindModel{
			ID:       1,
			Name:     "alice",
			Keep:     "keep",
			Skip:     "skip",
			Untagged: "untagged",
			hidden:   "secret",
		}
		s := New().Bind(&model).Ignore("keep")
		s.buildRowList()

		_assert.Equal(t, len(s.rowList), 1)
		_assert.Equal(t, s.rowList[0]["id"], 1)
		_assert.Equal(t, s.rowList[0]["name"], "alice")
		_, exists := s.rowList[0]["keep"]
		_assert.False(t, exists)
		_, exists = s.rowList[0]["-"]
		_assert.False(t, exists)
		_, exists = s.rowList[0]["hidden"]
		_assert.False(t, exists)
		_, exists = s.rowList[0]["Untagged"]
		_assert.False(t, exists)
	}
	{
		list := []sqlBindModel{
			{ID: 1, Name: "alice", Keep: "k1"},
			{ID: 2, Name: "bob", Keep: "k2"},
		}
		s := New().Bind(&list)
		s.buildRowList()

		_assert.Equal(t, len(s.rowList), 2)
		_assert.Equal(t, s.rowList[0]["id"], 1)
		_assert.Equal(t, s.rowList[0]["name"], "alice")
		_assert.Equal(t, s.rowList[1]["id"], 2)
		_assert.Equal(t, s.rowList[1]["name"], "bob")
	}
}

func TestSqlGetAndGetListBind(t *testing.T) {
	db := openSQLiteForSQLBindTest(t)

	_, err := db.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY,
			name TEXT,
			keep TEXT
		)
	`)
	_assert.NoError(t, err)

	_, err = db.Exec(`INSERT INTO users (id, name, keep) VALUES (1, 'alice', 'db_keep_1'), (2, 'bob', 'db_keep_2')`)
	_assert.NoError(t, err)

	{
		model := sqlBindModel{
			Keep:     "preserve",
			Untagged: "stay",
			hidden:   "secret",
		}
		row := New().
			Pool(db).
			Table("users").
			Field("id,name").
			Where("id = 1").
			Bind(&model).
			Get()

		_assert.Equal(t, row["id"], "1")
		_assert.Equal(t, row["name"], "alice")
		_assert.Equal(t, model.ID, 1)
		_assert.Equal(t, model.Name, "alice")
		_assert.Equal(t, model.Keep, "preserve")
		_assert.Equal(t, model.Skip, "")
		_assert.Equal(t, model.Untagged, "stay")
		_assert.Equal(t, model.hidden, "secret")
	}

	{
		list := []sqlBindModel{}
		rowList := New().
			Pool(db).
			Table("users").
			Field("id,name").
			Order("id ASC").
			Bind(&list).
			GetList()

		_assert.Equal(t, len(rowList), 2)
		_assert.Equal(t, len(list), 2)
		_assert.Equal(t, list[0].ID, 1)
		_assert.Equal(t, list[0].Name, "alice")
		_assert.Equal(t, list[0].Keep, "")
		_assert.Equal(t, list[1].ID, 2)
		_assert.Equal(t, list[1].Name, "bob")
	}
}

func TestSqlAddAddListSetAndCountBind(t *testing.T) {
	db := openSQLiteForSQLBindTest(t)

	_, err := db.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY,
			name TEXT,
			keep TEXT
		)
	`)
	_assert.NoError(t, err)

	{
		model := sqlBindModel{
			ID:       1,
			Name:     "alice",
			Keep:     "keep_1",
			Skip:     "skip",
			Untagged: "untagged",
			hidden:   "secret",
		}
		insertID := New().
			Pool(db).
			Table("users").
			Bind(&model).
			Add()

		_assert.Equal(t, insertID, int64(1))

		row := New().
			Pool(db).
			Table("users").
			Field("id,name,keep").
			Where("id = 1").
			Get()

		_assert.Equal(t, row["id"], "1")
		_assert.Equal(t, row["name"], "alice")
		_assert.Equal(t, row["keep"], "keep_1")
	}

	{
		list := []sqlBindModel{
			{ID: 2, Name: "bob", Keep: "keep_2"},
			{ID: 3, Name: "cindy", Keep: "keep_3"},
		}
		insertID := New().
			Pool(db).
			Table("users").
			Bind(&list).
			AddList()

		_assert.Equal(t, insertID, int64(3))

		rowList := New().
			Pool(db).
			Table("users").
			Field("id,name,keep").
			Order("id ASC").
			GetList()

		_assert.Equal(t, len(rowList), 3)
		_assert.Equal(t, rowList[1]["id"], "2")
		_assert.Equal(t, rowList[1]["name"], "bob")
		_assert.Equal(t, rowList[2]["id"], "3")
		_assert.Equal(t, rowList[2]["name"], "cindy")
	}

	{
		model := sqlBindModel{
			Name: "alice_updated",
			Keep: "should_be_ignored",
		}
		affected := New().
			Pool(db).
			Table("users").
			Where("id = 1").
			Bind(&model).
			Ignore("id", "keep").
			Set()

		_assert.Equal(t, affected, int64(1))

		row := New().
			Pool(db).
			Table("users").
			Field("id,name,keep").
			Where("id = 1").
			Get()

		_assert.Equal(t, row["name"], "alice_updated")
		_assert.Equal(t, row["keep"], "keep_1")
	}

	{
		var countInt int
		count := New().
			Pool(db).
			Table("users").
			Bind(&countInt).
			Count()

		_assert.Equal(t, count, int64(3))
		_assert.Equal(t, countInt, 3)
	}

	{
		var countNil *int64
		count := New().
			Pool(db).
			Table("users").
			Bind(countNil).
			Count()

		_assert.Equal(t, count, int64(3))
	}
}

func openSQLiteForSQLBindTest(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("sqlite3", ":memory:")
	_assert.NoError(t, err)
	db.SetMaxOpenConns(1)

	t.Cleanup(func() {
		_ = db.Close()
	})

	return db
}
