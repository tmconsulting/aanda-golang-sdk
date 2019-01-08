package aandaSdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/go-playground/validator.v2"
)

func validateStruct(i interface{}) error {
	if err := validator.ValidateStruct(i); err != nil {
		return err
	}
	return nil
}

func parseResp(body []byte, resp interface{}, validate func(interface{}) error, status func() *string) error {
	var (
		err  error
		vErr error
	)

	err = json.Unmarshal(body, resp)
	if validate != nil {
		vErr = validate(resp)
	}

	if err != nil || vErr != nil || (status != nil && status() != nil) {
		respErr := parseError(body)

		if respErr == nil {
			if err != nil {
				return err
			} else {
				return vErr
			}
		} else {
			return respErr
		}
	}

	return nil
}

func parseError(body []byte) error {
	// Try to parse known struct
	var aandaErr AandaError
	err := json.Unmarshal(body, &aandaErr)
	if err == nil && !aandaErr.IsEmpty() {
		return aandaErr.ToError()
	}

	var aandaErrMsg AandaErrorMsg
	err = json.Unmarshal(body, &aandaErrMsg)
	if err == nil && !aandaErrMsg.IsEmpty() {
		return aandaErrMsg.ToError()
	}

	//Try parse as JSON
	var answM map[string]interface{}
	err = json.Unmarshal(body, &answM)
	if err == nil {
		//if have only error key
		if err, ok := answM["error"].(string); ok {
			return errors.New(err)
		}
		//if have error and message key
		_, ok1 := answM["error"].(float64)
		message, ok2 := answM["message"].(string)
		if ok1 && ok2 {
			return errors.New(message)
		}
		//if have Status=Error
		st, ok3 := answM["Status"].(string)
		note, ok4 := answM["note"].(string)
		if ok3 && ok4 && st == "Error" {
			return errors.New(note)
		}
	}
	//may be simple string, try parse it
	var answS string
	err = json.Unmarshal(body, &answS)
	if err == nil {
		if strings.Index(answS, "Ошибка авторазации") == -1 {
			return errors.New("Authorization error")
		} else {
			return errors.New(answS)
		}
	}
	//may be part of XML, try parse in by regexp
	re := regexp.MustCompile(`Note="(.*)"`)
	res := re.FindStringSubmatch(string(body))
	if len(res) != 0 {
		return errors.New(res[1])
	}

	return nil
}

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
