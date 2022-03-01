package organizations

import (
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"testing"
)

func TestOrganization(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
}
