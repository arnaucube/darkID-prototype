# darkID
Blockchain based anonymous distributed ID system

### Main concept
The objective is to guarantee a decentralized login system, but making sure that registered users are real ones and there are no bots generating large amounts of accounts. Only the verified (by email or phone) users can generate an anonymous ID (the Public-Key blind signed).


![screenshot](https://raw.githubusercontent.com/arnaucode/darkID/master/documentation/screenshot01.png "screenshot")

![screenshot](https://raw.githubusercontent.com/arnaucode/darkID/master/documentation/screenshot02.png "screenshot")

## How it works?


#### Network infrastructure

![network](https://raw.githubusercontent.com/arnaucode/darkID/master/documentation/darkID-network.png "network")


#### Step by step process
1. Once all the nodes of the network are running, a new user can connect to the server-ID-signer.
2. The user registers a non anonymous user (using email, phone, password, etc), and performs the login with that user
3. The user, locally, generates a RSA key pair (private key & public key)
4. The user blinds his Public-Key with the server-ID-signer Public-Key
5. The user's Public-Key blinded, is sent to the server-ID-signer
6. The server-ID-signer Blind Signs the Public-Key blinded from the user, and returns it to the user
7. The user unblinds the Public-Key signed by the server-ID-signer, and now has the Public-Key Blind Signed by the server-ID-signer
8. The user sends the Public-Key blind signed to the p2p network
9. The peers verify that the Public-Key Blind Signed is correctly signed by the server-ID-signer, if it is, they add the Public-Key to the Ethereum Blockchain, inside a new block
10. Then, when the user wants to login into a platform, just needs to put his Public-Key
11. The platform goes to the Ethereum Blockchain, to check if this Public-Key is registered in the blockchain
12. The platform sends a message encrypted with the user Public-Key, and the user returns the message decrypted with the Private-Key, to verify that is the owner of that Public-Key


##### RSA encryption system
https://en.wikipedia.org/wiki/RSA_cryptosystem
- Public parameters: (e, n)
- Private parameters: (d, p, q, phi, sigma)
- Public-Key = (e, n)
- Private-Key = (d, n)
- Encryption:
![rsa](https://wikimedia.org/api/rest_v1/media/math/render/svg/fbfc70524a1ad983e6f3aac51226b9ca92fefb10 "rsa")
- Decryption:
![rsa](https://wikimedia.org/api/rest_v1/media/math/render/svg/10227461ee5f4784484f082d744ba5b8c468668c "rsa")


##### Blind signature process
https://en.wikipedia.org/wiki/Blind_signature
- m is the message (in our case, is the Public-Key of the user to be blinded)
![rsa](https://wikimedia.org/api/rest_v1/media/math/render/svg/a59b57fa153c8b327605672caadb0ecf59e5795a "rsa")

- server-ID-signer blind signs m'
![rsa](https://wikimedia.org/api/rest_v1/media/math/render/svg/e726b003ff1649f9254032cffae42d80577da787 "rsa")

- user can unblind m, to get m signed
![rsa](https://wikimedia.org/api/rest_v1/media/math/render/svg/e96fad0e1d46ec4c55986d1c8fc84e8c44259ecc "rsa")

- This works because RSA keys satisfy this equation
![rsa](https://wikimedia.org/api/rest_v1/media/math/render/svg/d6bd21fb4e25c311df07b50c313a248d978c3212 "rsa") and this ![rsa](https://wikimedia.org/api/rest_v1/media/math/render/svg/c13170a26e031125b417f22644fb64384c04eea7 "rsa")
