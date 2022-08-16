package scans

import (
	"encoding/json"
	"testing"

	"github.com/franela/goblin"
	"github.com/ion-channel/ionic/dependencies"
	"github.com/ion-channel/ionic/risk"
	"github.com/ion-channel/ionic/secrets"
	. "github.com/onsi/gomega"
)

func TestScanResults(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Untranslated Scan Results", func() {
		g.It("should parse untranslated scan results", func() {
			var untranslatedLicenseResult UntranslatedResults
			err := json.Unmarshal([]byte(SampleValidUntranslatedScanResultsLicense), &untranslatedLicenseResult)

			// validate the json parsing
			Expect(err).NotTo(HaveOccurred())
			Expect(untranslatedLicenseResult.License).NotTo(BeNil())

			var untranslatedVulnerabilityResult UntranslatedResults
			err = json.Unmarshal([]byte(SampleValidUntranslatedScanResultsVulnerability), &untranslatedVulnerabilityResult)
			// validate the json parsing
			Expect(err).To(BeNil())
			Expect(untranslatedVulnerabilityResult.Vulnerability).NotTo(BeNil())

			var untranslatedVirusResult UntranslatedResults
			err = json.Unmarshal([]byte(SampleValidUntranslatedScanResultsVirus), &untranslatedVirusResult)
			// validate the json parsing
			Expect(err).NotTo(HaveOccurred())
			Expect(untranslatedVirusResult.Virus).NotTo(BeNil())

			var untranslatedCommunityResult UntranslatedResults
			err = json.Unmarshal([]byte(SampleValidUntranslatedResultsCommunity), &untranslatedCommunityResult)
			// validate the json parsing
			Expect(err).NotTo(HaveOccurred())
			Expect(untranslatedCommunityResult.Community).NotTo(BeNil())
			Expect(untranslatedCommunityResult.Community.StarsTotalCount).To(Equal(2))

			var untranslatedMetricsResult UntranslatedResults
			err = json.Unmarshal([]byte(SampleValidUntranslatedScanResultsMetrics), &untranslatedMetricsResult)
			// validate the json parsing
			Expect(err).NotTo(HaveOccurred())
			Expect(untranslatedMetricsResult.Metrics).NotTo(BeNil())
			Expect(untranslatedMetricsResult.Metrics.Metrics.ID).To(Equal("pkg:github/yuchi/java-npm-semver"))

			var untranslatedRiskResult UntranslatedResults
			err = json.Unmarshal([]byte(SampleValidUntranslatedScanResultsRisk), &untranslatedRiskResult)
			// validate the json parsing
			Expect(err).NotTo(HaveOccurred())
			Expect(untranslatedRiskResult.Risk).NotTo(BeNil())
			Expect(untranslatedRiskResult.Risk.Risk[0].Value).To(Equal(35.45175730754987))

			var untranslatedExternalVulnerabilityResult UntranslatedResults
			err = json.Unmarshal([]byte(SampleValidUntranslatedResultsExternalVulnerability), &untranslatedExternalVulnerabilityResult)
			// validate the json parsing
			Expect(err).NotTo(HaveOccurred())
			Expect(untranslatedExternalVulnerabilityResult.ExternalVulnerabilities).NotTo(BeNil())
		})

		g.It("should translate untranslated scan results", func() {
			var untranslatedResult UntranslatedResults
			err := json.Unmarshal([]byte(SampleValidUntranslatedScanResultsLicense), &untranslatedResult)

			// validate the json parsing
			Expect(err).NotTo(HaveOccurred())
			Expect(untranslatedResult.AboutYML).To(BeNil())
			Expect(untranslatedResult.Community).To(BeNil())
			Expect(untranslatedResult.Coverage).To(BeNil())
			Expect(untranslatedResult.Dependency).To(BeNil())
			Expect(untranslatedResult.Difference).To(BeNil())
			Expect(untranslatedResult.Ecosystem).To(BeNil())
			Expect(untranslatedResult.ExternalVulnerabilities).To(BeNil())
			Expect(untranslatedResult.Vulnerability).To(BeNil())
			Expect(untranslatedResult.License).NotTo(BeNil())
			license := untranslatedResult.License
			Expect(license.Name).To(Equal("some license"))
			Expect(license.Type).To(HaveLen(1))
			Expect(license.Type[0].Name).To(Equal("a license"))

			// translate it
			translatedResult := untranslatedResult.Translate()

			// validate translated object
			Expect(translatedResult).NotTo(BeNil())
			Expect(translatedResult.Type).To(Equal("license"))
			Expect(translatedResult.Data).NotTo(BeNil())

			licenseResults, ok := translatedResult.Data.(LicenseResults)
			Expect(ok).To(BeTrue(), "Expected LicenseResults type")
			Expect(licenseResults.Type).To(HaveLen(1))
			Expect(licenseResults.Type[0].Name).To(Equal("a license"))
			Expect(licenseResults.Name).To(Equal("some license"))
			Expect(licenseResults.License.Type).To(HaveLen(1))
			Expect(licenseResults.License.Type[0].Name).To(Equal("a license"))
		})
	})

	g.Describe("Translated Scan Results", func() {
		g.It("should unmarshal a scan results with about yml data", func() {
			var r TranslatedResults
			err := json.Unmarshal([]byte(SampleValidScanResultsAboutYML), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("about_yml"))

			a, ok := r.Data.(AboutYMLResults)
			Expect(ok).To(Equal(true))
			Expect(a.Content).To(Equal("some content"))
		})

		g.It("should unmarshal a scan results with community data", func() {
			var r TranslatedResults
			err := json.Unmarshal([]byte(SampleValidScanResultsCommunity), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("community"))

			a, ok := r.Data.(CommunityResults)
			Expect(ok).To(Equal(true))
			Expect(a.CommittersTotalCount).To(Equal(7))
			Expect(a.Name).To(Equal("ion-channel/ion-connect"))
			Expect(a.URL).To(Equal("https://github.com/ion-channel/ion-connect"))
			Expect(a.StarsTotalCount).To(Equal(2))
			Expect(a.OldNames).To(Equal([]string{"old/name"}))
		})

		g.It("should unmarshal a scan results with coverage data", func() {
			var r TranslatedResults
			err := json.Unmarshal([]byte(SampleValidScanResultsCoverage), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("external_coverage"))

			c, ok := r.Data.(CoverageResults)
			Expect(ok).To(Equal(true))
			Expect(c.Value).To(Equal(42.0))
		})

		g.It("should unmarshal a scan results with dependency data", func() {
			var r TranslatedResults
			err := json.Unmarshal([]byte(SampleValidScanResultsDependency), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("dependency"))

			d, ok := r.Data.(DependencyResults)
			Expect(ok).To(Equal(true))
			Expect(len(d.Dependencies)).To(Equal(7))
			Expect(d.Meta.FirstDegreeCount).To(Equal(3))
			Expect(d.Meta.NoVersionCount).To(Equal(0))
			Expect(d.Meta.TotalUniqueCount).To(Equal(7))
			Expect(d.Meta.UpdateAvailableCount).To(Equal(2))
		})

		g.It("should unmarshal a scan results with difference data", func() {
			var r TranslatedResults
			err := json.Unmarshal([]byte(SampleValidScanResultsDifference), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("difference"))

			d, ok := r.Data.(DifferenceResults)
			Expect(ok).To(Equal(true))
			Expect(d.Checksum).To(Equal("checksumishere"))
			Expect(d.Difference).To(BeTrue())
		})

		g.It("should unmarshal a scan results with buildsystem data", func() {
			var r TranslatedResults
			err := json.Unmarshal([]byte(SampleValidScanResultsBuildsystems), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("buildsystems"))

			e, ok := r.Data.(BuildsystemResults)
			Expect(ok).To(Equal(true))
			Expect(len(e.Compilers)).To(Equal(1))
		})

		g.It("should marshal a scan result with buildsystem data", func() {
			r := &TranslatedResults{
				Type: "buildsystems",
				Data: BuildsystemResults{
					Compilers: []Compiler{
						Compiler{
							Name:    "Go",
							Version: "1.1.0",
						},
					},
					Dockerfile: Dockerfile{
						Images: []Image{
							Image{
								Name:    "golang",
								Version: "1.1.0",
							},
						},
						Dependencies: []dependencies.Dependency{
							dependencies.Dependency{
								Name:    "bash",
								Version: "",
							},
							dependencies.Dependency{
								Name:    "build-base",
								Version: "",
							},
							dependencies.Dependency{
								Name:    "curl",
								Version: "",
							},
						},
					},
				},
			}

			b, err := json.Marshal(r)
			Expect(err).To(BeNil())
			Expect(string(b)).To(Equal(SampleValidScanResultsBuildsystems))
		})

		g.It("should unmarshal a scan results with ecosystem data", func() {
			var r TranslatedResults
			err := json.Unmarshal([]byte(SampleValidScanResultsEcosystems), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("ecosystems"))

			e, ok := r.Data.(EcosystemResults)
			Expect(ok).To(Equal(true))
			Expect(len(e.Ecosystems)).To(Equal(3))
		})

		g.It("should marshal a scan result with ecosystem data", func() {
			r := &TranslatedResults{
				Type: "ecosystems",
				Data: EcosystemResults{
					Ecosystems: map[string]int{
						"Java":     2430,
						"Makefile": 210,
						"Ruby":     666,
					},
				},
			}

			b, err := json.Marshal(r)
			Expect(err).To(BeNil())
			Expect(string(b)).To(Equal(SampleValidScanResultsEcosystems))
		})

		g.It("should unmarshal a scan results with external vulnerabilities scan data", func() {
			var r TranslatedResults
			err := json.Unmarshal([]byte(SampleValidExternalVulnerabilities), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("external_vulnerability"))

			e, ok := r.Data.(ExternalVulnerabilitiesResults)
			Expect(ok).To(Equal(true))
			Expect(e.Critical).To(Equal(1))
			Expect(e.High).To(Equal(0))
			Expect(e.Medium).To(Equal(1))
			Expect(e.Low).To(Equal(0))
		})

		g.It("should unmarshal a scan results with license data", func() {
			var r TranslatedResults
			err := json.Unmarshal([]byte(SampleValidScanResultsLicense), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("license"))

			l, ok := r.Data.(LicenseResults)
			Expect(ok).To(Equal(true))
			Expect(l.License.Name).To(Equal("Not found"))
		})

		g.It("should unmarshal a scan results with metrics data", func() {
			var r TranslatedResults
			err := json.Unmarshal([]byte(SampleValidScanResultsMetrics), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("metrics"))

			e, ok := r.Data.(MetricsResults)
			Expect(ok).To(Equal(true))
			Expect(len(e.Metrics.Metrics)).ToNot(Equal(0))
		})

		g.It("should marshal a scan result with metrics data", func() {
			var r TranslatedResults
			err := json.Unmarshal([]byte(SampleValidScanResultsMetrics), &r)

			b, err := json.Marshal(r)
			Expect(err).To(BeNil())
			Expect(string(b)).To(Equal("{\"type\":\"metrics\",\"data\":{\"id\":\"pkg:github/yuchi/java-npm-semver\",\"metrics\":[{\"name\":\"committers_total_count\",\"bindings\":[{\"metric\":\"committers_total_count\",\"scope\":\"ecosystem\",\"category\":\"maintenance\",\"attribute\":\"size\",\"source\":\"github\"}],\"severity\":\"\",\"severity_rank\":0,\"value\":1,\"type\":\"\",\"sources\":[\"NVD\"],\"related_metrics\":null}]}}"))
		})

		g.It("should unmarshal a scan results with risk data", func() {
			var r TranslatedResults
			err := json.Unmarshal([]byte(SampleValidScanResultsRisk), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("risk"))

			e, ok := r.Data.(RiskResults)
			Expect(ok).To(Equal(true))
			Expect(len(e.Risk)).To(Equal(1))
		})

		g.It("should marshal a scan result with risk data", func() {
			r := &TranslatedResults{
				Type: "risk",
				Data: RiskResults{
					Risk: []risk.Scores{
						risk.Scores{
							Name:   "Software",
							Scopes: nil,
							Value:  0.0,
						},
					},
				},
			}

			b, err := json.Marshal(r)
			Expect(err).To(BeNil())
			Expect(string(b)).To(Equal(SampleValidScanResultsRisk))
		})

		g.It("should unmarshal a scan results with secrets data", func() {
			var r TranslatedResults
			err := json.Unmarshal([]byte(SampleValidScanResultsSecrets), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("secrets"))

			e, ok := r.Data.(SecretResults)
			Expect(ok).To(Equal(true))
			Expect(len(e.Secrets)).To(Equal(1))
		})

		g.It("should marshal a scan result with secrets data", func() {
			r := &TranslatedResults{
				Type: "secrets",
				Data: SecretResults{
					Secrets: []Secret{
						Secret{
							Secret: secrets.Secret{
								Rule:       "Slack Webhook",
								Match:      "\t\t\thttps://hooks.slack.com/services/T0F0****************************************",
								Confidence: 1,
							},
							File: "text.txt",
						},
					},
				},
			}

			b, err := json.Marshal(r)
			Expect(err).To(BeNil())
			Expect(string(b)).To(Equal(SampleValidScanResultsSecrets))
		})

		g.It("should unmarshal a scan results with virus data", func() {
			var r TranslatedResults
			err := json.Unmarshal([]byte(SampleValidScanResultsVirus), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("virus"))

			v, ok := r.Data.(VirusResults)
			Expect(ok).To(Equal(true))

			fn := FileNotes{
				"empty_file": []string{"file1", "file2", "file3"},
				"file1":      []string{"path/to/file"},
			}
			Expect(v.FileNotes).To(Equal(fn))
			cd := ClamavDetails{
				ClamavDbVersion: "1.1.0",
				ClamavVersion:   "1.0.0",
			}
			Expect(v.ClamavDetails).To(Equal(cd))
			Expect(v.KnownViruses).To(Equal(10))
		})

		g.It("should unmarshal a scan results with vulnerability data", func() {
			var r TranslatedResults
			err := json.Unmarshal([]byte(SampleValidScanResultsVulnerability), &r)

			Expect(err).To(BeNil())
			Expect(r.Type).To(Equal("vulnerability"))

			v, ok := r.Data.(VulnerabilityResults)
			Expect(ok).To(Equal(true))
			Expect(v.Meta.VulnerabilityCount).To(Equal(1))
			Expect(v.Vulnerabilities[0].Query.Name).To(Equal("broken"))
		})

		g.It("should return an error for an invalid results type", func() {
			var r TranslatedResults
			err := json.Unmarshal([]byte(SampleInvalidResults), &r)

			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("unsupported results type found:"))
		})
	})
}

