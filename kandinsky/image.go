package kandinsky

import (
	"encoding/base64"
	"os"
)

type Image struct {
	UUid     string   `json:"uuid"`
	Status   string   `json:"status"`
	Images   []string `json:"images"`
	Censored bool     `json:"censored"`
}

// ToByte is convert Image to []byte
func (i *Image) ToByte() ([]byte, error) {
	l := len(i.Images[0])
	var b = make([]byte, l)

	_, err := base64.StdEncoding.Decode(b, []byte(i.Images[0]))
	if err != nil {
		return nil, err
	}

	return b, nil
}

// ToFile is convert Image to os.File
func (i *Image) ToFile() (*os.File, error) {
	f, err := os.OpenFile(".temp.png", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return f, err
	}

	data, err := base64.StdEncoding.DecodeString(i.Images[0])
	if err != nil {
		return f, err
	}

	_, err = f.Write(data)
	if err != nil {
		return f, err
	}

	return f, nil
}

