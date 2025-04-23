package models

type GenerateKeyResp struct {
	Key string `json:"Key"`
}

type GenerateKeyReq struct {
	Size int    `json:"Size"`
	PRNG string `json:"PRNG"`
}

type GenerateKCVReq struct {
	Key    string `json:"Key"`
	Mode   string `json:"Mode"`
	Cipher string `json:"Cipher"`
}

type GenerateKCVResp struct {
	KCV string `json:"KCV"`
}
