package pallas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const assetSetsURL = "https://gdmf.apple.com/v2/pmv"

type AssetSetsResponse struct {
	PublicAssetSets *AssetSets `json:"PublicAssetSets"`
	AssetSets       *AssetSets `json:"AssetSets"`
}

type AssetSets struct {
	IOS   []AssetSet `json:"iOS"`
	MacOS []AssetSet `json:"macOS"`
}

type AssetSet struct {
	ProductVersion   string
	PostingDate      AssetSetTime
	ExpirationDate   AssetSetTime
	SupportedDevices []string
}

func GetAssetSets(httpClient *http.Client) (*AssetSetsResponse, error) {
	resp, err := httpClient.Get(assetSetsURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("pallas: invalid http status code, got: %d", resp.StatusCode)
	}

	var assetSets AssetSetsResponse

	if err := json.NewDecoder(resp.Body).Decode(&assetSets); err != nil {
		return nil, err
	}

	return &assetSets, nil
}

type AssetSetTime time.Time

func (a *AssetSetTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)

	if err != nil {
		return err
	}

	*a = AssetSetTime(t)

	return nil
}

func (a *AssetSetTime) Time() time.Time {
	if a == nil {
		return time.Time{}
	}

	return time.Time(*a)
}
