package keystore

// Provider keystore serializer provider
type Provider interface {
	Read(data []byte, password string) (*Key, error)
	Write(key *Key, password string, attrs map[string]interface{}) ([]byte, error)
	KdfTypeName() []string
}

// Key keystore handled key object
type Key struct {
	ID         []byte
	Address    string
	PrivateKey []byte
}

// KdfParams .
type KdfParams struct {
	DkLen int    `json:"dklen"` // DK length
	Salt  string `json:"salt"`  // salt string
}

type encryptedKeyJSONV3 struct {
	Address string     `json:"address"`
	Crypto  cryptoJSON `json:"crypto"`
	ID      string     `json:"id"`
	Version int        `json:"version"`
}

type cryptoJSON struct {
	Cipher       string                 `json:"cipher"`
	CipherText   string                 `json:"ciphertext"`
	CipherParams cipherparamsJSON       `json:"cipherparams"`
	KDF          string                 `json:"kdf"`
	KDFParams    map[string]interface{} `json:"kdfparams"`
	MAC          string                 `json:"mac"`
}

type cipherparamsJSON struct {
	IV string `json:"iv"`
}

var providers = []Provider{
	&Web3KeyStore{},
}

// Decrypt read key from keystore
func Decrypt(data []byte, password string) (*Key, error) {
	provider := &Web3KeyStore{}

	return provider.Read(data, password)
}

// Encrypt encrypt key as keystore data
func Encrypt(key *Key, password string, attrs map[string]interface{}) ([]byte, error) {
	provider := &Web3KeyStore{}

	return provider.Write(key, password, attrs)
}

func selectProvider(keystoreType string) (Provider, bool) {
	for _, provider := range providers {
		for _, support := range provider.KdfTypeName() {
			if support == keystoreType {
				return provider, true
			}
		}
	}

	return nil, false
}
