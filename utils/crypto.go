package utils

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func GenerateRandomId() string {
	ts := strconv.FormatInt(time.Now().UnixNano(), 10)
	mLoc := base64.StdEncoding.EncodeToString([]byte(ts))
	return mLoc
}

func FirstTimeSettup() {
	//Create a folder/directory at a full qualified path
	err := os.Mkdir("logfile", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Using Create() function
	myfile, e := os.Create("logfile/shiSocklogs.txt")
	if e != nil {
		log.Fatal(e)
	}
	if _, er := myfile.WriteString("Created Logfile folder and logfile/shiSocklogs.txt\n"); er != nil {
		log.Fatal(er)
	}
	myfile.Close()
}

func AddLog(str string) {
	f, err := os.OpenFile("logfile/shiSocklogs.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		FirstTimeSettup()
	} else {
		defer f.Close()
		if _, err = f.WriteString(str + "\n"); err != nil {
			panic(err)
		}
	}
}

func Encode(data []byte) string {
	hb := base64.StdEncoding.EncodeToString([]byte(data))
	return hb
}

// Decoding the base string to array of bytes
func Decode(data string) []byte {
	hb, _ := base64.StdEncoding.DecodeString(data)
	return hb
}

// Generating RSA private key
func GenerateRsaPrivateKey(size int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// Generating RSA public key
func GenerateRsaPublicKey(privateKey *rsa.PrivateKey) rsa.PublicKey {
	return privateKey.PublicKey
}

// This function can be use encrypt a plain text with rsa algorithm
func RsaEncrypt(publicKey rsa.PublicKey, data []byte) ([]byte, error) {
	// encryptedBytes, err := rsa.EncryptOAEP(
	// 	sha256.New(),
	// 	rand.Reader,
	// 	&publicKey,
	// 	data,
	// 	nil)
	encryptedBytes, err := rsa.EncryptPKCS1v15(
		rand.Reader,
		&publicKey,
		data)
	return encryptedBytes, err
}

// This function can be use decrypt a encrypted text with rsa algorithm
func RsaDecrypt(privateKey rsa.PrivateKey, data []byte) ([]byte, error) {
	decryptedBytes, err := rsa.DecryptPKCS1v15(rand.Reader, &privateKey, data)
	// decryptedBytes, err := privateKey.Decrypt(
	// 	nil,
	// 	data,
	// 	&rsa.OAEPOptions{Hash: crypto.SHA256})
	return decryptedBytes, err
}

//  This fucntion is used to dump/serialize the rsa public key
func DumpKey(key *rsa.PublicKey) ([]byte, error) {
	return x509.MarshalPKCS1PublicKey(key), nil
}

func ParseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("key type is not RSA")
}

// This function is used to load the rsa public key
func loadKey(byteKey []byte) (*rsa.PublicKey, error) {
	key, err := x509.ParsePKCS1PublicKey(byteKey)
	return key, err
}

func LoadKey(byteKey []byte) (*rsa.PublicKey, error) {
	key1, err1 := loadKey(byteKey)
	if err1 != nil {
		key2, err2 := ParseRsaPublicKeyFromPemStr(string(byteKey))
		if err2 != nil {
			return nil, err2
		} else {
			return key2, nil
		}
	} else {
		return key1, nil
	}
}

func VarifySignature(publicKey *rsa.PublicKey, digest []byte, signature []byte) bool {
	err := rsa.VerifyPSS(publicKey, crypto.SHA256, digest, signature, nil)
	if err != nil {
		err2 := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, digest, signature)
		if err2 != nil {
			return false
		} else {
			return true
		}
	} else {
		return true
	}
}

// Generate fixed size byte array
func GenerateAesKey(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

// This fucntion can be used for encrypting a plain text using AES-GCM algorithm
func AesEncryption(key []byte, data []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText, nil
}

// This fucntion can be used for decrypting the ciphertext encrypted using AES-GCM algorithm
func AesDecryption(key []byte, cipherText []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("1")
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)

	if err != nil {
		fmt.Println("2")
		return nil, err
	}

	noncesize := gcm.NonceSize()
	if len(cipherText) < noncesize {
		fmt.Println("3")
		return nil, err
	}

	nonce, cipherText := cipherText[:noncesize], cipherText[noncesize:]

	plainText, err := gcm.Open(nil, nonce, cipherText, nil)

	if err != nil {
		fmt.Println("4", err.Error())
		return nil, err
	}

	return plainText, nil
}
