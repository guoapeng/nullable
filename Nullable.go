package nullable

import (
	"database/sql"
	"encoding/json"
)

type NullableString struct {
	sql.NullString
}

func (value *NullableString) MarshalJSON() ([]byte, error) {
	if !value.Valid {
		return json.Marshal(nil)
	} else {
		return json.Marshal(value.String)
	}
}
