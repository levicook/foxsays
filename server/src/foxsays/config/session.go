package config

var Session session

type session struct {
	NewAuthenticationKey string `toml:"new_authentication_key"`
	OldAuthenticationKey string `toml:"old_authentication_key"`
	NewEncryptionKey     string `toml:"new_encryption_key"`
	OldEncryptionKey     string `toml:"old_encryption_key"`
}
