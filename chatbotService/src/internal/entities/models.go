package entities

type Message struct {
	UID       string `json:"id"bson:"_id"`
	Message   string `json:"msgBody"`
	Name      string `json:"name"`
	SName     string `json:"s_name"`
	CreatedAt string `bson:"created_at"json:"createdAt"`
}
