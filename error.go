package aandaSdk

import (
	"errors"
	"fmt"
	"strings"
)

type AandaError struct {
	Status string `json:"status" json:"Status"`
	Code   int    `json:"code" json:"Code"`
	Type   string `json:"type" json:"Type"`
	Note   string `json:"note" json:"Note"`
}

func (o *AandaError) IsEmpty() bool {
	return o.Status == "" && o.Code == 0 && o.Type == "" && o.Note == ""
}

func (o *AandaError) ToError() error {
	msgs := []string{}
	for _, s := range []string{o.Type, o.Note} {
		if s != "" {
			msgs = append(msgs, s)
		}
	}

	return errors.New(strings.Join(msgs, "; "))
}

func (o *AandaError) Error() string {
	return fmt.Sprintf("[%s] %d - %s; %s", o.Status, o.Code, o.Type, o.Note)
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
