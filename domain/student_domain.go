package domain

type Student struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Class int     `json:"class"`
	Mark  float32 `json:"mark"`
	Email string  `json:"email"`
}
