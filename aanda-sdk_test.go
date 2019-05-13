package aandaSdk

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/nbio/st"
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

	_, err := api.CityListRequest(context.Background(), 3)

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
		reqQuery      string
		reqMimeType   string
		reqData       []byte
		reqCtxValue   int

		respMethodName string
		respQuery      string
		respMimeType   string
		respData       []byte
		respCtxValue   int
	)

	api := NewApi(auth)
	api.url = server.URL

	api.RegisterEventHandler(BeforeRequestSend, func(ctx context.Context, methodName, query, mimeType string, data []byte) {
		reqMethodName = methodName
		reqQuery = query
		reqMimeType = mimeType

		params, _ := url.ParseQuery(string(data))

		reqData = []byte(params.Get("JSON"))

		reqCtxValue = ctx.Value("value").(int)

	}).RegisterEventHandler(AfterResponseReceive, func(ctx context.Context, methodName, query, mimeType string, data []byte) {
		respMethodName = methodName
		respQuery = query
		respMimeType = mimeType

		respData = data

		respCtxValue = ctx.Value("value").(int)

	})

	value := rand.Intn(99999) + 1

	ctx := context.WithValue(context.Background(), "value", value)

	_, err := api.OrderListRequest(ctx, OrderListRequest{
		LastName: "last name",
	})

	st.Expect(t, err, nil)

	st.Expect(t, reqMethodName, "OrderListRequest")
	st.Expect(t, reqQuery, "")
	st.Expect(t, reqMimeType, "application/x-www-form-urlencoded")
	testAuth(reqData)
	st.Expect(t, reqCtxValue, value)

	st.Expect(t, respMethodName, "OrderListRequest")
	st.Expect(t, respQuery, "")
	st.Expect(t, respMimeType, "text/json")
	st.Expect(t, string(respData), "[]")
	st.Expect(t, respCtxValue, value)
}
