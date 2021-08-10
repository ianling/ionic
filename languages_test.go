package ionic

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestLanguages(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
	g.Describe("Languages", func() {
		var server *bogus.Bogus
		var h, p string
		var client *IonClient
		g.BeforeEach(func() {
			server = bogus.New()
			h, p = server.HostPort()
			client, _ = New(fmt.Sprintf("http://%v:%v", h, p))
		})
		g.AfterEach(func() {
			server.Close()
		})
		g.It("should get languages", func() {
			server.AddPath("/v1/metadata/getLanguages").
				SetMethods("POST").
				SetPayload([]byte(sampleValidGetLanguages)).
				SetStatus(http.StatusOK)
			langs, err := client.GetLanguages("text with english", "sometoken")
			Expect(err).NotTo(HaveOccurred())
			Expect(langs).To(HaveLen(1))
			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			Expect(hitRecords[0].Header.Get("Authorization")).To(Equal("Bearer sometoken"))
			Expect(langs[0].Name).To(Equal("English"))
		})
	})
}

const sampleValidGetLanguages = `{"data":[{"name":"English","confidence":1}],"meta":{"total_count":1,"offset":0,"last_update":"0001-01-01T00:00:00Z"}}`
