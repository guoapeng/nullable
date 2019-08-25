package nullable

import (
	"database/sql"
	"encoding/json"
)

type String struct {
	sql.NullString
}

func (value *String) MarshalJSON() ([]byte, error) {
	if !value.Valid {
		return json.Marshal(nil)
	} else {
		return json.Marshal(value.String)
	}
}

type Int64 struct {
	sql.NullInt64
}

func (value *Int64) MarshalJSON() ([]byte, error) {
	if !value.Valid {
		return json.Marshal(nil)
	} else {
		return json.Marshal(value.Int64)
	}
}

type Float64 struct {
	sql.NullFloat64
}

func (value *Float64) MarshalJSON() ([]byte, error) {
	if !value.Valid {
		return json.Marshal(nil)
	} else {
		return json.Marshal(value.Float64)
	}
}

type Bool struct {
	sql.NullBool
}

func (value *Bool) MarshalJSON() ([]byte, error) {
	if !value.Valid {
		return json.Marshal(nil)
	} else {
		return json.Marshal(value.Bool)
	}
}
