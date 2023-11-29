package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brandon-a-pinto/nebula/user-service/configs"
	"github.com/brandon-a-pinto/nebula/user-service/internal/domain/dto"
)

var config = configs.LoadConfig()

func HttpLog(msg, typ string) error {
	dto := dto.CreateLogInput{
		Msg:  msg,
		Type: typ,
	}

	jsonValue, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("broker-service:%s/logs", config.BrokerServerPort), bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
