package kandinsky

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

var (
	ErrEmptyURL     = errors.New("kandinsky url is not exists in env or in config file")
	ErrEmptyKey     = errors.New("kandinsky auth key is not exists in env or config file")
	ErrEmptySecret  = errors.New("kandinsky auth secret is not exists in env or config file")
	ErrAuth         = errors.New("kandinsky authentication error")
	ErrStatusNot200 = errors.New("kandinsky status is not 200")
)

// Kandinsky struct, all fields are required
// https://fusionbrain.ai/docs/ru/doc/api-dokumentaciya/
type Kandinsky struct {
	key    string
	secret string
	model  Model
}

// Model is the message from kandinsky API after auth
// [
//
//	{
//	    "id": 4,
//	    "name": "Kandinsky",
//	      "version": 3.0,
//	      "type": "TEXT2IMAGE"
//	}
//
// ]
type Model struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Version float32 `json:"version"`
	Type    string  `json:"type"`
}

// Params for generate image
//	{
//		"type": "GENERATE",
//		"style": "string",
//		"width": 1024,
//		"height": 1024,
//		"num_images": 1,
//		"negativePromptUnclip": "яркие цвета, кислотность, высокая контрастность",
//		"generateParams": {
//			"query": "Пушистый кот в очках",
//		}
//	}
type Params struct {
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	NumImages      int    `json:"num_images"`
	Type           string `json:"type"`
	Style          string `json:"style"`
	NegativePrompt string `json:"negativePromptUnclip"`
	GenerateParams struct {
		Query string `json:"query"`
	} `json:"generateParams"`
}

// UUID response with UUID from Kandinsky API
//	{
//		"uuid": "string",
//		"status": "INITIAL"
//	}
type UUID struct {
	ID     string `json:"uuid"`
	Status string `json:"status"`
}

// ErrResponse from Kandinsky API
//
//	{
//		"timestamp": "2024-03-04T13:46:55.473+00:00",
//		"status": 400,
//		"error": "Bad Request",
//		"message": "Failed to convert value of type 'java.lang.String' to required type 'int'; For input string: \"\"4\"\"",
//		"path": "/key/api/v1/text2image/run"
//	}
type ErrResponse struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

// New return Kandinsky instance
func New(key, secret string) (*Kandinsky, error) {
	if key == "" {
		return nil, ErrEmptyKey
	}

	if secret == "" {
		return nil, ErrEmptySecret
	}

	k := &Kandinsky{
		key:    key,
		secret: secret,
		model:  Model{},
	}

	return k, nil
}

// SetModel send auth request to url and set image UUID to Kandinsky instance from json response:
// [
//
//	{
//	    "id": 4,
//	    "name": "Kandinsky",
//	      "version": 3.0,
//	      "type": "TEXT2IMAGE"
//	}
//
// ]
func (k *Kandinsky) SetModel(url string) error {
	// create GET request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// create auth header
	req.Header.Add("X-Key", "Key " + k.key)
	req.Header.Add("X-Secret", "Secret " + k.secret)

	// create client
	client := http.Client{}

	// Do request to Kandinsky API
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// check status code from received from API
	if res.StatusCode != http.StatusOK {
		return ErrStatusNot200
	}

	dec := json.NewDecoder(res.Body)

	// read open bracket
	_, err = dec.Token()
	if err != nil {
		return err
	}

	var m Model

	// while the array contains values
	for dec.More() {
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			return err
		}
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		return err
	}

	if m.Id == 0 {
		return ErrAuth
	}

	k.model = m

	return nil
}

// GetUUID send POST request with params to url and return response:
//
//	{
//		"uuid": "string",
//		"status": "INITIAL"
//	}
func (k *Kandinsky) GetUUID(url string, params Params) (UUID, error) {
	u := UUID{}

	if k.model.Id == 0 {
		k.model.Id = 4
	}

	// marshall params to json bytes
	b, err := json.Marshal(&params)
	if err != nil {
		return u, err
	}

	// generate command string
	curlCommand := fmt.Sprintf(`curl --location --request POST 'https://api-key.fusionbrain.ai/key/api/v1/text2image/run' --header 'X-Key: Key %s' --header 'X-Secret: Secret %s' -F 'params=%s
	};type=application/json' --form 'model_id="%d"'`, k.key, k.secret, string(b), k.model.Id)

	// create command
	cmd := exec.Command("sh", "-c", curlCommand)

	// buffers for standard out and error
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// run command
	err = cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return u, err
	}

	// out to string
	s := out.String()
	// if response status not 200
	if strings.Contains(s, "error") {
		e := ErrResponse{}
		err = json.Unmarshal(out.Bytes(), &e)
		if err != nil {
			return u, err
		}

		return u, errors.New("error from Kandinsky API: status " + strconv.Itoa(e.Status) + " " + e.Error + " > " + e.Message)
	}

	// unmarshal out data to UUID struct
	err = json.Unmarshal(out.Bytes(), &u)
	if err != nil {
		return u, err
	}

	return u, nil
}
