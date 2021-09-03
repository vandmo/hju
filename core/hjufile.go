package core

import (
	"encoding/json"
	"os"
	"sort"
)

type HjuFile struct {
	Repositories []string `json:"repositories"`
}

func ParseHjuFile() ([]string, error) {
	jsonFile, fileOpenErr := os.Open("hju.json")
	if fileOpenErr != nil {
		return nil, fileOpenErr
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	var hju HjuFile
	decoderErr := decoder.Decode(&hju)
	if decoderErr != nil {
		return nil, decoderErr
	}
	return hju.Repositories, nil
}

func WriteHjuFile(repositories []string) error {
	sort.Strings(repositories)
	jsonFile, fileErr := os.Create("hju.json")
	if fileErr != nil {
		return fileErr
	}
	defer jsonFile.Close()
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ")
	encoderErr := encoder.Encode(&HjuFile{Repositories: repositories})
	if encoderErr != nil {
		return encoderErr
	}
	return nil
}
