import os
import base64
import requests
import json
from cryptography.hazmat.primitives.asymmetric import padding as asymmetric_padding
from cryptography.hazmat.primitives import hashes
from cryptography.hazmat.primitives.serialization import load_pem_private_key, load_pem_public_key
import warnings
import sys
warnings.filterwarnings("ignore")

class RSAKeysNotFound(Exception):
    pass


class Pandora:
    """
A Python HTTPS client for the Pandora Secrets Management System. It can be used to retrieve secrets from or post secrets to Pandora.
    """

    def __init__(self, url: str, username: str, password: str, privateKeyPath: str, publicKeyPath: str) -> None:
        if not url.endswith("/"):
            url = url + "/"
        self._url = url
        self._auth_headers = {
            "username": username,
            "password": password
        }
        if not (os.path.isfile(privateKeyPath) and os.path.isfile(publicKeyPath)):
            raise RSAKeysNotFound("RSA keys were not found.")
        self._privateKey = privateKeyPath
        self._publicKey = publicKeyPath

    def _RSA_Encrypt(self, message: str) -> str:
        """Encrypt a message with the RSA public key."""
        with open(self._publicKey, "rb") as key_file:
            public_key = load_pem_public_key(key_file.read())

        ciphertext = public_key.encrypt(
            message.encode("utf-8"),
            asymmetric_padding.OAEP(
                mgf=asymmetric_padding.MGF1(algorithm=hashes.SHA256()),
                algorithm=hashes.SHA256(),
                label=None
            )
        )

        return base64.b64encode(ciphertext).decode('utf-8')

    def _RSA_Decrypt(self, encrypted_message: str) -> str:
        """Decrypt a message with the RSA private key."""
        with open(self._privateKey, "rb") as key_file:
            private_key = load_pem_private_key(key_file.read(), password=None)

        ciphertext = base64.b64decode(encrypted_message)

        plaintext = private_key.decrypt(
            ciphertext,
            asymmetric_padding.OAEP(
                mgf=asymmetric_padding.MGF1(algorithm=hashes.SHA256()),
                algorithm=hashes.SHA256(),
                label=None
            )
        )

        return plaintext.decode("utf-8")

    def getSecret(self, secretID: str) -> tuple[str, bool, int]:
        """Retrieve a secret from the server using GET request."""
        secretRequest = {
            "secret_id": secretID
        }

        r = requests.get(url=f"{self._url}secret", json=secretRequest, headers=self._auth_headers, verify=False)
        if r.status_code == 200:
            re = json.loads(r.text)
            return self._RSA_Decrypt(re["text"]), True, 200
        else:
            return "", False, r.status_code

    def addSecret(self, secretID: str, secret: str) -> tuple[bool, int]:
        """Post a secret to the server using POST request."""
        secret_data = {
            "secret_id": secretID,
            "secret": self._RSA_Encrypt(secret)
        }

        r = requests.post(url=f"{self._url}secret", json=secret_data, headers=self._auth_headers, verify=False)
        return r.status_code == 200, r.status_code
pnd = Pandora("https://localhost:8080/","mustafa", "0770", "private_key.pem", "public_key.pem")
secret, ok, status = pnd.getSecret("tokens")
if not ok:
    sys.exit(status)
print(secret)
