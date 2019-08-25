package nullable

import (
	"database/sql"
	"testing"
)

func TestNullableString(t *testing.T) {
	for _, unit := range []struct {
		value    *String
		expected string
	}{
		{&String{sql.NullString{"aquas", true}}, "\"aquas\""},
		{&String{sql.NullString{"aquas", false}}, "null"},
		{&String{sql.NullString{"", false}}, "null"},
	} {
		if actually, _ := unit.value.MarshalJSON(); string(actually) != unit.expected {
			t.Errorf("combination: [%v], type:%T, actually: [%v], type: %T", unit.expected, unit.expected, string(actually), string(actually))
		}
	}
}

func TestNullableInt64(t *testing.T) {
	for _, unit := range []struct {
		value    *Int64
		expected string
	}{
		{&Int64{sql.NullInt64{0, false}}, "null"},
		{&Int64{sql.NullInt64{10, false}}, "null"},
		{&Int64{sql.NullInt64{10, true}}, "10"},
	} {
		if actually, _ := unit.value.MarshalJSON(); string(actually) != unit.expected {
			t.Errorf("combination: [%v], type:%T, actually: [%v], type: %T", unit.expected, unit.expected, string(actually), string(actually))
		}
	}
}
