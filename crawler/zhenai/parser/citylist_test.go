package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_file.html")

	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)

	const resultSize = 470

	expectedUrls := []string{
		"", "", "",
	}

	expectedCitys := []string{
		"", "", "",
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("excepted url %d : %s ;but was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d request , but had %d ", resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("items should have %d request , but had %d ", resultSize, len(result.Items))
	}

	for i, city := range expectedCitys {
		if result.Items[i].(string) != city {
			t.Errorf("excepted city %d : %s ;but was %s", i, city, result.Items[i].(string))
		}
	}
}
