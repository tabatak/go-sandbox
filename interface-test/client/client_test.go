package client

import "testing"

type MockAPIClient struct {
	GetFunc func(key string) string
}

func (c *MockAPIClient) Get(key string) string {
	return c.GetFunc(key)
}

func TestGetReturnCase1(t *testing.T) {
	retStr := "case1"

	// モッククライアント
	client := &MockAPIClient{
		GetFunc: func(key string) string {
			// 常に指定した文字列を返す
			return retStr
		},
	}

	got := client.Get("testkey")
	if retStr != got {
		t.Errorf("want=%s, got=%s", retStr, got)
	}
}

func TestGetReturnCase2(t *testing.T) {
	retStr := "case2"

	// モッククライアント
	client := &MockAPIClient{
		GetFunc: func(key string) string {
			// 常に指定した文字列を返す
			return retStr
		},
	}

	got := client.Get("testkey")
	if retStr != got {
		t.Errorf("want=%s, got=%s", retStr, got)
	}
}
