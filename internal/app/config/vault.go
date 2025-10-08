package config

import (
	"context"
	"log"
	"os"

	vault "github.com/hashicorp/vault/api"
)

func LoadVaultSecrets() map[string]string {
	vaultAddr := os.Getenv("VAULT_ADDR")       // e.g. "http://127.0.0.1:8200"
	vaultToken := os.Getenv("VAULT_TOKEN")     // your root or app token
	vaultSecretPath := os.Getenv("VAULT_PATH") // e.g. "secret/data/bumimedika"

	if vaultAddr == "" || vaultToken == "" || vaultSecretPath == "" {
		log.Println("[WARN] Vault not configured, fallback to .env")
		return nil
	}

	client, err := vault.NewClient(&vault.Config{Address: vaultAddr})
	if err != nil {
		log.Fatalf("Failed to create Vault client: %v", err)
	}
	client.SetToken(vaultToken)

	secret, err := client.KVv2("secret").Get(context.Background(), "bumimedika")
	if err != nil {
		log.Fatalf("Failed to read Vault secret: %v", err)
	}

	result := make(map[string]string)
	for k, v := range secret.Data {
		result[k] = v.(string)
	}
	return result
}
