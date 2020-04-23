package secret

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	cipher "github.com/RuNpiXelruN/secrets-cli-app/cipher"
)

// VaultService type
type VaultService interface {
	readKeyValues(r io.Reader) error
	loadKeyVals() error
	saveKeyValues() error
	writeKeyValues(w io.Writer) error
	Get(key string) (string, error)
	Set(key, val string) error
	ListKeys() ([]string, error)
	RemoveKey(key string) error
}

// Vault type
type Vault struct {
	encodingKey string
	filepath    string
	keyValues   map[string]string
	mu          sync.Mutex
}

// NewFileVault func
func NewFileVault(encodingKey, filepath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filepath:    filepath,
	}
}

func (v *Vault) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}

func (v *Vault) loadKeyVals() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	defer f.Close()

	r, err := cipher.DecryptReader(v.encodingKey, f)
	if err != nil {
		return err
	}

	return v.readKeyValues(r)
}

func (v *Vault) saveKeyValues() error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	streamWriter, err := cipher.EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}

	return v.writeKeyValues(streamWriter)
}

func (v *Vault) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(&v.keyValues)
}

// RemoveKey func
func (v *Vault) RemoveKey(key string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	err := v.loadKeyVals()
	if err != nil {
		return err
	}

	_, ok := v.keyValues[key]
	if !ok {
		return fmt.Errorf("secrets: cannot find the key, %v", key)
	}

	delete(v.keyValues, key)
	err = v.saveKeyValues()
	if err != nil {
		return err
	}

	return nil
}

// ListKeys func
func (v *Vault) ListKeys() ([]string, error) {
	v.mu.Lock()
	defer v.mu.Unlock()

	var keys []string

	err := v.loadKeyVals()
	if err != nil {
		return nil, err
	}

	for k := range v.keyValues {
		keys = append(keys, k)
	}

	return keys, nil
}

// Get func
func (v *Vault) Get(key string) (string, error) {
	v.mu.Lock()
	defer v.mu.Unlock()

	err := v.loadKeyVals()
	if err != nil {
		return "", err
	}

	val, ok := v.keyValues[key]
	if !ok {
		return "", fmt.Errorf("secrets: cannot find the key, %v", key)
	}
	return val, nil
}

// Set func
func (v *Vault) Set(key, val string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	err := v.loadKeyVals()
	if err != nil {
		return err
	}

	v.keyValues[key] = val
	err = v.saveKeyValues()
	if err != nil {
		return err
	}
	return nil
}
