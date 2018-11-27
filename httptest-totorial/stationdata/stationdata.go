package stationdata

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// GetLineInfoFromPrefCode は都道府県コードをもとに沿線情報を取得する。
func GetLineInfoFromPrefCode(prefCode int) (string, error) {
	url := fmt.Sprintf("%s/api/p/%d.json", os.Getenv("STATIONDATA_API_DOMAIN"), prefCode)

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	if res.StatusCode != 200 {
		return "", fmt.Errorf("statuscode=%d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	return string(body), nil
}
