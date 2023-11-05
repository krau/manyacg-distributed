package lskypro


type commonResp struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}


type tokensResp struct {
	commonResp
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
}