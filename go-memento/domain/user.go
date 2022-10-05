package domain

type User struct {
	name   string
	age    int
	active bool
}

type Users []User

func NewUser(name string, age int, active bool) User {
	return User{
		name:   name,
		age:    age,
		active: active,
	}
}

type UserMemento struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

func (u User) Memento() UserMemento {
	return UserMemento{
		Name:   u.name,
		Age:    u.age,
		Active: u.active,
	}
}

func (m UserMemento) Restore() User {
	return User{
		name:   m.Name,
		age:    m.Age,
		active: m.Active,
	}
}

func (u Users) Memento() []UserMemento {
	t := make([]UserMemento, 0, len(u))
	for idx := range u {
		t = append(t, u[idx].Memento())
	}
	return t
}
