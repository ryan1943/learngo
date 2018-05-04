package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test.html")
	if err != nil {
		panic(err)
	}
	result := ParseCity(contents)
	const resultSize = 20

	expectedUrls := []string{
		"http://album.zhenai.com/u/108415017",
		"http://album.zhenai.com/u/1314495053",
		"http://album.zhenai.com/u/110171680",
	}
	expectedUsers := []string{
		"User 惠儿", "User 风中的蒲公英", "User 幽诺",
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url %d: %s; but was %s",
				i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d items; but had %d",
			resultSize, len(result.Items))
	}

	for i, user := range expectedUsers {
		if result.Items[i].(string) != user {
			t.Errorf("expected user %d: %s; but was %s",
				i, user, result.Items[i].(string))
		}
	}
}
