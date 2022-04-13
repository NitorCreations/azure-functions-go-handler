package function

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
)

type Allocatable interface {
	Allocate() error
}

// Convenience type for binary input and outpu bindings.
// Handles "null" values correctly.
type Binary struct {
	Buffer *bytes.Buffer
}

func (b *Binary) Allocate() error {
	if b.Buffer == nil {
		b.Buffer = &bytes.Buffer{}
	}
	return nil
}

func (b Binary) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte(`"null"`)) {
		return nil
	}

	buffer := bytes.NewBuffer(data[1 : len(data)-1])
	decoder := base64.NewDecoder(base64.StdEncoding, buffer)
	_, err := b.Buffer.ReadFrom(decoder)
	return err
}

func (b Binary) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Buffer.Bytes())
}
