package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
)

type Kms struct {
	Client *kms.Client
}

func NewKms(c context.Context) (*Kms, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, er.W(err)
	}

	cli := kms.NewFromConfig(cfg)

	return &Kms{
		Client: cli,
	}, nil
}

func (k *Kms) EncryptText(keyID string, plaintext string) ([]byte, error) {
	output, err := k.Client.Encrypt(context.TODO(), &kms.EncryptInput{
		KeyId:     &keyID,
		Plaintext: []byte(plaintext),
	})
	if err != nil {
		return nil, er.W(err)
	}
	return output.CiphertextBlob, nil
}

func (k *Kms) DecryptText(ciphertext []byte) (string, error) {
	output, err := k.Client.Decrypt(context.TODO(), &kms.DecryptInput{
		CiphertextBlob: ciphertext,
	})
	if err != nil {
		return "", er.W(err)
	}
	return string(output.Plaintext), nil
}
