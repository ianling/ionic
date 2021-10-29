package ionic

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/franela/goblin"
	"github.com/gomicro/bogus"
	"github.com/ion-channel/ionic/teamusers"
	. "github.com/onsi/gomega"
)

func TestTeamUsers(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Teams", func() {
		var server *bogus.Bogus
		var client *IonClient

		g.BeforeEach(func() {
			server = bogus.New()
			h, p := server.HostPort()
			client, _ = New(fmt.Sprintf("http://%v:%v", h, p))
		})

		g.It("should create a team user", func() {
			server.AddPath("/v1/teamUsers/createTeamUser").
				SetMethods("POST").
				SetPayload([]byte(SampleCreateTeamUser)).
				SetStatus(http.StatusOK)

			opts := CreateTeamUserOptions{
				Role:   "admin",
				Status: "active",
				TeamID: "team",
				UserID: "user",
			}
			tu, err := client.CreateTeamUser(opts, "atoken")
			Expect(err).To(BeNil())
			Expect(tu.ID).To(Equal("teamuser"))
			Expect(tu.Role).To(Equal("admin"))
			Expect(tu.Status).To(Equal("active"))
			Expect(tu.TeamID).To(Equal("team"))
			Expect(tu.UserID).To(Equal("user"))
		})

		g.It("should update a team user", func() {
			server.AddPath("/v1/teamUsers/updateTeamUser").
				SetMethods("PUT").
				SetPayload([]byte(SampleUpdateTeamUser)).
				SetStatus(http.StatusOK)

			tu := &teamusers.TeamUser{
				ID:   "someid",
				Role: "member",
			}

			tu, err := client.UpdateTeamUser(tu, "atoken")
			Expect(err).To(BeNil())
			Expect(tu.ID).To(Equal("someid"))
			Expect(tu.Role).To(Equal("member"))

			hr := server.HitRecords()
			Expect(len(hr)).To(Equal(1))
			Expect(hr[0].Verb).To(Equal("PUT"))
			Expect(string(hr[0].Body)).To(Equal(`{"id":"someid","team_id":"","user_id":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","deleted_at":"0001-01-01T00:00:00Z","status":"","role":"member"}`))

		})

		g.It("should delete a team user", func() {
			server.AddPath("/v1/teamUsers/deleteTeamUser").
				SetMethods("DELETE").
				SetStatus(http.StatusNoContent)

			tu := &teamusers.TeamUser{
				ID: "someid",
			}

			err := client.DeleteTeamUser(tu, "atoken")
			Expect(err).To(BeNil())
			Expect(tu.ID).To(Equal("someid"))

			hr := server.HitRecords()
			Expect(len(hr)).To(Equal(1))
			Expect(hr[0].Verb).To(Equal("DELETE"))
		})
	})
}

const (
	SampleCreateTeamUser = `{"data":{"id":"teamuser","created_at":"2018-01-05T23:59:58.160Z","updated_at":"2018-01-05T23:59:58.160Z","team_id":"team","user_id":"user","role":"admin","status":"active"}}`
	SampleUpdateTeamUser = `{"data":{"id":"someid","created_at":"2018-01-05T23:59:58.160Z","updated_at":"2018-01-05T23:59:58.160Z","team_id":"team","user_id":"user","role":"member","status":"active"}}`
)
