package parse

import (
	"fmt"
	"mason/internal/models"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

func Yaml(file []byte) models.Mason {
	config := models.Mason{}

	err := yaml.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return format(config)
}

func format(data models.Mason) models.Mason {
	for idx, val := range data.Workflows {
		// regular express to replace spaces and underscores with dashes
		re := regexp.MustCompile(`[\s_\/]+`)
		val.Name = re.ReplaceAllString(strings.ToLower(val.Name), "-")

		data.Workflows[idx].Name = val.Name
	}

	return data
}