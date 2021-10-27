package pallas

type AssetAudienceID string

const (
	AssetAudienceIOSRelease           AssetAudienceID = "01c1d682-6e8f-4908-b724-5501fe3f5e5c"
	AssetAudienceIOS14SecurityUpdates AssetAudienceID = "c724cb61-e974-42d3-a911-ffd4dce11eda"
)

type AssetType string

const (
	AssetTypeMobileSoftwareUpdate AssetType = "com.apple.MobileAsset.SoftwareUpdate"
	AssetTypeMacSoftwareUpdate    AssetType = "com.apple.MobileAsset.MacSoftwareUpdate"
	AssetTypeSFRSoftwareUpdate    AssetType = "com.apple.MobileAsset.SFRSoftwareUpdate"
)
