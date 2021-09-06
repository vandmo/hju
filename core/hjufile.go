package core

import (
	"encoding/json"
	"os"
	"sort"

	"github.com/vandmo/hju/git"
)

type hjuFileJson struct {
	Repositories []string `json:"repositories"`
}

type HjuFile struct {
	Repositories []string
	Folders      []string
}

func ParseHjuFileOrNew() (*HjuFile, error) {
	hjuFile, err := ParseHjuFile()
	if err != nil && os.IsNotExist(err) {
		return &HjuFile{Repositories: make([]string, 0)}, nil
	}
	return hjuFile, err
}

func ParseHjuFile() (*HjuFile, error) {
	jsonFile, fileOpenErr := os.Open("hju.json")
	if fileOpenErr != nil {
		return nil, fileOpenErr
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	var hju hjuFileJson
	decoderErr := decoder.Decode(&hju)
	if decoderErr != nil {
		return nil, decoderErr
	}
	var folders []string
	for _, repo := range hju.Repositories {
		folders = append(folders, git.FolderName(repo))
	}
	return &HjuFile{Repositories: hju.Repositories, Folders: folders}, nil
}

func (hf *HjuFile) ContainsFolder(folderName string) bool {
	for _, existing := range hf.Folders {
		if existing == folderName {
			return true
		}
	}
	return false
}

func (hf *HjuFile) Add(repo string) {
	hf.Repositories = append(hf.Repositories, repo)
}

func (hf *HjuFile) RemoveByFolder(folder string) {
	old := hf.Repositories
	hf.Repositories = hf.Repositories[:0]
	for _, repository := range old {
		if git.FolderName(repository) != folder {
			hf.Repositories = append(hf.Repositories, repository)
		}
	}
}

func (hf *HjuFile) Write() error {
	sort.Strings(hf.Repositories)
	jsonFile, fileErr := os.Create("hju.json")
	if fileErr != nil {
		return fileErr
	}
	defer jsonFile.Close()
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ")
	encoderErr := encoder.Encode(&hjuFileJson{Repositories: hf.Repositories})
	if encoderErr != nil {
		return encoderErr
	}
	return nil
}
