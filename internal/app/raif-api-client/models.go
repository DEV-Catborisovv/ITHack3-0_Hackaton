package raifapiclient

type QRResponse struct {
	QrId             string `json:"qrId"`
	QrStatus         string `json:"qrStatus"`
	QrExpirationDate string `json:"qrExpirationDate"`
	Payload          string `json:"payload"`
	QrUrl            string `json:"qrUrl"`
	SubscriptionId   string `json:"subscriptionId"`
}
