package raifapiclient_methods

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tgBotHakaton/configs"
	rac "tgBotHakaton/internal/app/raif-api-client"
)

func GetOrderDataFromUid(uid string) (error, *rac.QRResponse) {
	config := configs.NewConfig()

	client := rac.GetNewClient()
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", rac.GetOrderDataFromUidURL, uid), nil)
	if err != nil {
		return err, nil
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.APIClient.SecretKey))

	resp, err := client.Do(req)
	if err != nil {
		return err, nil
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Error with status code: %d, expected status-code 200", resp.StatusCode), nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}

	respJson := rac.QRResponse{}
	err = json.Unmarshal(body, &respJson)
	if err != nil {
		return err, nil
	}

	return nil, &respJson
}
