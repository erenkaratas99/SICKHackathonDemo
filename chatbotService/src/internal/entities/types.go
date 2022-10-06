package entities

type MsgRequestModel struct {
	MsgBody string `json:"msgBody"`
	Name    string `json:"name"`
	SName   string `json:"SName"`
}

type SAPRespModel struct {
	ID string `json:"id"`
}
