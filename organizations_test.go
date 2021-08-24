package ionic

import (
	"fmt"
	"github.com/ion-channel/ionic/requests"
	"net/http"
	"testing"

	"github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestOrganizations(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Organizations", func() {
		server := bogus.New()
		h, p := server.HostPort()
		client, _ := New(fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get an organization", func() {
			server.AddPath("/v1/organizations/getOrganization/cd98e4e1-6926-4989-8ef8-f326cd5956fc").
				SetMethods("GET").
				SetPayload([]byte(SampleValidOrganization)).
				SetStatus(http.StatusOK)

			org, err := client.GetOrganization("cd98e4e1-6926-4989-8ef8-f326cd5956fc", "atoken")
			Expect(err).To(BeNil())
			Expect(org.ID).To(Equal("cd98e4e1-6926-4989-8ef8-f326cd5956fc"))
			Expect(org.Name).To(Equal("ion-channel"))
		})

		g.It("should get organizations in bulk", func() {
			server.AddPath("/v1/organizations/getOrganizations").
				SetMethods("POST").
				SetPayload([]byte(fmt.Sprintf(`{"data":[%v,%v]}`, SampleValidOrganization, SampleValidOrganization))).
				SetStatus(http.StatusOK)

			orgs, err := client.GetOrganizations(requests.ByIDs{IDs: []string{"some-id", "another-id"}}, "atoken")
			Expect(err).To(BeNil())
			Expect(len(*orgs)).To(Equal(2))
		})

		g.It("should create an organization", func() {
			server.AddPath("/v1/organizations/createOrganization").
				SetMethods("POST").
				SetPayload([]byte(SampleCreateOrganization)).
				SetStatus(http.StatusOK)

			opts := CreateOrganizationOptions{
				Name: "test-org",
			}
			org, err := client.CreateOrganization(opts, "atoken")
			Expect(err).To(BeNil())
			Expect(org.ID).To(Equal("5c4a8a84-efa0-4357-91f6-9f9e95f7dd1a"))
			Expect(org.Name).To(Equal("test-org"))
		})
	})
}

const (
	SampleValidOrganization  = `{"data":{"id":"cd98e4e1-6926-4989-8ef8-f326cd5956fc","created_at":"2016-09-09T22:06:49.487Z","updated_at":"2016-09-09T22:06:49.487Z","name":"ion-channel"}}`
	SampleCreateOrganization = `{"data":{"id":"5c4a8a84-efa0-4357-91f6-9f9e95f7dd1a","created_at":"2018-01-05T23:59:58.160Z","updated_at":"2018-01-05T23:59:58.160Z","name":"test-org"}}`
)
