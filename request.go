package deepl

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const url = "https://www.deepl.com/jsonrpc"

func request(c call, timeout time.Duration) (reply, error) {
	jsonStr, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return reply{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cache-Control", "no-cache")

	client := &http.Client{Timeout: timeout}
	res, err := client.Do(req)
	if err != nil {
		return reply{}, err
	}
	defer res.Body.Close()

	jsonStr, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return reply{}, err
	}

	r := reply{}
	if err = json.Unmarshal(jsonStr, &r); err != nil {
		return reply{}, err
	}

	return r, nil
}
