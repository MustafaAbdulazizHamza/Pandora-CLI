package clientLogic

func secretGetDelete(url, username, password, secretID, privateKey string, isDelete bool) error {
	req := RequestedSecret{
		SecretID: secretID,
	}
	body, err := toJson(req)
	if err != nil {
		return err
	}
	method := "GET"
	if isDelete {
		method = "DELETE"
	}
	res, err := sendHTTPRequest(method, url+"secret", username, password, string(body))
	if err != nil {
		return err
	}
	response, err := ParseResponse(res)
	if err != nil {
		return err
	}
	if !isDelete && response.Status == "200" {
		response.Text, err = decryptWithPrivateKey(response.Text, privateKey)
		if err != nil {
			return err
		}
	}
	PrintOutResponse(response)
	return nil

}

func secretPostUpdate(url, username, password, secret, secretID, publicKey string, isUpdate bool) error {
	secret, err := encryptWithPublicKey(secret, publicKey)
	if err != nil {
		return err
	}
	sec := Secret{SecretID: secretID, Secret: secret}
	body, err := toJson(sec)
	if err != nil {
		return err
	}
	method := "POST"
	if isUpdate {
		method = "PATCH"
	}
	res, err := sendHTTPRequest(method, url+"secret", username, password, string(body))
	if err != nil {
		return err
	}
	response, err := ParseResponse(res)
	if err != nil {
		return err
	}

	PrintOutResponse(response)
	return nil
}
func GetSecret(url, username, password, secretID, privateKey string) error {
	err := secretGetDelete(url, username, password, secretID, privateKey, false)
	return err
}

func PostSecret(url, username, password, secretID, secret, publicKey string) error {
	err := secretPostUpdate(url, username, password, secret, secretID, publicKey, false)
	return err

}
func DeleteSecret(url, username, password, secretID string) error {
	err := secretGetDelete(url, username, password, secretID, "", true)
	return err
}
func UpdateSecret(url, username, password, secretID, secret, publicKey string) error {
	err := secretPostUpdate(url, username, password, secret, secretID, publicKey, true)
	return err
}
