package models

type GenerateKeyResp struct {
	Key string `json:"Key"`
}

type GenerateKeyReq struct {
	Size int    `json:"Size"`
	PRNG string `json:"PRNG"`
}

type GenerateKCVReq struct {
	Key string `json:"Key"`
}

type GenerateKCVResp struct {
	KCV string `json:"KCV"`
}
