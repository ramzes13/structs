package structs

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/jackc/pgtype"
)

func TestDiff_GetFields(t *testing.T) {
	type A struct {
		Name sql.NullString `db:"name"`
		IP   sql.NullString `db:"ip"`
	}

	m := GetFields(A{})

	expectedMap := []string{
		"name",
		"ip",
	}

	if !reflect.DeepEqual(m, expectedMap) {
		t.Errorf("The exprected map %+v does't correspond to %+v", expectedMap, m)
	}
}

func TestDiff_OmitCompare(t *testing.T) {
	type A struct {
		Name sql.NullString `db:"name,omitcompare"`
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
		"ip": sql.NullString{String: "", Valid: false},
	}

	if !reflect.DeepEqual(m, expectedMap) {
		t.Errorf("The exprected map %+v does't correspond to %+v", expectedMap, m)
	}
}

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

func TestDiff_EqualJSON(t *testing.T) {
	type A struct {
		Name pgtype.JSONB `db:"name"`
	}

	a := A{
		Name: pgtype.JSONB{Bytes: []byte(`{"a": "test", "b": "test2"}`), Status: pgtype.Present},
	}

	b := A{
		Name: pgtype.JSONB{Bytes: []byte(`{"b": "test2", "a": "test"}`), Status: pgtype.Present},
	}

	m := GenerateDiff(a, b)

	expectedMap := map[string]interface{}{}

	if !reflect.DeepEqual(m, expectedMap) {
		t.Errorf("The exprected map %+v does't correspond to %+v", expectedMap, m)
	}
}

func TestDiff_EqualJSON2(t *testing.T) {
	type A struct {
		Name pgtype.JSONB `db:"name"`
	}

	a := A{
		Name: pgtype.JSONB{Bytes: []byte(`{"a": "test", "b": "test2"}`), Status: pgtype.Present},
	}

	b := A{
		Name: pgtype.JSONB{Bytes: []byte(`{"a": "test", "b": "test2"}`), Status: pgtype.Present},
	}

	m := GenerateDiff(a, b)

	expectedMap := map[string]interface{}{}

	if !reflect.DeepEqual(m, expectedMap) {
		t.Errorf("The exprected map %+v does't correspond to %+v", expectedMap, m)
	}
}
