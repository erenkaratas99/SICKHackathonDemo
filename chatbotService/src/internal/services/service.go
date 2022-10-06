package services

import (
	"SICKHackathon/chatbotService/src/internal/repositories"
	"SICKHackathon/shared/types"
	"bytes"
	"encoding/json"
	"github.com/erenkaratas99/COApiCore/pkg/customErrors"
)

type Service struct {
	repo   *repositories.Repository
	client *RestClient
}

func NewService(r *repositories.Repository, client *RestClient) *Service {
	return &Service{repo: r, client: client}
}

func (s *Service) SendMessageService(msgReq *types.MsgCommModel) error {
	reqBodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBytes).Encode(msgReq)
	if err != nil {
		return customErrors.BindErr
	}
	err = s.client.MakePostRequest("http://127.0.0.1:8001/sap-service/pushmsg", reqBodyBytes.Bytes())
	if err != nil {
		return err
	}
	err = s.repo.WriteMSG(msgReq)
	if err != nil {
		return err
	}
	return nil
}
