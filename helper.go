package aandaSdk

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type MustString string

func (o *MustString) UnmarshalJSON(data []byte) error {
	var v interface{}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v := v.(type) {
	case string:
		*o = MustString(v)
	case int:
		*o = MustString(strconv.Itoa(v))
	default:
		*o = MustString(fmt.Sprint(v))
	}

	return nil
}
