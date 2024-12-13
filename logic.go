package main

import (
	"github.com/MustafaAbdulazizHamza/Pandora-CLI/clientLogic"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

func loadConfig(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Failed to open config file: ", err.Error())
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("failed to decode config file: %s\n", err.Error())
	}

}
func getExecutableLocation() string {
	executablePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %s\n", err)
	}

	absPath, err := filepath.Abs(executablePath)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %s\n", err)
	}
	return filepath.Dir(absPath)
}

func addUser(url, authusername, authpassword, username, password string) {
	if err := clientLogic.AddUser(url, authusername, authpassword, username, password); err != nil {
		log.Fatalf("Failed to add user: %s\n", err.Error())
	}
}
func deleteUser(url, authusername, authpassword, username string) {
	if err := clientLogic.DeleteUser(url, authusername, authpassword, username); err != nil {
		log.Fatalf("Failed to delete user: %s\n", err.Error())
	}
}
func updateUserCredentials(url, authusername, authpassword, username, password string) {
	if err := clientLogic.UpdateUserCredentials(url, authusername, authpassword, username, password); err != nil {
		log.Fatalf("Failed to update user credentials: %s\n", err.Error())
	}
}
func getSecret(url, username, password, secretID, privateKey string) {
	if err := clientLogic.GetSecret(url, username, password, secretID, privateKey); err != nil {
		log.Fatalf("Failed to get secret: %s\n", err.Error())
	}
}

func addSecret(url, username, password, secretID, secret, publicKey string) {
	if err := clientLogic.PostSecret(url, username, password, secretID, secret, publicKey); err != nil {
		log.Fatalf("Failed to add secret: %s\n", err.Error())
	}
}
func updateSecret(url, username, password, secretID, secret, publicKey string) {
	if err := clientLogic.UpdateSecret(url, username, password, secretID, secret, publicKey); err != nil {
		log.Fatalf("Failed to update secret: %s\n", err.Error())
	}
}
func deleteSecret(url, username, password, secretID string) {
	if err := clientLogic.DeleteSecret(url, username, password, secretID); err != nil {
		log.Fatalf("Failed to delete secret: %s\n", err.Error())
	}
}
