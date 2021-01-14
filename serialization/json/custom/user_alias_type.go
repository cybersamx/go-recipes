package custom

import "encoding/json"

type UserAliasType struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Note     string `json:"note,omitempty"`
	Tags     string `json:"-"`
}

func (u UserAliasType) MarshalJSON() ([]byte, error) {
	type UserAlias User // Must create an alias otherwise

	alias := UserAlias{
		ID:    u.ID,
		Email: u.Email,
		Note:  u.Note,
		Tags:  u.Tags,
	}

	return json.Marshal(alias)
}
