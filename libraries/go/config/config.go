package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

var DefaultProvider Provider

func init() {
	file := os.Getenv("CONFIG_FILE")
	c, err := LoadConfigFile(file)
	if err != nil {
		log.Println("Failed to load configuration file")
	}
	DefaultProvider = New(c)
}

type Provider interface {
	Has(string) bool
	Get(string) Value
}

type Config struct {
	Map map[string]interface{}
}

type Value struct {
	raw interface{}
}

func mustGetDefaultProvider() Provider {
	if DefaultProvider == nil {
		log.Println("Default provider not defined")
	}
	return DefaultProvider
}

func Has(path string) bool  { return mustGetDefaultProvider().Has(path) }
func Get(path string) Value { return mustGetDefaultProvider().Get(path) }

func LoadConfigFile(file string) (map[interface{}]interface{}, error) {
	var c map[interface{}]interface{}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	return c, nil
}

func New(content map[interface{}]interface{}) Provider {
	m, err := convertKeysToString(content)
	if err != nil {
		return &Config{}
	}
	return &Config{
		Map: m,
	}
}

func convertKeysToString(m map[interface{}]interface{}) (map[string]interface{}, error) {
	n := make(map[string]interface{})

	for k, v := range m {
		// Assert that the key is a string
		str, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("config key is not a string")
		}

		if vMap, ok := v.(map[interface{}]interface{}); ok {
			var err error
			v, err = convertKeysToString(vMap)
			if err != nil {
				return nil, err
			}
		}

		n[str] = v
	}

	return n, nil
}

func (c *Config) Has(path string) bool {
	v := c.Get(path)
	return v.raw != nil
}

func (c *Config) Get(path string) Value {
	return Value{
		raw: reduce(strings.Split(path, "."), c.Map),
	}
}

func reduce(parts []string, value interface{}) interface{} {
	if len(parts) == 0 {
		return value
	}
	valueMap, ok := value.(map[string]interface{})
	if !ok {
		return nil
	}
	value, ok = valueMap[parts[0]]
	if !ok {
		return nil
	}
	return reduce(parts[1:], value)
}

func (v Value) Int(defaults ...int) int {
	defaults = append(defaults, 0)
	if v.raw == nil {
		return defaults[0]
	}
	r, ok := v.raw.(int)
	if !ok {
		return defaults[0]
	}
	return r
}

func (v Value) String(defaults ...string) string {
	defaults = append(defaults, "")
	if v.raw == nil {
		return defaults[0]
	}
	return fmt.Sprintf("%s", v.raw)
}

func (v Value) Bool(defaults ...bool) bool {
	defaults = append(defaults, false)
	if v.raw == nil {
		return defaults[0]
	}
	r, ok := v.raw.(bool)
	if !ok {
		return defaults[0]
	}
	return r
}
