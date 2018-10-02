package aandaSdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MustString string

func (o *MustString) UnmarshalJSON(data []byte) error {
	var v interface{}

	if string(data) == "null" {
		return nil
	}

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

type MustFloat64 float64

func (o *MustFloat64) UnmarshalJSON(data []byte) error {
	var v interface{}

	if string(data) == "null" {
		return nil
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v := v.(type) {
	case float64:
		*o = MustFloat64(v)
	case float32:
		*o = MustFloat64(v)
	case int64:
		*o = MustFloat64(v)
	case int:
		*o = MustFloat64(v)
	case string:
		f, _ := strconv.ParseFloat(strings.Trim(v, `"`), 64)
		*o = MustFloat64(f)
	default:
		return errors.New(fmt.Sprintf("error unmarshal %#v as float64", v))
	}

	return nil
}

type MustInt int

func (o *MustInt) UnmarshalJSON(data []byte) error {
	var v interface{}

	if string(data) == "null" {
		return nil
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch v := v.(type) {
	case float64:
		*o = MustInt(v)
	case float32:
		*o = MustInt(v)
	case int64:
		*o = MustInt(v)
	case int:
		*o = MustInt(v)
	case string:
		f, _ := strconv.Atoi(v)
		*o = MustInt(f)
	default:
		return errors.New(fmt.Sprintf("error unmarshal %#v as int", v))
	}

	return nil
}
