package pkg

type City struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewCity(id int, name string) *City {
	return &City{
		ID:   id,
		Name: name,
	}
}
