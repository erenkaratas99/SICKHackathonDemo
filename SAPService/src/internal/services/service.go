package services

import (
	"SICKHackathon/SAPService/src/internal/kafka/producer"
	"SICKHackathon/SAPService/src/internal/repositories"
	"SICKHackathon/shared/helper"
	"SICKHackathon/shared/types"
	"os/exec"
	"time"
)

type Service struct {
	repo     *repositories.Repository
	client   *RestClient
	producer *producer.KafkaProducer
}

func NewService(r *repositories.Repository, client *RestClient, producer *producer.KafkaProducer) *Service {
	return &Service{repo: r, client: client, producer: producer}
}

func (s *Service) PreNLPService(cbResp *types.MsgCommModel) ([][]string, error) {

	err := helper.CreateFile("raw.txt", cbResp.MsgBody)
	time.Sleep(1 * time.Second)
	//exec.Command("C://Users//KULLANICI//Desktop//untitled0.py").Run()
	helper.RunPython()
	time.Sleep(25 * time.Second)
	keyword, err := helper.ReadFile("keywords.txt")
	err = s.repo.WriteKeyword(*keyword)
	if err != nil {
		return nil, err
	}
	//keyword := cbResp.MsgBody
	byteform := []byte(*keyword)
	err = s.producer.ProduceMsg("topic_0", byteform)
	if err != nil {
		return nil, err
	}
	err = s.client.DoGetReqToSAPCloud("https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PURCHASEORDER_PROCESS_SRV/A_PurchaseOrder?$top=50&$inlinecount=allpages")
	if err != nil {
		return nil, err
	}
	time.Sleep(5 * time.Second)
	exec.Command("script.py").Run()
	csv, err := helper.ReadCsvFile("answer.csv")
	if err != nil {
		return nil, err
	}
	return csv, nil
}

//buraya iki tane private method yaz ==> python dosyasına girdi versin txt olarak
//diğeri de işlenmiş veriyi alsın
