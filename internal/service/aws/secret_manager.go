package aws

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/sso/types"
	"github.com/skyrocketOoO/serverx/internal/domain/er"
)

type SecretManager struct {
	Client *secretsmanager.Client
}

func NewSecretManager(c context.Context) (*SecretManager, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, er.W(err)
	}

	cli := secretsmanager.NewFromConfig(cfg)

	return &SecretManager{
		Client: cli,
	}, nil
}

// GetSecret fetches the secret string from AWS Secrets Manager
func (sm *SecretManager) GetSecret(ctx context.Context, secretID string) (string, error) {
	out, err := sm.Client.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: &secretID,
	})
	if err != nil {
		return "", fmt.Errorf("failed to get secret: %w", err)
	}
	if out.SecretString == nil {
		return "", fmt.Errorf("secret string is nil")
	}
	return *out.SecretString, nil
}

// SetSecret creates or updates a secret in AWS Secrets Manager
func (sm *SecretManager) SetSecret(ctx context.Context, secretID, value string) error {
	// Try to update first
	_, err := sm.Client.UpdateSecret(ctx, &secretsmanager.UpdateSecretInput{
		SecretId:     &secretID,
		SecretString: &value,
	})
	// If not found, create new
	if err != nil {
		var notFound *types.ResourceNotFoundException
		if ok := errors.As(err, &notFound); ok {
			_, err = sm.Client.CreateSecret(ctx, &secretsmanager.CreateSecretInput{
				Name:         &secretID,
				SecretString: &value,
			})
		}
	}

	if err != nil {
		return fmt.Errorf("failed to set secret: %w", err)
	}

	return nil
}
