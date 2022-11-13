package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"shortlink/internal/models"
)

func UnmarshalRequest(ctx context.Context, r *http.Request, target *models.ShortLinkRedisData) error {

	if r.Body == nil {
		return errors.New("empty request sent")
	}

	data, errReadData := ioutil.ReadAll(r.Body)
	if errReadData != nil {
		return errors.New("unable to read data")
	}

	err := json.Unmarshal(data, target)
	if err != nil {
		return err
	}

	return nil
}

// Validate URL
func IsValidUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Host != ""
}

func MarshalRequest(ctx context.Context, data *models.ShortLinkRedisData) ([]byte, error) {
	respJson, err := json.Marshal(data)
	if err != nil {
		fmt.Print("Utils Error : Unable to marshal json response")
		return nil, err
	}

	return respJson, nil
}
