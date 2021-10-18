package events

import (
	"encoding/json"
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestUserEvents(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("User Events", func() {
		g.It("should unmarshal a user event action", func() {
			var ue UserEvent
			err := json.Unmarshal([]byte(SampleUserEvent), &ue)

			Expect(err).To(BeNil())
			Expect(ue.Action).To(Equal("forgot_password"))
		})
	})
}

const (
	SampleUserEvent = `{"data": {"user":{"name":"foo"}}, "action":"forgot_password"}`
)
