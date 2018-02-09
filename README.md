# darkID: A proof of concept of an anonymous decentralized identification system based on blockchain
Blockchain based anonymous decentralized ID system

( Full slides in: https://github.com/arnaucode/darkID/blob/master/darkID-presentation.pdf )

## 1.- Main concept
The main idea behind darkID is to implement a proof of concept of a decentralized system that allows platforms to identify verified users, but ensuring their anonymity.
The main point, is to ensure anonymity, but at the same time, allow users to verify their identity, to ensure that no fake accounts are being used.


## 2.- How it works

- Verify the non anonymous ID of an user
  - Based on the: username, email, phone, ID card, etc
- From a logged and verified user, generate an anonymous darkID (Public Key, and save the Private Key into the storage), and get that darkID signed by an authority (serverIDsigner), with high reputation, and without knowing what is signing, to ensure the anonymity of the user
- Once the darkID (Public Key) is signed by an authority server, add the darkID to the blockchain (ethereum, or some other)
- Use the darkID to authenticate in platforms (just need to point the darkID reference in the blockchain), and the platform will send a challenge to the user, to resolve it with the Private Key assigned to the darkID from the darkID desktop app.

```go
type DarkID struct {
	ID             string         `json:"id"`
	PrivK          string         `json:"privK"` //path of the PrivK file
	PubK           string         `json:"pubK"`  //path of the PubK file
	PublicKey      string         `json:"publicKey"`
	Date           time.Time      `json:"date"`
	Hashed         []byte         `json:"hashed"`
	UnblindedSig   []byte         `json:"unblindedsig"`
	Verified       bool           `json:"verified"`
	ServerVerifier *rsa.PublicKey `json:"serververifier"`
	SignerID       string         `json:"signerid"`
	BlockchainRef  string         `json:"blockchainref"`
}
```

### 2.1.- Infraestructure

![network](https://raw.githubusercontent.com/arnaucode/darkID/master/documentation/darkID-network.png "network")


#### 2.2.1.- serverIDsigner
- The server where the user creates a non anonymous account
- Also is the server that blind signs the anonymous ID of the users
- This server must be a recognized authority, or based on some reputation system. As only the reliable serverIDsigner will be used by the users to trust their ID validation
- The serverIDsigner, can be different servers, based on a public reputation, where the users trust the servers that their are using to verify their darkIDs.

#### 2.2.2.- ethereum Smart Contract
Where the darkID is stored. In this case, is a very simple smart contract.

#### 2.2.3.- Desktop app
An electron desktop app implemented in Angularjs and Go lang.

![screenshot](https://raw.githubusercontent.com/arnaucode/darkID/master/documentation/darkID-screenshot01.png "screenshot")

![screenshot](https://raw.githubusercontent.com/arnaucode/darkID/master/documentation/darkID-screenshot02.png "screenshot")

![screenshot](https://raw.githubusercontent.com/arnaucode/darkID/master/documentation/darkID-screenshot03.png "screenshot")


### 2.2.- Step by step process

- Once all the nodes of the network are running, a new user can connect to the serverIDsigner.
- The user registers a non anonymous user (using email, phone, password, etc), and performs the login with that user
- The user, locally, generates a RSA key pair (private key & public key)
- The user blinds his Public-Key with the serverIDsigner Public-Key
- The user's Public-Key blinded, is sent to the serverIDsigner
- The serverIDsigner Blind Signs the Public-Key blinded from the user, and returns it to the user
- The user unblinds the Public-Key signed by the serverIDsigner, and now has the Public-Key Blind Signed by the serverIDsigner
- The user sends the Public-Key blind signed to the p2p network
- The peers verify that the Public-Key Blind Signed is correctly signed by the serverIDsigner, if it is, they add the Public-Key to the Ethereum Blockchain, inside a new block

![creation](https://raw.githubusercontent.com/arnaucode/darkID/master/documentation/darkID-creation.png "creation")

- Then, when the user wants to login into a platform, just needs to put his Public-Key
- The platform goes to the Ethereum Blockchain, to check if this Public-Key is registered in the blockchain
- The platform sends a message encrypted with the user Public-Key, and the user returns the message decrypted with the Private-Key, to verify that is the owner of that Public-Key

![login](https://raw.githubusercontent.com/arnaucode/darkID/master/documentation/darkID-login.png "login")

The user Private Keys are stored only in the user local directory 'darkID/keys'.

## 3.- Cryptography
### 3.1.- RSA encryption system
https://en.wikipedia.org/wiki/RSA_cryptosystem
- Public parameters: (e, n)
- Private parameters: (d, p, q, phi, sigma)
- Public-Key = (e, n)
- Private-Key = (d, n)
- Encryption:
![rsa](https://wikimedia.org/api/rest_v1/media/math/render/svg/fbfc70524a1ad983e6f3aac51226b9ca92fefb10 "rsa")
- Decryption:
![rsa](https://wikimedia.org/api/rest_v1/media/math/render/svg/10227461ee5f4784484f082d744ba5b8c468668c "rsa")


### 3.2.- Blind signature process
https://en.wikipedia.org/wiki/Blind_signature
- m is the message (in our case, is the Public-Key of the user to be blinded)
![rsa](https://wikimedia.org/api/rest_v1/media/math/render/svg/a59b57fa153c8b327605672caadb0ecf59e5795a "rsa")

- serverIDsigner blind signs m'
![rsa](https://wikimedia.org/api/rest_v1/media/math/render/svg/e726b003ff1649f9254032cffae42d80577da787 "rsa")

- user can unblind m, to get m signed
![rsa](https://wikimedia.org/api/rest_v1/media/math/render/svg/e96fad0e1d46ec4c55986d1c8fc84e8c44259ecc "rsa")

- This works because RSA keys satisfy this equation
![rsa](https://wikimedia.org/api/rest_v1/media/math/render/svg/d6bd21fb4e25c311df07b50c313a248d978c3212 "rsa") and this ![rsa](https://wikimedia.org/api/rest_v1/media/math/render/svg/c13170a26e031125b417f22644fb64384c04eea7 "rsa")


## 4.- Conclusions
- This is just a proof of concept, as an extra [project in a university small subject](https://github.com/arnaucode/darkID/blob/master/darkID-presentation.pdf). Is not a full finalized project.
- Cryptographic [blind signature](https://en.wikipedia.org/wiki/Blind_signature) is very powerful, and also, can be combined with [homomorphic encryption](https://en.wikipedia.org/wiki/Homomorphic_encryption) properties, to make secure and anonymous systems.
- In this proof of concept, the smart contract integration is not finished, and is using ethereum Smart Contracts, but can be implemented using any other blockchain technology. At the beginning [I tried to implement a complete p2p network and a blockchain from scratch](https://github.com/arnaucode/blockchainIDsystem) (but was not a good idea having a short amount of time).
- A decentralized anonymized login system over blockchain, can have lots of applications for example, can be useful to authentication for centralized and decentralized platforms, for voting systems, for health systems, for exchanges, for anonymous reputation systems, etc.
