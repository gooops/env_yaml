package env_yaml

import (
	"log"

	"gopkg.in/yaml.v2"

	"github.com/gooops/env_strings"
)

const (
	ENV_YAML_KEY = "ENV_YAML"
	ENV_YAML_EXT = ".env"
)

type EnvYaml struct {
	*env_strings.EnvStrings
}

func NewEnvYaml(envName string, envExt, configType string) *EnvYaml {
	if envName == "" {
		panic("env_yaml: env name could not be nil")
	}

	return &EnvYaml{
		EnvStrings: env_strings.NewEnvStrings(envName, envExt, configType),
	}
}

func (p *EnvYaml) Marshal(v interface{}) (data []byte, err error) {
	return yaml.Marshal(v)
}

// func (p *EnvYaml) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
// 	return yaml.MarshalIndent(v, prefix, indent)
// }

func (p *EnvYaml) Unmarshal(data []byte, v interface{}) (err error) {
	strData := ""
	if strData, err = p.Execute(string(data)); err != nil {
		return
	}
	log.Println(strData)
	err = yaml.Unmarshal([]byte(strData), v)

	return
}

func Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

// func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
// 	return yaml.MarshalIndent(v, prefix, indent)
// }

func Unmarshal(data []byte, v interface{}) error {
	envYaml := NewEnvYaml(ENV_YAML_KEY, ENV_YAML_EXT, "yaml")
	return envYaml.Unmarshal(data, v)
}
