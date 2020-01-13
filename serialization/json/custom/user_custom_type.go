package custom

import "encoding/json"

type SensitiveString struct {
	string  // Underlying type
}

func (ss SensitiveString) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal("")
	return data, err
}

func (ss *SensitiveString) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	ss.string = s

	return nil
}

type UserCustomType struct {
	ID        string          `json:"id"`
	Email     string          `json:"email"`
	Password  SensitiveString `json:"password"`
	Note      string          `json:"note,omitempty"`
	Tags      string          `json:"-"`
}
