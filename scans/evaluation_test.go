package scans

import (
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"testing"
)

func TestEvaluation(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Evaluation", func() {
		g.Describe("Translating", func() {
			g.It("should translate an untranslated evaluation", func() {
				ee := &Evaluation{
					UntranslatedResults: &UntranslatedResults{
						License: &LicenseResults{},
					},
				}
				Expect(ee.UntranslatedResults).NotTo(BeNil())
				Expect(ee.TranslatedResults).To(BeNil())

				err := ee.Translate()
				Expect(err).To(BeNil())
				Expect(ee.UntranslatedResults).To(BeNil())
				Expect(ee.TranslatedResults).NotTo(BeNil())
				Expect(ee.TranslatedResults.Type).To(Equal("license"))
				Expect(ee.Results).NotTo(BeNil())
				Expect(len(ee.Results)).NotTo(Equal(0))
			})

			g.It("should not translate an already translated summary", func() {
				ee := &Evaluation{
					UntranslatedResults: &UntranslatedResults{
						License: &LicenseResults{},
					},
				}
				Expect(ee.UntranslatedResults).NotTo(BeNil())
				Expect(ee.TranslatedResults).To(BeNil())

				err := ee.Translate()
				Expect(err).To(BeNil())
				Expect(ee.UntranslatedResults).To(BeNil())
				Expect(ee.TranslatedResults).NotTo(BeNil())
				Expect(ee.TranslatedResults.Type).To(Equal("license"))
				Expect(ee.Results).NotTo(BeNil())
				Expect(len(ee.Results)).NotTo(Equal(0))

				err = ee.Translate()
				Expect(err).To(BeNil())
				Expect(ee.UntranslatedResults).To(BeNil())
				Expect(ee.TranslatedResults).NotTo(BeNil())
				Expect(ee.TranslatedResults.Type).To(Equal("license"))
				Expect(ee.Results).NotTo(BeNil())
				Expect(len(ee.Results)).NotTo(Equal(0))
			})
		})
	})
}
