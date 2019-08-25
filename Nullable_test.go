package nullable

import (
	"database/sql"
	"testing"
)

func TestNullableString(t *testing.T){

	for _, unit := range []struct {
		m        *NullableString
		n        int
		expected string
	}{
	{&NullableString{sql.NullString{"aquas", true}}, 0,  "\"aquas\""},
		{&NullableString{sql.NullString{"aquas", false}}, 0,  "null"},
		{&NullableString{sql.NullString{"", false}}, 0,  "null"},
	} {
		if actually, _ := unit.m.MarshalJSON(); string(actually) != unit.expected {
			t.Errorf("combination: [%v], type:%T, actually: [%v], type: %T", unit.expected, unit.expected, string(actually), string(actually))
		}
	}
}