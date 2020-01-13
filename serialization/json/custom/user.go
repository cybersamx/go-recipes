package custom

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Note      string    `json:"note,omitempty"`
	Tags      string    `json:"-"`
}
