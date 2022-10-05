package repo

import (
	"encoding/json"

	"github.com/patterns/go-memento/domain"
)

func New() *Repo {
	return &Repo{
		data: nil,
	}
}

type Repo struct {
	data []byte
}

func (r *Repo) Save(u domain.User) {
	memento := u.Memento()
	bytes, err := json.Marshal(memento)
	if err != nil {
		panic(err)
	}
	r.data = bytes
}

func (r *Repo) Read() domain.User {
	var memento domain.UserMemento
	err := json.Unmarshal(r.data, &memento)
	if err != nil {
		panic(err)
	}
	return memento.Restore()
}
