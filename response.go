package pallas

type Response struct {
	AssetSetID      string          `json:"AssetSetId"`
	Assets          []AssetResponse `json:"Assets"`
	LegacyXMLURL    string          `json:"LegacyXmlUrl"`
	Nonce           string          `json:"Nonce"`
	PallasNonce     string          `json:"PallasNonce"`
	PostingDate     string          `json:"PostingDate"`
	SessionID       string          `json:"SessionId"`
	Transformations Transformations `json:"Transformations"`
}

type AssetResponse struct {
	ActualMinimumSystemPartition          int64          `json:"ActualMinimumSystemPartition"`
	AssetType                             string         `json:"AssetType"`
	Build                                 string         `json:"Build"`
	InstallationSize                      string         `json:"InstallationSize"`
	InstallationSizeSnapshot              string         `json:"InstallationSize-Snapshot"`
	MinimumSystemPartition                int64          `json:"MinimumSystemPartition"`
	OSVersion                             string         `json:"OSVersion"`
	PrerequisiteBuild                     string         `json:"PrerequisiteBuild"`
	PrerequisiteOSVersion                 string         `json:"PrerequisiteOSVersion"`
	RSEPDigest                            string         `json:"RSEPDigest"`
	Ramp                                  bool           `json:"Ramp"`
	RescueMinimumSystemPartition          int64          `json:"RescueMinimumSystemPartition"`
	SEPDigest                             string         `json:"SEPDigest"`
	SUConvReqd                            bool           `json:"SUConvReqd"`
	SUDocumentationID                     string         `json:"SUDocumentationID"`
	SUInstallTonightEnabled               bool           `json:"SUInstallTonightEnabled"`
	SUMultiPassEnabled                    bool           `json:"SUMultiPassEnabled"`
	SUProductSystemName                   string         `json:"SUProductSystemName"`
	SUPublisher                           string         `json:"SUPublisher"`
	SupportedDeviceModels                 []string       `json:"SupportedDeviceModels"`
	SupportedDevices                      []string       `json:"SupportedDevices"`
	SystemPartitionPadding                map[string]int `json:"SystemPartitionPadding"`
	SystemVolumeSealingOverhead           int64          `json:"SystemVolumeSealingOverhead"`
	AssetReceipt                          AssetReceipt   `json:"_AssetReceipt"`
	CompressionAlgorithm                  string         `json:"_CompressionAlgorithm"`
	DownloadSize                          int64          `json:"_DownloadSize"`
	EventRecordingServiceURL              string         `json:"_EventRecordingServiceURL"`
	IsZipStreamable                       bool           `json:"_IsZipStreamable"`
	Measurement                           string         `json:"_Measurement"`
	MeasurementAlgorithm                  string         `json:"_MeasurementAlgorithm"`
	UnarchivedSize                        int64          `json:"_UnarchivedSize"`
	AssetDefaultGarbageCollectionBehavior string         `json:"__AssetDefaultGarbageCollectionBehavior"`
	BaseURL                               string         `json:"__BaseURL"`
	CanUseLocalCacheServer                bool           `json:"__CanUseLocalCacheServer"`
	HideInstallAlert                      bool           `json:"__HideInstallAlert"`
	QueuingServiceURL                     string         `json:"__QueuingServiceURL"`
	RelativePath                          string         `json:"__RelativePath"`
}

type AssetReceipt struct {
	AssetReceipt   string `json:"AssetReceipt"`
	AssetSignature string `json:"AssetSignature"`
}

type Transformations struct {
	RSEPDigest  string `json:"RSEPDigest"`
	SEPDigest   string `json:"SEPDigest"`
	Measurement string `json:"_Measurement"`
}
