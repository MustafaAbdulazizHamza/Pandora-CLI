package clientLogic

import "fmt"

func AddUser(url, authusername, authpassword, username, password string) error {
	err := userOperation(url, authusername, authpassword, username, password, "POST")
	return err
}
func DeleteUser(url, authusername, authpassword, username string) error {
	err := userOperation(url, authusername, authpassword, username, " ", "DELETE")
	return err

}
func UpdateUserCredentials(url, authusername, authpassword, username, password string) error {
	err := userOperation(url, authusername, authpassword, username, password, "PATCH")
	return err

}
func userOperation(url, authusername, authpassword, username, password, operation string) error {
	user := User{
		Username: username,
		Password: password,
	}
	body, err := toJson(user)
	if err != nil {
		return err
	}
	res, err := sendHTTPRequest(operation, url+"user", authusername, authpassword, string(body))
	if err != nil {
		return err
	}
	response, err := ParseResponse(res)
	if err != nil {
		fmt.Println("f")
		return err
	}
	PrintOutResponse(response)
	return nil

}
