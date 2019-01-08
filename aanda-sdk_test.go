package aandaSdk

import (
	"encoding/json"
	"github.com/nbio/st"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestApi_CityListRequest(t *testing.T) {
	auth := Auth{
		BuyerId:  "B1",
		UserId:   "U1",
		Password: "P1",
		Language: "L1",
	}

	server := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		body, _ := json.Marshal(&CityListResponse{})
		resp.Write(body)
	}))
	defer server.Close()

	api := NewApi(auth)
	api.url = server.URL

	_, err := api.CityListRequest(3)

	st.Expect(t, err, nil)
}

func TestApi_OrderListRequest(t *testing.T) {
	auth := Auth{
		BuyerId:  "B1",
		UserId:   "U1",
		Password: "P1",
		Language: "L1",
	}

	testAuth := func(data []byte) {
		reqParam := OrderListRequest{}

		json.Unmarshal(data, &reqParam)

		existAuth := Auth{
			BuyerId:  reqParam.BuyerId,
			UserId:   reqParam.UserId,
			Password: reqParam.Password,
			Language: reqParam.Language,
		}

		if !reflect.DeepEqual(auth, existAuth) {
			t.Error("failed to set auth")
		}
	}

	server := httptest.NewServer(http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		req.ParseForm()

		testAuth([]byte(req.Form.Get("JSON")))

		body, _ := json.Marshal(&[]OrderListResponse{})

		resp.Header().Set("Content-Type", "text/json")
		resp.Write(body)
	}))
	defer server.Close()

	var (
		reqMethodName string
		reqMimeType   string
		reqData       []byte

		respMethodName string
		respMimeType   string
		respData       []byte
	)

	api := NewApi(auth)
	api.url = server.URL

	api.RegisterEventHandler(BeforeRequestSend, func(methodName, mimeType string, data []byte) {
		reqMethodName = methodName
		reqMimeType = mimeType

		params, _ := url.ParseQuery(string(data))

		reqData = []byte(params.Get("JSON"))

	}).RegisterEventHandler(AfterResponseReceive, func(methodName, mimeType string, data []byte) {
		respMethodName = methodName
		respMimeType = mimeType

		respData = data

	})

	_, err := api.OrderListRequest(OrderListRequest{
		LastName: "last name",
	})

	st.Expect(t, err, nil)

	st.Expect(t, reqMethodName, "OrderListRequest")
	st.Expect(t, reqMimeType, "application/x-www-form-urlencoded")
	testAuth(reqData)

	st.Expect(t, respMethodName, "OrderListRequest")
	st.Expect(t, respMimeType, "text/json")
	st.Expect(t, string(respData), "[]")
}
