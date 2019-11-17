package fetcher

import (
	"io/ioutil"
	"net/http"
	"fmt"
)

const URL  ="\"http://www.zhenai.com/zhenghun"
func Fetch(url string) ([]byte,  error) {
	resp,err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("wrong status code: %d ",resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