const (
	SampleValidScanResultsAboutYML = `{"type":"about_yml", "data":{"message": "foo message", "valid": true, "content": "some content"}}`

	SampleValidScanResultsBuildsystems  = `{"type":"buildsystems","data":{"compilers":[{"name":"Go","version":"1.1.0"}],"docker_file":{"images":[{"name":"golang","version":"1.1.0"}],"dependencies":[{"name":"bash","version":"","latest_version":"","org":"","type":"","package":"","scope":"","requirement":"","dependencies":null,"confidence":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","outdated_version":{"major_behind":0,"minor_behind":0,"patch_behind":0}},{"name":"build-base","version":"","latest_version":"","org":"","type":"","package":"","scope":"","requirement":"","dependencies":null,"confidence":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","outdated_version":{"major_behind":0,"minor_behind":0,"patch_behind":0}},{"name":"curl","version":"","latest_version":"","org":"","type":"","package":"","scope":"","requirement":"","dependencies":null,"confidence":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","outdated_version":{"major_behind":0,"minor_behind":0,"patch_behind":0}}]}}}`
	SampleValidScanResultsCommunity     = `{"type":"community", "data":{"old_names":["old/name"],"stars_total_count":2,"committers_total_count":7,"name":"ion-channel/ion-connect","url":"https://github.com/ion-channel/ion-connect"}}`
	SampleValidScanResultsCoverage      = `{"type":"external_coverage", "data":{"value":42.0}}`
	SampleValidScanResultsDependency    = `{"type":"dependency","data":{"dependencies":[{"requirement":">1.0","latest_version":"2.0","org":"net.sourceforge.javacsv","name":"javacsv","type":"maven","package":"jar","version":"2.0","scope":"compile"},{"latest_version":"4.12","org":"junit","name":"junit","type":"maven","package":"jar","version":"4.11","scope":"test"},{"latest_version":"1.4-atlassian-1","org":"org.hamcrest","name":"hamcrest-core","type":"maven","package":"jar","version":"1.3","scope":"test"},{"latest_version":"4.5.2","org":"org.apache.httpcomponents","name":"httpclient","type":"maven","package":"jar","version":"4.3.4","scope":"compile"},{"latest_version":"4.4.5","org":"org.apache.httpcomponents","name":"httpcore","type":"maven","package":"jar","version":"4.3.2","scope":"compile"},{"latest_version":"99.0-does-not-exist","org":"commons-logging","name":"commons-logging","type":"maven","package":"jar","version":"1.1.3","scope":"compile"},{"latest_version":"20041127.091804","org":"commons-codec","name":"commons-codec","type":"maven","package":"jar","version":"1.6","scope":"compile"}],"meta":{"first_degree_count":3,"no_version_count":0,"total_unique_count":7,"update_available_count":2}}}`
	SampleValidScanResultsEcosystems    = `{"type":"ecosystems","data":{"Java":2430,"Makefile":210,"Ruby":666}}`
	SampleValidScanResultsLicense       = `{"type":"license","data":{"license":{"name":"Not found","type":[]}}}`
	SampleValidScanResultsMetrics       = `{"type":"metrics","data":{"id":"pkg:github/yuchi/java-npm-semver","metrics":[{"name":"committers_total_count","value":1,"sources":["NVD"],"bindings":[{"metric":"committers_total_count","scope":"ecosystem","category":"maintenance","attribute":"size","source":"github"}]}]}}`
	SampleValidScanResultsRisk          = `{"type":"risk","data":[{"name":"Software","value":0,"scopes":null}]}`
	SampleValidScanResultsSecrets       = `{"type":"secrets","data":[{"rule":"Slack Webhook","match":"\t\t\thttps://hooks.slack.com/services/T0F0****************************************","confidence":1,"file":"text.txt"}]}`
	SampleValidScanResultsVirus         = `{"type":"virus","data":{"known_viruses":10,"engine_version":"","scanned_directories":1,"scanned_files":2,"infected_files":1,"data_scanned":"some cool data was scanned","data_read":"we read some data","time":"10PM","file_notes": {"empty_file":["file1","file2","file3"], "file1": ["path/to/file"]},"clam_av_details":{"clamav_version":"1.0.0","clamav_db_version":"1.1.0"}}}`
	SampleValidScanResultsVulnerability = `{"type":"vulnerability","data":{"vulnerabilities":[{"id":316274974,"name":"hadoop","org":"apache","version":"2.8.0","up":null,"edition":null,"aliases":null,"created_at":"2017-02-13T20:02:32.785Z","updated_at":"2017-02-13T20:02:32.785Z","title":null,"references":null,"part":null,"language":null,"source_id":1,"external_id":"cpe:/a:apache:hadoop:2.8.0","vulnerabilities":[{"id":92596,"external_id":"CVE-2017-7669","title":"CVE-2017-7669","summary":"In Apache Hadoop 2.8.0, 3.0.0-alpha1, and 3.0.0-alpha2, the LinuxContainerExecutor runs docker commands as root with insufficient input validation. When the docker feature is enabled, authenticated users can run commands as root.","score":"8.5","score_version":"2.0","score_system":"CVSS","score_details":{"cvssv2":{"vectorString":"(AV:N/AC:M/Au:S/C:C/I:C/A:C)","accessVector":"NETWORK","accessComplexity":"MEDIUM","authentication":"SINGLE","confidentialityImpact":"COMPLETE","integrityImpact":"COMPLETE","availabilityImpact":"COMPLETE","baseScore":8.5},"cvssv3":{"vectorString":"AV:N/AC:H/PR:L/UI:N/S:U/C:H/I:H/A:H","attackVector":"NETWORK","attackComplexity":"HIGH","privilegesRequired":"LOW","userInteraction":"NONE","scope":"UNCHANGED","confidentialityImpact":"HIGH","integrityImpact":"HIGH","availabilityImpact":"HIGH","baseScore":7.5,"baseSeverity":"HIGH"}},"vector":"NETWORK","access_complexity":"MEDIUM","vulnerability_authentication":"SINGLE","confidentiality_impact":"COMPLETE","integrity_impact":"COMPLETE","availability_impact":"COMPLETE","vulnerability_source":null,"assessment_check":null,"scanner":null,"recommendation":"","references":[{"type":"UNKNOWN","source":"","url":"http://www.securityfocus.com/bid/98795","text":"http://www.securityfocus.com/bid/98795"},{"type":"UNKNOWN","source":"","url":"https://mail-archives.apache.org/mod_mbox/hadoop-user/201706.mbox/%3C4A2FDA56-491B-4C2A-915F-C9D4A4BDB92A%40apache.org%3E","text":"https://mail-archives.apache.org/mod_mbox/hadoop-user/201706.mbox/%3C4A2FDA56-491B-4C2A-915F-C9D4A4BDB92A%40apache.org%3E"}],"modified_at":"2017-06-09T16:21:00.000Z","published_at":"2017-06-05T01:29:00.000Z","created_at":"2017-07-12T23:07:35.491Z","updated_at":"2017-07-12T23:07:35.491Z","source_id":1}],"query":{"name":"broken"}}],"meta":{"vulnerability_count":1}}}`
	SampleInvalidResults                = `{"type":"fooresult", "data":"I pitty the foo"}`
	SampleValidScanResultsDifference    = `{"data": {"checksum": "checksumishere","difference": true},"type": "difference"}`
	SampleValidExternalVulnerabilities  = `{"type":"external_vulnerability","data":{"critical":1,"high":0,"medium":1,"low": 0}}`

	SampleValidUntranslatedResultsExternalVulnerability = `{"external_vulnerability":{"critical":43, "high":262, "medium":0, "low":79}, "source":{"name":"Fortify", "url":""}, "notes":"", "raw":{"fpr":"/ion/fortify.zip"}}`
	SampleValidUntranslatedResultsCommunity             = `{"type":"community", "data":{"old_names":["old/name"],"stars_total_count":2,"committers_total_count":7,"name":"ion-channel/ion-connect","url":"https://github.com/ion-channel/ion-connect"}}`
	SampleValidUntranslatedScanResultsLicense           = `{"license": {"license": {"type": [{"name": "a license"}], "name": "some license"}}}`
	SampleValidUntranslatedScanResultsMetrics           = `{"metrics": {"id":"pkg:github/yuchi/java-npm-semver","metrics":[{"name":"committers_total_count","value":1,"bindings":[{"metric":"committers_total_count","scope":"ecosystem","category":"maintenance","attribute":"size","source":"github"},{"metric":"committers_total_count","scope":"ecosystem","category":"sustainability","attribute":"visibility","source":"github"}]},{"name":"actors_total_count","value":4,"bindings":[{"metric":"actors_total_count","scope":"ecosystem","category":"maintenance","attribute":"size","source":"github"},{"metric":"actors_total_count","scope":"ecosystem","category":"sustainability","attribute":"visibility","source":"github"},{"metric":"actors_total_count","scope":"supply_chain","category":"reputation","attribute":"popularity","source":"github"}]},{"name":"committers_monthly_count","value":[{"month":"08-2017","count":2},{"month":"12-2016","count":3}],"bindings":[]},{"name":"releases_total_count","value":0,"bindings":[{"metric":"releases_total_count","scope":"ecosystem","category":"maintenance","attribute":"activeness","source":"github"},{"metric":"releases_total_count","scope":"ecosystem","category":"maintenance","attribute":"responsiveness","source":"github"},{"metric":"releases_total_count","scope":"ecosystem","category":"sustainability","attribute":"process_maturity","source":"github"},{"metric":"releases_total_count","scope":"technology","category":"reliability","attribute":"effectiveness","source":"github"},{"metric":"releases_total_count","scope":"technology","category":"reliability","attribute":"maturity","source":"github"},{"metric":"releases_total_count","scope":"technology","category":"reliability","attribute":"usability","source":"github"}]},{"name":"releases_monthly_count","value":[],"bindings":[]},{"name":"releases_last_at","value":"0001-01-01T00:00:00Z","bindings":[]},{"name":"pull_requests_total_count","value":0,"bindings":[{"metric":"pull_requests_total_count","scope":"ecosystem","category":"sustainability","attribute":"process_maturity","source":"github"}]},{"name":"pull_requests_last_at","value":"0001-01-01T00:00:00Z","bindings":[]},{"name":"pull_requests_monthly_count","value":[],"bindings":[]},{"name":"issues_last_at","value":"2020-07-09T16:16:02Z","bindings":[]},{"name":"issues_open_monthly_count","value":[{"month":"08-2017","count":1}],"bindings":[]},{"name":"issues_closed_monthly_count","value":[{"month":"08-2017","count":1}],"bindings":[]},{"name":"issues_closed_mttr_monthly","value":[{"month":"2017-08","mttr":0}],"bindings":[]},{"name":"issues_closed_mttr","value":-172800,"bindings":[{"metric":"issues_closed_mttr","scope":"ecosystem","category":"maintenance","attribute":"responsiveness","source":"github"},{"metric":"issues_closed_mttr","scope":"ecosystem","category":"sustainability","attribute":"process_maturity","source":"github"},{"metric":"issues_closed_mttr","scope":"technology","category":"reliability","attribute":"effectiveness","source":"github"},{"metric":"issues_closed_mttr","scope":"technology","category":"reliability","attribute":"usability","source":"github"}]},{"name":"commits_total_count","value":5,"bindings":[{"metric":"commits_total_count","scope":"ecosystem","category":"maintenance","attribute":"activeness","source":"github"}]},{"name":"commits_monthly_count","value":[{"month":"08-2017","count":2},{"month":"12-2016","count":3}],"bindings":[]},{"name":"actors_monthly_count","value":[{"month":"11-2019","count":2},{"month":"08-2017","count":2},{"month":"12-2016","count":1},{"month":"07-2020","count":1}],"bindings":[]},{"name":"actions_total_count","value":23,"bindings":[{"metric":"actions_total_count","scope":"ecosystem","category":"maintenance","attribute":"activeness","source":"github"}]},{"name":"actions_last_at","value":"2020-07-09T16:16:02Z","bindings":[]},{"name":"actions_first_at","value":"2016-12-18T00:00:00Z","bindings":[]},{"name":"actions_monthly_count","value":[{"month":"07-2020","count":1},{"month":"12-2016","count":4},{"month":"11-2019","count":6},{"month":"08-2017","count":12}],"bindings":[]},{"name":"contributing_actors_total_count","value":3,"bindings":[{"metric":"contributing_actors_total_count","scope":"ecosystem","category":"maintenance","attribute":"size","source":"github"}]},{"name":"contributing_actors_monthly_count","value":[{"month":"11-2019","count":1},{"month":"12-2016","count":1},{"month":"08-2017","count":2}],"bindings":[]},{"name":"contributing_actions_total_count","value":9,"bindings":[{"metric":"contributing_actions_total_count","scope":"ecosystem","category":"maintenance","attribute":"activeness","source":"github"}]},{"name":"contributing_actions_last_at","value":"2019-11-25T16:44:27Z","bindings":[]},{"name":"contributing_actions_monthly_count","value":[{"month":"11-2019","count":2},{"month":"08-2017","count":3},{"month":"12-2016","count":4}],"bindings":[]},{"name":"new_actors_monthly_count","value":[{"month":"2020-07","count":1},{"month":"2019-11","count":1},{"month":"2017-08","count":1},{"month":"2016-12","count":1}],"bindings":[]},{"name":"median_working_hour","value":16,"bindings":[]},{"name":"average_monthly_actions","value":0.367905056772835,"bindings":[{"metric":"average_monthly_actions","scope":"ecosystem","category":"maintenance","attribute":"activeness","source":"github"}]},{"name":"average_monthly_actors","value":0.095975232201609,"bindings":[{"metric":"average_monthly_actors","scope":"ecosystem","category":"maintenance","attribute":"size","source":"github"},{"metric":"average_monthly_actors","scope":"ecosystem","category":"sustainability","attribute":"regeneration_ability","source":"github"},{"metric":"average_monthly_actors","scope":"ecosystem","category":"sustainability","attribute":"visibility","source":"github"},{"metric":"average_monthly_actors","scope":"supply_chain","category":"reputation","attribute":"popularity","source":"github"}]},{"name":"average_monthly_commits","value":0.0799793601680075,"bindings":[{"metric":"average_monthly_commits","scope":"ecosystem","category":"maintenance","attribute":"activeness","source":"github"}]},{"name":"average_monthly_committers","value":0.0479876161008045,"bindings":[{"metric":"average_monthly_committers","scope":"ecosystem","category":"maintenance","attribute":"size","source":"github"},{"metric":"average_monthly_committers","scope":"ecosystem","category":"sustainability","attribute":"regeneration_ability","source":"github"},{"metric":"average_monthly_committers","scope":"ecosystem","category":"sustainability","attribute":"visibility","source":"github"}]},{"name":"average_monthly_contributing_actions","value":0.143962848302414,"bindings":[{"metric":"average_monthly_contributing_actions","scope":"ecosystem","category":"maintenance","attribute":"activeness","source":"github"}]},{"name":"average_monthly_contributing_actors","value":0.063983488134406,"bindings":[{"metric":"average_monthly_contributing_actors","scope":"ecosystem","category":"maintenance","attribute":"size","source":"github"},{"metric":"average_monthly_contributing_actors","scope":"ecosystem","category":"sustainability","attribute":"regeneration_ability","source":"github"},{"metric":"average_monthly_contributing_actors","scope":"ecosystem","category":"sustainability","attribute":"visibility","source":"github"}]},{"name":"average_monthly_growth_new_actors","value":0,"bindings":[{"metric":"average_monthly_growth_new_actors","scope":"ecosystem","category":"sustainability","attribute":"regeneration_ability","source":"github"},{"metric":"average_monthly_growth_new_actors","scope":"ecosystem","category":"sustainability","attribute":"visibility","source":"github"}]},{"name":"average_monthly_growth_new_contributing_actors","value":0,"bindings":[{"metric":"average_monthly_growth_new_contributing_actors","scope":"ecosystem","category":"sustainability","attribute":"regeneration_ability","source":"github"},{"metric":"average_monthly_growth_new_contributing_actors","scope":"ecosystem","category":"sustainability","attribute":"visibility","source":"github"}]},{"name":"average_monthly_issues_closed","value":0.0159958720336015,"bindings":[{"metric":"average_monthly_issues_closed","scope":"ecosystem","category":"maintenance","attribute":"activeness","source":"github"},{"metric":"average_monthly_issues_closed","scope":"ecosystem","category":"maintenance","attribute":"responsiveness","source":"github"}]},{"name":"average_monthly_issues_open","value":0.0479876161008045,"bindings":[{"metric":"average_monthly_issues_open","scope":"ecosystem","category":"sustainability","attribute":"visibility","source":"github"}]},{"name":"average_monthly_new_actors","value":4,"bindings":[]},{"name":"average_monthly_passive_actions","value":0.223942208470421,"bindings":[{"metric":"average_monthly_passive_actions","scope":"ecosystem","category":"maintenance","attribute":"activeness","source":"github"},{"metric":"average_monthly_passive_actions","scope":"ecosystem","category":"sustainability","attribute":"visibility","source":"github"}]},{"name":"average_monthly_passive_actors","value":0.0799793601680075,"bindings":[{"metric":"average_monthly_passive_actors","scope":"ecosystem","category":"maintenance","attribute":"size","source":"github"},{"metric":"average_monthly_passive_actors","scope":"ecosystem","category":"sustainability","attribute":"regeneration_ability","source":"github"},{"metric":"average_monthly_passive_actors","scope":"ecosystem","category":"sustainability","attribute":"visibility","source":"github"},{"metric":"average_monthly_passive_actors","scope":"supply_chain","category":"reputation","attribute":"popularity","source":"github"}]},{"name":"average_monthly_pr_comments","value":0,"bindings":[{"metric":"average_monthly_pr_comments","scope":"ecosystem","category":"sustainability","attribute":"process_maturity","source":"github"}]},{"name":"average_monthly_pull_requests","value":0,"bindings":
	[{"metric":"average_monthly_pull_requests","scope":"ecosystem","category":"sustainability","attribute":"process_maturity","source":"github"}]},{"name":"average_monthly_releases","value":0,"bindings":[{"metric":"average_monthly_releases","scope":"ecosystem","category":"maintenance","attribute":"activeness","source":"github"},{"metric":"average_monthly_releases","scope":"ecosystem","category":"maintenance","attribute":"responsiveness","source":"github"},{"metric":"average_monthly_releases","scope":"ecosystem","category":"sustainability","attribute":"process_maturity","source":"github"},{"metric":"average_monthly_releases","scope":"technology","category":"reliability","attribute":"effectiveness","source":"github"},{"metric":"average_monthly_releases","scope":"technology","category":"reliability","attribute":"maturity","source":"github"},{"metric":"average_monthly_releases","scope":"technology","category":"reliability","attribute":"usability","source":"github"}]},{"name":"contributing_actor_activity_rate","value":1,"bindings":[{"metric":"contributing_actor_activity_rate","scope":"ecosystem","category":"sustainability","attribute":"workload_balance","source":"github"}]},{"name":"median_working_hour_diversity_index","value":1.03972,"bindings":[{"metric":"median_working_hour_diversity_index","scope":"ecosystem","category":"sustainability","attribute":"heterogeniety","source":"github"}]},{"name":"passive_actions_total_count","value":14,"bindings":[{"metric":"passive_actions_total_count","scope":"ecosystem","category":"maintenance","attribute":"activeness","source":"github"}]},{"name":"passive_actors_total_count","value":4,"bindings":[{"metric":"passive_actors_total_count","scope":"ecosystem","category":"maintenance","attribute":"size","source":"github"},{"metric":"passive_actors_total_count","scope":"ecosystem","category":"sustainability","attribute":"visibility","source":"github"},{"metric":"passive_actors_total_count","scope":"supply_chain","category":"reputation","attribute":"popularity","source":"github"}]},{"name":"pr_comment_total_count","value":0,"bindings":[{"metric":"pr_comment_total_count","scope":"ecosystem","category":"sustainability","attribute":"process_maturity","source":"github"}]},{"name":"actor_role_diversity_index","value":0,"bindings":[{"metric":"actor_role_diversity_index","scope":"ecosystem","category":"sustainability","attribute":"heterogeniety","source":"github"},{"metric":"actor_role_diversity_index","scope":"ecosystem","category":"sustainability","attribute":"workload_balance","source":"github"}]},{"name":"contributing_actor_activity_rate_monthly","value":[{"month":"2016-12","count":1},{"month":"2017-08","count":0.16666667},{"month":"2019-11","count":1}],"bindings":[]},{"name":"passive_actors_monthly_count","value":[{"month":"2017-08","count":2},{"month":"2019-11","count":2},{"month":"2020-07","count":1}],"bindings":[]},{"name":"passive_actions_last_at","value":"2020-07-09T16:16:02Z","bindings":[]},{"name":"passive_actions_monthly_count","value":[{"month":"2017-08","count":9},{"month":"2019-11","count":4},{"month":"2020-07","count":1}],"bindings":[]},{"name":"new_contributing_actors_monthly_count","value":[],"bindings":[]},{"name":"pr_comment_monthly_count","value":[],"bindings":[]}]}}`

	SampleValidUntranslatedScanResultsRisk          = `{"risk": [{"name": "software", "value": 35.45175730754987, "scopes": null}]}`
	SampleValidUntranslatedScanResultsVulnerability = `{"vulnerabilities": {"meta": {"vulnerability_count": 0},"vulnerabilities": []}}`
	SampleValidUntranslatedScanResultsVirus         = `{"clam_av_details":{"clamav_db_version":"Tue Apr 24 12:26:01 2018\n","clamav_version":"ClamAV 0.99.4"},"clamav":{"data_read":"2.78 MB (ratio 1.68:1)","data_scanned":"4.66 MB","engine_version":"0.99.4","file_notes":{"empty_file":["/workspace/851c1261-471c-4713-bdc4-fabb0c2d0f6a/xunit-plugin-1-102/xunit-plugin-master/src/main/resources/util/taglib"]},"infected_files":0,"known_viruses":6480116,"scanned_directories":132,"scanned_files":305,"time":"18.655 sec (0 m 18 s)"}      }`
)
