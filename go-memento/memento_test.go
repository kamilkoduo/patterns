package go_memento

import (
	"testing"

	"github.com/patterns/go-memento/repo"
	"github.com/stretchr/testify/assert"

	"github.com/patterns/go-memento/domain"
)

func TestMemento(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		r := repo.New()
		user := domain.NewUser("Alice", 20, true)
		r.Save(user)
		assert.Equal(t, user, r.Read())

		user.Memento()
	})
}
