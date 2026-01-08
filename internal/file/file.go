package file

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/hadlow/genomdb/types"
)

func Get(path string) (types.Config, error) {
	// get file contents
	contents, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
		return types.Config{}, fmt.Errorf("error reading file")
	}

	// parse yaml into object
	request, err := ParseYaml(contents)

	if err != nil {
		log.Fatal(err)
		return types.Config{}, fmt.Errorf("error parsing YAML")
	}

	return request, nil
}

func ParseYaml(contents []byte) (types.Config, error) {
	var request types.Config

	err := yaml.Unmarshal(contents, &request)

	if err != nil {
		return request, err
	}

	return request, nil
}
