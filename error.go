package aandaSdk

import "fmt"

type AandaError struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Type   string `json:"type"`
	Note   string `json:"note"`
}

func (o *AandaError) Error() string {
	return fmt.Sprintf("[%s] %d - %s; %s", o.Status, o.Code, o.Type, o.Note)
}
