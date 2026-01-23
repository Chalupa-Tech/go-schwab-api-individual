package models

type UserPreference struct {
	Accounts     []UserPreferenceAccount `json:"accounts"`
	StreamerInfo []StreamerInfo          `json:"streamerInfo"`
	Offers       []Offer                 `json:"offers"`
}

type UserPreferenceAccount struct {
	AccountNumber      string `json:"accountNumber"`
	PrimaryAccount     bool   `json:"primaryAccount"`
	Type               string `json:"type"`
	NickName           string `json:"nickName"`
	AccountColor       string `json:"accountColor"`
	DisplayAcctId      string `json:"displayAcctId"`
	AutoPositionEffect bool   `json:"autoPositionEffect"`
}

type StreamerInfo struct {
	StreamerSocketUrl      string `json:"streamerSocketUrl"`
	SchwabClientCustomerId string `json:"schwabClientCustomerId"`
	SchwabClientCorrelId   string `json:"schwabClientCorrelId"`
	SchwabClientChannel    string `json:"schwabClientChannel"`
	SchwabClientFunctionId string `json:"schwabClientFunctionId"`
}

type Offer struct {
	Level2Permissions bool   `json:"level2Permissions"`
	MktDataPermission string `json:"mktDataPermission"`
}
