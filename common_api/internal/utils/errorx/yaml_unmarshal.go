package errorx

import "gopkg.in/yaml.v3"

func UnmarshalYAML(data []byte, v interface{}) error {
    return yaml.Unmarshal(data, v)
}
