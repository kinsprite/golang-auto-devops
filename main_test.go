package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mainRouter(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"healthz", args{"healthz"}, `{"message":"OK"}`},
		{"ping", args{"ping"}, `{"message":"pong"}`},
		{"hello", args{"hello"}, `{"message":"world"}`},
		{"add(2, 8)", args{"add"}, `{"result":10}`},
	}

	router := mainRouter()

	ts := httptest.NewServer(router)
	defer ts.Close()

	getResp := func(uri string) string {
		resp, err := http.Get(fmt.Sprintf("%s/%s", ts.URL, uri))
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		return string(body)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := getResp(tt.args.uri); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("mainRouter() = %v, want %v", got, tt.want)
			// }

			got := getResp(tt.args.uri)
			assert.Equal(t, tt.want, got)
		})
	}
}
