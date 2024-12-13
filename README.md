# Pandora CLI

Pandora CLI is an HTTPS client tool designed to provide a simple and intuitive interface for interacting with the Pandora Secrets Management System. It allows users to perform all necessary actions related to the Pandora server through a command-line interface.

## User Guide

### Prerequisites
- Golang installed on the machine to compile the code.

### Installation
1. Clone the repository:
```shell
   git clone https://github.com/MustafaAbdulazizHamza/Pandora-CLI.git
```
2. Generate or obtain your own RSA key pair (key.pem files).
3. Compile the code:
```shell
go build -o pandora-cli
```
4. Edit the configuration file with your credentials. Leave the root credentials blank if you are not the root user.
5. Run the compiled executable. You can move it to any location, but ensure that the configuration file remains in the same directory as the executable.
### User Management
User management is a crucial part of any system. In Pandora, user management is primarily handled by the root user, with some exceptions allowing individual users to update their own passwords.
1. To add a new user:
```shell
pandora-cli user add -u username1 -p password1
```
2. To update user credentials:
```shell
pandora-cli user update -u username1 -p password2
```
3. To delete a user:
```shell
pandora-cli user delete -u username1
```
### Secret Management
Pandora was initially developed as a secrets management system, with its primary responsibility being to allow users to securely add, retrieve, update, and delete their secrets in a centralized manner through simple API calls.
1. To add a secret:
```shell
pandora-cli secret add -i secret_id -s secret1
```
2. To retrieve a secret:
```shell
pandora-cli secret get -i secret_id
```
3. To update a secret:
```shell
pandora-cli secret update -i secret_id -s secret2
```
4. To delete a secret:
```shell
pandora-cli delete -i secret_id
```
