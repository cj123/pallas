package pallas

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	pallasRequestURL = "https://gdmf.apple.com/v2/assets"

	ClientVersion = 2
)

type Request struct {
	ClientVersion           int             `json:"ClientVersion"`
	AssetType               AssetType       `json:"AssetType"`
	AssetAudience           AssetAudienceID `json:"AssetAudience"`
	ProductType             string          `json:"ProductType"`
	HWModelStr              string          `json:"HWModelStr"`
	ProductVersion          string          `json:"ProductVersion"`
	BuildVersion            string          `json:"BuildVersion"`
	RequestedProductVersion string          `json:"RequestedProductVersion,omitempty"`
}

func MakeRequest(req *Request, httpClient *http.Client) (*Response, error) {
	requestData := new(bytes.Buffer)

	if err := json.NewEncoder(requestData).Encode(req); err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, pallasRequestURL, requestData)

	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Accept", "application/json")
	httpReq.Header.Add("User-Agent", "pallasutil")

	resp, err := httpClient.Do(httpReq)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("pallas: invalid http status code, got: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	parts := bytes.Split(body, []byte("."))

	if len(parts) == 0 {
		return nil, ErrInvalidResponseFormat
	}

	trimmed := parts[1]
	trimmed = bytes.ReplaceAll(trimmed, []byte("-"), []byte("+"))
	trimmed = bytes.ReplaceAll(trimmed, []byte("_"), []byte("/"))

	out := make([]byte, base64.StdEncoding.WithPadding(base64.NoPadding).DecodedLen(len(trimmed)))

	_, err = base64.StdEncoding.WithPadding(base64.NoPadding).Decode(out, trimmed)

	if err != nil {
		return nil, err
	}

	var response Response

	if err := json.Unmarshal(out, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

var ErrInvalidResponseFormat = errors.New("pallas: invalid response format")
