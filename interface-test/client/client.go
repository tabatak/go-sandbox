package client

type APIClient interface {
	Get(key string) string
}

type APIClientImpl struct {
}

func (rc *APIClientImpl) Get(key string) string {
	// 外部APIを利用して値を取得
	value := "some value"

	return value
}
