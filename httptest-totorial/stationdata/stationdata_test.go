package stationdata

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetLineInfoFromPrefCode(t *testing.T) {
	wantedResponse := `{"line": [{"line_cd": 100, "line_name": "test line"}]}`

	// 外部API代わりのHTTPサーバ
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, wantedResponse)
		},
	))
	defer ts.Close()

	// テスト実行
	// 環境変数を指定
	os.Setenv("STATIONDATA_API_DOMAIN", ts.URL)
	line, err := GetLineInfoFromPrefCode(1)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if line != wantedResponse {
		t.Errorf("invalid response. want=%s, got=%s", wantedResponse, line)
	}
}

func TestGetLineInfoFromPrefStatusCode500(t *testing.T) {

	// 外部API代わりのHTTPサーバ
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(500)
		},
	))
	defer ts.Close()

	// テスト実行
	// 環境変数を指定
	os.Setenv("STATIONDATA_API_DOMAIN", ts.URL)
	line, err := GetLineInfoFromPrefCode(1)
	if err == nil {
		t.Fatal("no errors")
	}
	if err.Error() != "statuscode=500" {
		t.Fatalf("invalid error: %v", err)
	}
	if line != "" {
		t.Errorf("invalid response. got=%s", line)
	}
}
