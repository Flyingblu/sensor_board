package influxclient

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type InfluxClient struct {
	InfluxUrl string
	Token     string
	Database  string
}

func (ifClient *InfluxClient) ReceiveData(token, device, measurement string, value float64) error {

	if ifClient.Token != token {
		return errors.New("Incorrect token! ")
	}

	writeData := fmt.Sprintf("%v,device=%v value=%v", measurement, device, value)
	reader := strings.NewReader(writeData)
	fmt.Println("Hello!")
	_, err := http.Post(ifClient.InfluxUrl+"/write?db="+ifClient.Database, "", reader)

	if err != nil {
		return err
	}
	return nil
}
