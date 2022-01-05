package structs

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestDiff_NotEqual(t *testing.T) {
	type A struct {
		Name sql.NullString `db:"name"`
		IP   sql.NullString `db:"ip"`
	}

	a := A{
		Name: sql.NullString{String: "Name", Valid: true},
	}

	b := A{
		IP: sql.NullString{String: "IP", Valid: true},
	}

	m := GenerateDiff(a, b)

	expectedMap := map[string]interface{}{
		"name": sql.NullString{String: "Name", Valid: true},
		"ip":   sql.NullString{String: "", Valid: false},
	}

	if !reflect.DeepEqual(m, expectedMap) {
		t.Errorf("The exprected map %+v does't correspond to %+v", expectedMap, m)
	}
}

func TestDiff_Equal(t *testing.T) {
	type A struct {
		Name sql.NullString `db:"name"`
		IP   sql.NullString `db:"ip"`
	}

	a := A{
		Name: sql.NullString{String: "Name", Valid: true},
		IP:   sql.NullString{String: "IP", Valid: true},
	}

	b := A{
		Name: sql.NullString{String: "Name", Valid: true},
		IP:   sql.NullString{String: "IP", Valid: true},
	}

	m := GenerateDiff(a, b)

	expectedMap := map[string]interface{}{}

	if !reflect.DeepEqual(m, expectedMap) {
		t.Errorf("The exprected map %+v does't correspond to %+v", expectedMap, m)
	}
}
