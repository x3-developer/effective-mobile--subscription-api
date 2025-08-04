package req

import (
	"encoding/json"
	"io"
	"net/url"
)

func DecodeBody[T any](body io.ReadCloser) (T, error) {
	var payload T

	if err := json.NewDecoder(body).Decode(&payload); err != nil {
		return payload, err
	}
	return payload, nil
}

func DecodeQuery[T any](query url.Values) (T, error) {
	var payload T

	queryMap := make(map[string]string)
	for key, values := range query {
		if len(values) > 0 {
			queryMap[key] = values[0]
		}
	}

	jsonData, err := json.Marshal(queryMap)
	if err != nil {
		return payload, err
	}

	if err := json.Unmarshal(jsonData, &payload); err != nil {
		return payload, err
	}

	return payload, nil
}
