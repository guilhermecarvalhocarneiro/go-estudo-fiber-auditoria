package pkg

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSONMap map[string]interface{}

func (j *JSONMap) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Falha ao escanear valor JSON: tipo inválido")
	}

	var data map[string]interface{}
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return fmt.Errorf("Falha ao escanear valor JSON: %v", err)
	}

	*j = JSONMap(data)
	return nil
}

func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}
