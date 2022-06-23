package models

type User struct {
	ID    uint64
	Email string
	Senha string
}

func (u User) GetID() uint64 {
	return u.ID
}
