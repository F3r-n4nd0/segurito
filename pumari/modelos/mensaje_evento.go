package modelos

import (
	"encoding/json"
	"time"
)

type MensanjeEvento struct {
	UserName string     `json:"user_name"`
	Tipo     TipoEvento `json:"tipo"`
	Fecha    time.Time  `json:"fecha"`
}

func UnMarshal(bytes []byte) (*MensanjeEvento, error) {
	var evento MensanjeEvento
	err := json.Unmarshal(bytes, &evento)
	if err != nil {
		return nil, err
	}
	return &evento, nil
}
