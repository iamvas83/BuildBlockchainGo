package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

type Wallet struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  []byte
}

func (w *Wallet) GenerateRSAKeys() (*rsa.PrivateKey, []byte, error) {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	publickey := &privatekey.PublicKey

	//Encode public key to PEM format

	pubASN1, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal public key: %v", err)
	}

	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})

	return privatekey, pubPEM, nil
}

func (w *Wallet) NewWallet() (*Wallet, error) {
	privateKey, publickey, err := w.GenerateRSAKeys()
	if err != nil {
		return nil, err
	}

	wallet := &Wallet{
		PrivateKey: privateKey,
		PublicKey:  publickey,
	}
	return wallet, nil
}

func (w *Wallet) SignTransaction(tx *Transaction) (string, error) {
	data := tx.Sender + tx.Receiver + fmt.Sprintf("%f", tx.Amount) + fmt.Sprintf("%t", tx.Coinbase)

	hashedData := sha256.Sum256([]byte(data))

	signature, err := rsa.SignPKCS1v15(rand.Reader, w.PrivateKey, crypto.SHA256, hashedData[:])

	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %v", err)
	}

	encodedSignature := base64.StdEncoding.EncodeToString(signature)

	return encodedSignature, nil

}

func (w *Wallet) VerifyTransaction(tx *Transaction, publickey []byte, signature string) error {
	data := tx.Sender + tx.Receiver + fmt.Sprintf("%f", tx.Amount) + fmt.Sprintf("%t", tx.Coinbase)

	hashedData := sha256.Sum256([]byte(data))

	signatureBytes, err := base64.StdEncoding.DecodeString(signature)

	if err != nil {
		return fmt.Errorf("failed to decode signature %v", err)
	}

	rsaPublicKey, err := ConvertPemToPublicKey(publickey)
	if err != nil {
		return fmt.Errorf("invalid public key: %v", err)
	}
	err = rsa.VerifyPKCS1v15(rsaPublicKey, crypto.SHA256, hashedData[:], signatureBytes)
	if err != nil {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func ConvertPemToPublicKey(pemPub []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemPub)

	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pubKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}
	return pubKey, nil
}
