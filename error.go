package aandaSdk

import (
	"errors"
	"fmt"
	"strings"
)

type AandaError struct {
	Status string     `json:"status" json:"Status"`
	Code   MustString `json:"code" json:"Code"`
	Type   string     `json:"type" json:"Type"`
	Note   string     `json:"note" json:"Note"`
}

func (o *AandaError) IsEmpty() bool {
	return o.Status == "" && (o.Code == "" || o.Code == "0") && o.Type == "" && o.Note == ""
}

func (o *AandaError) ToError() error {
	var msgs []string
	for _, s := range []string{o.Type, o.Note} {
		if s != "" {
			msgs = append(msgs, s)
		}
	}

	return errors.New(strings.Join(msgs, "; "))
}

func (o *AandaError) Error() string {
	return fmt.Sprintf("[%s] %s - %s; %s", o.Status, o.Code, o.Type, o.Note)
}

type AandaErrorMsg struct {
	Err string `json:"error"`
}

func (o *AandaErrorMsg) IsEmpty() bool {
	return o.Err == ""
}

func (o *AandaErrorMsg) ToError() error {
	return errors.New(o.Err)
}

func (o *AandaErrorMsg) Error() string {
	return o.Err
}
