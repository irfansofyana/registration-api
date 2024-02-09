package config

import (
	"github.com/siuyin/dflt"
	"os"
)

type Config struct {
	PrivateKey []byte
	PublicKey  []byte
}

var Instance Config

func init() {
	privateKeyPath := dflt.EnvString("PRIVATE_KEY_PATH", "./private.pem")
	privateKey, err := os.ReadFile(privateKeyPath)
	if err != nil {
		panic(err)
	}

	publicKeyPath := dflt.EnvString("PUBLIC_KEY_PATH", "./public.pem")
	publicKey, err := os.ReadFile(publicKeyPath)
	if err != nil {
		panic(err)
	}

	Instance = Config{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}
