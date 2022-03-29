package ionic

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/franela/goblin"
	"github.com/gomicro/bogus"
	. "github.com/onsi/gomega"
)

func TestScores(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
	g.Describe("Scores", func() {
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
		g.It("should lookup scores", func() {
			server.AddPath("/v1/score/getScores").
				SetMethods("POST").
				SetPayload([]byte(sampleValidScoresResponse)).
				SetStatus(http.StatusOK)
			scoreResults, err := client.GetScores([]string{"pkg:github/antlr/stringtemplate4", "pkg:github/asomov/snakeyaml", "pkg:github/fasterxml/java-classmate", "pkg:github/testvendor001/testproduct001"}, "sometoken")
			Expect(err).NotTo(HaveOccurred())
			Expect(scoreResults).To(HaveLen(4))

			hitRecords := server.HitRecords()
			Expect(hitRecords).To(HaveLen(1))
			Expect(hitRecords[0].Header.Get("Authorization")).To(Equal("Bearer sometoken"))
		})
	})
}

const (
	sampleValidScoresResponse = `{"data":[{"name":"pkg:github/antlr/stringtemplate4","value":0,"scopes":null},{"name":"pkg:github/asomov/snakeyaml","value":0,"scopes":null},{"name":"pkg:github/fasterxml/java-classmate","value":0,"scopes":null},{"name":"pkg:github/testvendor001/testproduct001","value":12.5,"scopes":[{"name":"ecosystem","value":12.5,"categories":[{"name":"sustainability","value":12.5,"attributes":[{"name":"visibility","value":12.5}]}]},{"name":"supply chain","value":12.5,"categories":[{"name":"reputation","value":12.5,"attributes":[{"name":"popularity","value":12.5},{"name":"public trustworthiness","value":12.5}]}]},{"name":"technology","value":12.5,"categories":[{"name":"reliability","value":12.5,"attributes":[{"name":"effectiveness","value":12.5},{"name":"usability","value":12.5}]}]}]}],"meta":{"total_count":0,"offset":0}}`
)
