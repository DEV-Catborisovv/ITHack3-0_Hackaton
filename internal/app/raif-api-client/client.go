package raifapiclient

import "net/http"

func GetNewClient() *http.Client {
	client := http.Client{}
	return &client
}
