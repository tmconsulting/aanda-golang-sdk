package test

import (
	"errors"
	"github.com/nbio/st"
	"github.com/tmconsulting/aanda-sdk"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestOrderRequest_ok(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("orderRequest_answOk.txt"))

	orderReq := aandaSdk.OrderRequest{}
	data, err := zApi.OrderRequest(orderReq)

	st.Expect(t, err, nil)
	st.Expect(t, data.Status, "Ok")
}

func TestOrderRequest_err1(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("orderRequest_answErr1.txt"))

	orderReq := aandaSdk.OrderRequest{}
	_, err := zApi.OrderRequest(orderReq)

	st.Expect(t, err, errors.New("Количество гостей не равно NumberOfGuests"))
}

func TestOrderRequest_err2(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("orderRequest_answErr2.txt"))

	orderReq := aandaSdk.OrderRequest{}
	_, err := zApi.OrderRequest(orderReq)

	st.Expect(t, err, errors.New("Неверный формат DepartureDate[]"))
}

func TestOrderRequest_err3(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("orderRequest_answErr3.txt"))

	orderReq := aandaSdk.OrderRequest{}
	_, err := zApi.OrderRequest(orderReq)

	st.Expect(t, err, errors.New("Ошибка! - Номер 1642 неактивен! Выберите другой номер."))
}
