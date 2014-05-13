//
// This file simplies the json Encoding and Decoding of sql-safe values
// By using the String and Int64 structs declared below, the MarshalJSON
// function has been overwritten to facilitate cleaner json encoding for
// the rest api.
//

package libraries

import (
	"database/sql"
	"fmt"
)

type String struct {
	sql.NullString
}

func (s String) MarshalJSON() ([]byte, error) {
	if s.Valid == false {
		return []byte("null"), nil
	} else {
		return []byte("\"" + s.String + "\""), nil
	}
}

type Int64 struct {
	sql.NullInt64
}

func (s Int64) MarshalJSON() ([]byte, error) {
	if s.Valid == false {
		return nil, nil
	} else {
		return []byte(fmt.Sprintf("%d", s.Int64)), nil
	}
}
