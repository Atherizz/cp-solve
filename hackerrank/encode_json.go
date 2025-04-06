package main

import (
	"encoding/json"
	"bytes"
	"io"
)

type Manager struct {
    FullName       string `json:"full_name,omitempty"`
    Position       string `json:"position,omitempty"`
    Age            int32  `json:"age,omitempty"`
    YearsInCompany int32  `json:"years_in_company,omitempty"`

}

func EncodeManager(manager *Manager) (io.Reader, error) {
	jsonData, err := json.Marshal(manager)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonData), nil
}
