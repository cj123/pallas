package main

import (
	"fmt"
	"net/http"

	"github.com/cj123/pallas"
)

func main() {
	r := &pallas.Request{
		ClientVersion:           pallas.ClientVersion,
		AssetType:               pallas.AssetTypeMobileSoftwareUpdate,
		AssetAudience:           pallas.AssetAudienceIOS14SecurityUpdates,
		ProductType:             "iPhone10,6",
		HWModelStr:              "D221AP",
		ProductVersion:          "14.8",
		BuildVersion:            "18H17",
		RequestedProductVersion: "15.0",
	}

	resp, err := pallas.MakeRequest(r, http.DefaultClient)

	if err != nil {
		panic(err)
	}

	for _, asset := range resp.Assets {
		fmt.Printf("%s (%s) - Download: %s\n", asset.OSVersion, asset.Build, asset.BaseURL+asset.RelativePath)
	}
}
