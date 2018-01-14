package test

import (
	"errors"
	"github.com/nbio/st"
	"github.com/tmconsulting/aanda-golang-sdk"
	"testing"
)

func TestOrderRequest_ok(t *testing.T) {
	testRequest("orderRequest_answOk.txt")
	orderReq := aandaSdk.OrderRequest{}
	data, err := zApi.OrderRequest(orderReq)

	st.Expect(t, err, nil)
	st.Expect(t, data.Status, "Ok")
}

func TestOrderRequest_err1(t *testing.T) {
	testRequest("orderRequest_answErr1.txt")
	orderReq := aandaSdk.OrderRequest{}
	_, err := zApi.OrderRequest(orderReq)

	st.Expect(t, err, errors.New("Количество гостей не равно NumberOfGuests"))
}

func TestOrderRequest_err2(t *testing.T) {
	testRequest("orderRequest_answErr2.txt")
	orderReq := aandaSdk.OrderRequest{}
	_, err := zApi.OrderRequest(orderReq)

	st.Expect(t, err, errors.New("Неверный формат DepartureDate[]"))
}

func TestOrderRequest_err3(t *testing.T) {
	testRequest("orderRequest_answErr3.txt")
	orderReq := aandaSdk.OrderRequest{}
	_, err := zApi.OrderRequest(orderReq)

	st.Expect(t, err, errors.New("Ошибка! - Номер 1642 неактивен! Выберите другой номер."))
}

func TestOrderRequest_err4(t *testing.T) {
	testRequest("orderRequest_answErr4.txt")
	orderReq := aandaSdk.OrderRequest{}
	_, err := zApi.OrderRequest(orderReq)

	st.Expect(t, err, errors.New("Не верный статус"))
}
