package pallas

type AssetAudienceID string

const (
	AssetAudienceIOSRelease           AssetAudienceID = "01c1d682-6e8f-4908-b724-5501fe3f5e5c"
	AssetAudienceIOS14SecurityUpdates AssetAudienceID = "c724cb61-e974-42d3-a911-ffd4dce11eda"
	AssetAudienceIOSInternal          AssetAudienceID = "ce9c2203-903b-4fb3-9f03-040dc2202694"
	AssetAudienceIOS11Beta            AssetAudienceID = "b7580fda-59d3-43ae-9488-a81b825e3c73"
	AssetAudienceIOS12Beta            AssetAudienceID = "ef473147-b8e7-4004-988e-0ae20e2532ef"
	AssetAudienceIOS13Beta            AssetAudienceID = "d8ab8a45-ee39-4229-891e-9d3ca78a87ca"
	AssetAudienceIOS14Beta            AssetAudienceID = "84da8706-e267-4554-8207-865ae0c3a120"
	AssetAudienceIOS14PublicBeta      AssetAudienceID = "dbbb0481-d521-4cdf-a2a4-5358affc224b"
	AssetAudienceIOS15DeveloperBeta   AssetAudienceID = "ce48f60c-f590-4157-a96f-41179ca08278"

	AssetAudienceTVOSRelease AssetAudienceID = "356d9da0-eee4-4c6c-bbe5-99b60eadddf0"
	AssetAudienceTVOS11Beta  AssetAudienceID = "ebd90ea1-6216-4a7c-920e-666faccb2d50"
	AssetAudienceTVOS12Beta  AssetAudienceID = "5b220c65-fe50-460b-bac5-b6774b2ff475"
	AssetAudienceTVOS13Beta  AssetAudienceID = "975af5cb-019b-42db-9543-20327280f1b2"
	AssetAudienceTVOS14Beta  AssetAudienceID = "65254ac3-f331-4c19-8559-cbe22f5bc1a6"
	AssetAudienceTVOS15Beta  AssetAudienceID = "4d0dcdf7-12f2-4ebf-9672-ac4a4459a8bc"

	AssetAudienceWatchOSRelease AssetAudienceID = "b82fcf9c-c284-41c9-8eb2-e69bf5a5269f"
	AssetAudienceWatchOS4Beta   AssetAudienceID = "f659e06d-86a2-4bab-bcbb-61b7c60969ce"
	AssetAudienceWatchOS5Beta   AssetAudienceID = "e841259b-ad2e-4046-b80f-ca96bc2e17f3"
	AssetAudienceWatchOS6Beta   AssetAudienceID = "d08cfd47-4a4a-4825-91b5-3353dfff194f"
	AssetAudienceWatchOS7Beta   AssetAudienceID = "ff6df985-3cbe-4d54-ba5f-50d02428d2a3"
	AssetAudienceWatchOS8Beta   AssetAudienceID = "b407c130-d8af-42fc-ad7a-171efea5a3d0"

	AssetAudienceAudioOSRelease AssetAudienceID = "0322d49d-d558-4ddf-bdff-c0443d0e6fac"
	AssetAudienceAudioOS14Beta  AssetAudienceID = "b05ddb59-b26d-4c89-9d09-5fda15e99207"
	AssetAudienceAudioOS15Beta  AssetAudienceID = "58ff8d56-1d77-4473-ba88-ee1690475e40"

	AssetAudienceMacOS                AssetAudienceID = "60b55e25-a8ed-4f45-826c-c1495a4ccc65"
	AssetAudienceMacOS11CustomerBeta  AssetAudienceID = "215447a0-bb03-4e18-8598-7b6b6e7d34fd"
	AssetAudienceMacOS11DeveloperBeta AssetAudienceID = "ca60afc6-5954-46fd-8cb9-60dde6ac39fd"
	AssetAudienceMacOS11PublicBeta    AssetAudienceID = "02eb66c-8e37-451f-b0f2-ffb3e878560b"
	AssetAudienceMacOS12CustomerBeta  AssetAudienceID = "a3799e8a-246d-4dee-b418-76b4519a15a2"
	AssetAudienceMacOS12DeveloperBeta AssetAudienceID = "298e518d-b45e-4d36-94be-34a63d6777ec"
	AssetAudienceMacOS12PublicBeta    AssetAudienceID = "9f86c787-7c59-45a7-a79a-9c164b00f866"
)

var AssetAudiences = []AssetAudienceID{
	AssetAudienceIOSRelease,
	AssetAudienceIOS14SecurityUpdates,
	AssetAudienceIOSInternal,
	AssetAudienceIOS11Beta,
	AssetAudienceIOS12Beta,
	AssetAudienceIOS13Beta,
	AssetAudienceIOS14Beta,
	AssetAudienceIOS14PublicBeta,
	AssetAudienceIOS15DeveloperBeta,

	AssetAudienceTVOSRelease,
	AssetAudienceTVOS11Beta,
	AssetAudienceTVOS12Beta,
	AssetAudienceTVOS13Beta,
	AssetAudienceTVOS14Beta,
	AssetAudienceTVOS15Beta,

	AssetAudienceWatchOSRelease,
	AssetAudienceWatchOS4Beta,
	AssetAudienceWatchOS5Beta,
	AssetAudienceWatchOS6Beta,
	AssetAudienceWatchOS7Beta,
	AssetAudienceWatchOS8Beta,

	AssetAudienceAudioOSRelease,
	AssetAudienceAudioOS14Beta,
	AssetAudienceAudioOS15Beta,

	AssetAudienceMacOS,
	AssetAudienceMacOS11CustomerBeta,
	AssetAudienceMacOS11DeveloperBeta,
	AssetAudienceMacOS11PublicBeta,
	AssetAudienceMacOS12CustomerBeta,
	AssetAudienceMacOS12DeveloperBeta,
	AssetAudienceMacOS12PublicBeta,
}

type AssetType string

const (
	AssetTypeMobileSoftwareUpdate AssetType = "com.apple.MobileAsset.SoftwareUpdate"
	AssetTypeMacSoftwareUpdate    AssetType = "com.apple.MobileAsset.MacSoftwareUpdate"
	AssetTypeSFRSoftwareUpdate    AssetType = "com.apple.MobileAsset.SFRSoftwareUpdate"
)

var AssetTypes = []AssetType{
	AssetTypeMobileSoftwareUpdate,
	AssetTypeMacSoftwareUpdate,
	AssetTypeSFRSoftwareUpdate,
}
