package entities

type Msg struct {
	ID  string       `json:"id"bson:"_id"`
	Msg *ChatbotResp `json:"msgBody"`
}

type Keyword struct {
	ID        string `json:"id"bson:"_id"`
	Keyword   string `json:"keyword"`
	CreatedAt string `json:"createdAt"`
}
