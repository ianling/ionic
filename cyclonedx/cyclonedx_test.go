package cyclonedx

import (
	"github.com/CycloneDX/cyclonedx-go"
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestSPDX(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("CycloneDX", func() {
		g.It("should return the top-level project when no dependencies requested", func() {
			metadata := cyclonedx.Metadata{}
			topLevelComponentExternalRefs := []cyclonedx.ExternalReference{{
				URL:  "git@github.com:ion-channel/my_app.git@some-branch",
				Type: "vcs",
			}}
			topLevelComponent := cyclonedx.Component{
				BOMRef:             "pkg:my_app",
				Type:               "application",
				Name:               "my_app",
				Version:            "1.2.3",
				Publisher:          "Ion Channel",
				ExternalReferences: &topLevelComponentExternalRefs,
			}
			metadata.Component = &topLevelComponent

			dependency := cyclonedx.Component{
				BOMRef:    "pkg:react_5.5.5",
				Type:      "library",
				Name:      "react",
				Version:   "5.5.5",
				Publisher: "facebook",
			}
			components := []cyclonedx.Component{dependency}

			doc := cyclonedx.NewBOM()
			doc.Metadata = &metadata
			doc.Components = &components

			projects, err := ProjectsFromCycloneDX(doc, false)

			Expect(err).To(BeNil())
			Expect(projects).NotTo(BeNil())
			Expect(len(projects)).To(Equal(1))
			Expect(*projects[0].Name).To(Equal(topLevelComponent.Name))
			Expect(*projects[0].Type).To(Equal("git"))
			Expect(len(projects[0].Aliases)).To(Equal(1))
			Expect(projects[0].Aliases[0].Name).To(Equal(topLevelComponent.Name))
			Expect(projects[0].Aliases[0].Org).To(Equal(topLevelComponent.Publisher))
			Expect(projects[0].Aliases[0].Version).To(Equal(topLevelComponent.Version))
			Expect(*projects[0].Branch).To(Equal("some-branch"))
			Expect(*projects[0].Source).To(Equal("git@github.com:ion-channel/my_app.git"))
		})

		g.It("should return all projects when dependencies requested", func() {
			metadata := cyclonedx.Metadata{}
			topLevelComponentExternalRefs := []cyclonedx.ExternalReference{{
				URL:  "git@github.com:ion-channel/my_app.git@some-branch",
				Type: "vcs",
			}}
			topLevelComponent := cyclonedx.Component{
				BOMRef:             "pkg:my_app",
				Type:               "application",
				Name:               "my_app",
				Version:            "1.2.3",
				Publisher:          "Ion Channel",
				ExternalReferences: &topLevelComponentExternalRefs,
			}
			metadata.Component = &topLevelComponent

			dependency := cyclonedx.Component{
				BOMRef:    "pkg:react_5.5.5",
				Type:      "library",
				Name:      "react",
				Version:   "5.5.5",
				Publisher: "facebook",
			}
			components := []cyclonedx.Component{dependency}

			doc := cyclonedx.NewBOM()
			doc.Metadata = &metadata
			doc.Components = &components

			projects, err := ProjectsFromCycloneDX(doc, true)

			Expect(err).To(BeNil())
			Expect(projects).NotTo(BeNil())
			Expect(len(projects)).To(Equal(2))
			Expect(*projects[0].Name).To(Equal(topLevelComponent.Name))
			Expect(*projects[0].Type).To(Equal("git"))
			Expect(len(projects[0].Aliases)).To(Equal(1))
			Expect(projects[0].Aliases[0].Name).To(Equal(topLevelComponent.Name))
			Expect(projects[0].Aliases[0].Org).To(Equal(topLevelComponent.Publisher))
			Expect(projects[0].Aliases[0].Version).To(Equal(topLevelComponent.Version))
			Expect(*projects[0].Branch).To(Equal("some-branch"))
			Expect(*projects[0].Source).To(Equal("git@github.com:ion-channel/my_app.git"))

			Expect(*projects[1].Name).To(Equal(dependency.Name))
			Expect(*projects[1].Type).To(Equal("source_unavailable"))
			Expect(len(projects[1].Aliases)).To(Equal(1))
			Expect(projects[1].Aliases[0].Name).To(Equal(dependency.Name))
			Expect(projects[1].Aliases[0].Org).To(Equal(dependency.Publisher))
			Expect(projects[1].Aliases[0].Version).To(Equal(dependency.Version))
			Expect(*projects[1].Branch).To(Equal(""))
			Expect(*projects[1].Source).To(Equal(""))
		})
	})
}
