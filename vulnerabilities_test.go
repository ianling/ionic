package ionic

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gomicro/bogus"
	"github.com/ion-channel/ionic/pagination"
	. "github.com/onsi/gomega"
)

func TestVulns(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Vulnerabilities", func() {
		server := bogus.New()
		server.Start()
		h, p := server.HostPort()
		client, _ := New("secret", fmt.Sprintf("http://%v:%v", h, p))

		g.It("should get vulnerabilities", func() {
			server.AddPath("/v1/vulnerability/getVulnerabilities").
				SetMethods("GET").
				SetPayload([]byte(SampleVulnerabilitiesResponse)).
				SetStatus(http.StatusOK)
			vulns, err := client.GetVulnerabilities("jdk", "", pagination.AllItems)

			Expect(err).To(BeNil())
			Expect(len(vulns)).To(Equal(21))
		})

		g.It("should get a vulnerability", func() {
			server.AddPath("/v1/vulnerability/getVulnerability").
				SetMethods("GET").
				SetPayload([]byte(SampleVulnerabilityResponse)).
				SetStatus(http.StatusOK)
			vuln, err := client.GetVulnerability("CVE-2013-4164")

			Expect(err).To(BeNil())
			Expect(vuln.Title).To(Equal("CVE-2013-4164"))
			Expect(vuln.Vector).To(Equal("NETWORK"))
			Expect(vuln.AccessComplexity).To(Equal("MEDIUM"))
		})

		g.It("should get a raw vulnerability", func() {
			server.AddPath("/v1/vulnerability/getVulnerability").
				SetMethods("GET").
				SetPayload([]byte(SampleVulnerabilityResponse)).
				SetStatus(http.StatusOK)
			bodyBytes, err := client.GetRawVulnerability("CVE-2013-4164")

			Expect(err).To(BeNil())
			Expect(string(bodyBytes)).To(ContainSubstring("CVE-2013-4164"))
		})
	})
}

const (
	SampleVulnerabilitiesResponse = `{"data":[{"id":2661,"external_id":"CVE-2000-1099","title":"CVE-2000-1099","summary":"Java Runtime Environment in Java Development Kit (JDK) 1.2.2_05 and earlier can allow an untrusted Java class to call into a disallowed class, which could allow an attacker to escape the Java sandbox and conduct unauthorized activities.","score":"5.1","vector":"NETWORK","access_complexity":"HIGH","vulnerability_authentication":"NONE","confidentiality_impact":"PARTIAL","integrity_impact":"PARTIAL","availability_impact":"PARTIAL","vulnerability_source":"http://nvd.nist.gov","assessment_check":null,"scanner":null,"recommendation":"","modified_at":"2008-09-10T15:06:37.993Z","published_at":"2001-01-09T00:00:00.000Z","created_at":"2017-02-10T23:51:01.262Z","updated_at":"2017-02-10T23:51:01.262Z","source_id":1},{"id":4403,"external_id":"CVE-2002-0058","title":"CVE-2002-0058","summary":"Vulnerability in Java Runtime Environment (JRE) allows remote malicious web sites to hijack or sniff a web client's sessions, when an HTTP proxy is being used, via a Java applet that redirects the session to another server, as seen in (1) Netscape 6.0 through 6.1 and 4.79 and earlier, (2) Microsoft VM build 3802 and earlier as used in Internet Explorer 4.x and 5.x, and possibly other implementations that use vulnerable versions of SDK or JDK.","score":"5.0","vector":"NETWORK","access_complexity":"LOW","vulnerability_authentication":"NONE","confidentiality_impact":"PARTIAL","integrity_impact":"NONE","availability_impact":"NONE","vulnerability_source":"http://nvd.nist.gov","assessment_check":null,"scanner":null,"recommendation":"","modified_at":"2008-09-10T15:11:10.383Z","published_at":"2002-03-15T00:00:00.000Z","created_at":"2017-02-10T23:51:01.910Z","updated_at":"2017-02-10T23:51:01.910Z","source_id":1},{"id":4421,"external_id":"CVE-2002-0076","title":"CVE-2002-0076","summary":"Java Runtime Environment (JRE) Bytecode Verifier allows remote attackers to escape the Java sandbox and execute commands via an applet containing an illegal cast operation, as seen in (1) Microsoft VM build 3802 and earlier as used in Internet Explorer 4.x and 5.x, (2) Netscape 6.2.1 and earlier, and possibly other implementations that use vulnerable versions of SDK or JDK, aka a variant of the \"Virtual Machine Verifier\" vulnerability.","score":"7.5","vector":"NETWORK","access_complexity":"LOW","vulnerability_authentication":"NONE","confidentiality_impact":"PARTIAL","integrity_impact":"PARTIAL","availability_impact":"PARTIAL","vulnerability_source":"http://nvd.nist.gov","assessment_check":null,"scanner":null,"recommendation":"","modified_at":"2008-09-05T16:27:06.967Z","published_at":"2002-03-19T00:00:00.000Z","created_at":"2017-02-10T23:51:01.910Z","updated_at":"2017-02-10T23:51:01.910Z","source_id":1},{"id":11343,"external_id":"CVE-2005-0471","title":"CVE-2005-0471","summary":"Sun Java JRE 1.1.x through 1.4.x writes temporary files with long filenames that become predictable on a file system that uses 8.3 style short names, which allows remote attackers to write arbitrary files to known locations and facilitates the exploitation of vulnerabilities in applications that rely on unpredictable file names.","score":"5.0","vector":"NETWORK","access_complexity":"LOW","vulnerability_authentication":"NONE","confidentiality_impact":"NONE","integrity_impact":"PARTIAL","availability_impact":"NONE","vulnerability_source":"http://nvd.nist.gov","assessment_check":null,"scanner":null,"recommendation":"","modified_at":"2008-09-05T16:46:23.540Z","published_at":"2005-03-14T00:00:00.000Z","created_at":"2017-02-10T23:51:32.371Z","updated_at":"2017-02-10T23:51:32.371Z","source_id":1},{"id":16192,"external_id":"CVE-2006-0614","title":"CVE-2006-0614","summary":"Unspecified vulnerability in Sun Java JDK and JRE 5.0 Update 3 and earlier, SDK and JRE 1.3.x through 1.3.1_16 and 1.4.x through 1.4.2_08 allows remote attackers to bypass Java sandbox security and obtain privileges via unspecified vectors involving the reflection APIs, aka the \"first issue.\"","score":"6.4","vector":"NETWORK","access_complexity":"LOW","vulnerability_authentication":"NONE","confidentiality_impact":"PARTIAL","integrity_impact":"PARTIAL","availability_impact":"NONE","vulnerability_source":"http://nvd.nist.gov","assessment_check":null,"scanner":null,"recommendation":"","modified_at":"2011-03-07T21:30:26.970Z","published_at":"2006-02-08T21:02:00.000Z","created_at":"2017-02-10T23:51:56.090Z","updated_at":"2017-02-10T23:51:56.090Z","source_id":1},{"id":16193,"external_id":"CVE-2006-0615","title":"CVE-2006-0615","summary":"Multiple unspecified vulnerabilities in Sun Java JDK and JRE 5.0 Update 4 and earlier, SDK and JRE 1.4.x through 1.4.2_09 allow remote attackers to bypass Java sandbox security and obtain privileges via unspecified vectors involving the reflection APIs, aka the \"second and third issues.\"","score":"4.0","vector":"NETWORK","access_complexity":"HIGH","vulnerability_authentication":"NONE","confidentiality_impact":"PARTIAL","integrity_impact":"PARTIAL","availability_impact":"NONE","vulnerability_source":"http://nvd.nist.gov","assessment_check":null,"scanner":null,"recommendation":"","modified_at":"2011-03-07T21:30:27.097Z","published_at":"2006-02-08T21:02:00.000Z","created_at":"2017-02-10T23:51:56.090Z","updated_at":"2017-02-10T23:51:56.090Z","source_id":1},{"id":16194,"external_id":"CVE-2006-0616","title":"CVE-2006-0616","summary":"Unspecified vulnerability in Sun Java JDK and JRE 5.0 Update 4 and earlier allows remote attackers to bypass Java sandbox security and obtain privileges via unspecified vectors involving the reflection APIs, aka the \"fourth issue.\"","score":"4.0","vector":"NETWORK","access_complexity":"HIGH","vulnerability_authentication":"NONE","confidentiality_impact":"PARTIAL","integrity_impact":"PARTIAL","availability_impact":"NONE","vulnerability_source":"http://nvd.nist.gov","assessment_check":null,"scanner":null,"recommendation":"","modified_at":"2011-03-07T21:30:27.187Z","published_at":"2006-02-08T21:02:00.000Z","created_at":"2017-02-10T23:51:56.090Z","updated_at":"2017-02-10T23:51:56.090Z","source_id":1},{"id":16195,"external_id":"CVE-2006-0617","title":"CVE-2006-0617","summary":"Multiple unspecified vulnerabilities in Sun Java JDK and JRE 5.0 Update 5 and earlier allow remote attackers to bypass Java sandbox security and obtain privileges via unspecified vectors involving the reflection APIs, aka the \"fifth, sixth, and seventh issues.\"","score":"4.0","vector":"NETWORK","access_complexity":"HIGH","vulnerability_authentication":"NONE","confidentiality_impact":"PARTIAL","integrity_impact":"PARTIAL","availability_impact":"NONE","vulnerability_source":"http://nvd.nist.gov","assessment_check":null,"scanner":null,"recommendation":"","modified_at":"2011-03-07T21:30:27.283Z","published_at":"2006-02-08T21:02:00.000Z","created_at":"2017-02-10T23:51:56.090Z","updated_at":"2017-02-10T23:51:56.090Z","source_id":1},{"id":17972,"external_id":"CVE-2006-2426","title":"CVE-2006-2426","summary":"Sun Java Runtime Environment (JRE) 1.5.0_6 and earlier, JDK 1.5.0_6 and earlier, and SDK 1.5.0_6 and earlier allows remote attackers to cause a denial of service (disk consumption) by using the Font.createFont function to create temporary files of arbitrary size in the %temp% directory.","score":"6.4","vector":"NETWORK","access_complexity":"LOW","vulnerability_authentication":"NONE","confidentiality_impact":"NONE","integrity_impact":"PARTIAL","availability_impact":"PARTIAL","vulnerability_source":"http://nvd.nist.gov","assessment_check":{"system":[],"href":[],"name":[]},"scanner":{"system":[],"href":[],"name":[]},"recommendation":"","modified_at":"2013-09-11T00:55:33.570Z","published_at":"2006-05-17T06:06:00.000Z","created_at":"2017-02-10T23:51:58.031Z","updated_at":"2017-02-10T23:51:58.031Z","source_id":1},{"id":20654,"external_id":"CVE-2006-5201","title":"CVE-2006-5201","summary":"Multiple packages on Sun Solaris, including (1) NSS; (2) Java JDK and JRE 5.0 Update 8 and earlier, SDK and JRE 1.4.x up to 1.4.2_12, and SDK and JRE 1.3.x up to 1.3.1_19; (3) JSSE 1.0.3_03 and earlier; (4) IPSec/IKE; (5) Secure Global Desktop; and (6) StarOffice, when using an RSA key with exponent 3, removes PKCS-1 padding before generating a hash, which allows remote attackers to forge a PKCS #1 v1.5 signature that is signed by that RSA key and prevents these products from correctly verifying X.509 and other certificates that use PKCS #1.","score":"4.0","vector":"NETWORK","access_complexity":"HIGH","vulnerability_authentication":"NONE","confidentiality_impact":"NONE","integrity_impact":"PARTIAL","availability_impact":"PARTIAL","vulnerability_source":"http://nvd.nist.gov","assessment_check":null,"scanner":null,"recommendation":"","modified_at":"2011-03-07T21:42:45.673Z","published_at":"2006-10-10T00:06:00.000Z","created_at":"2017-02-10T23:51:59.541Z","updated_at":"2017-02-10T23:51:59.541Z","source_id":1},{"id":21417,"external_id":"CVE-2006-6009","title":"CVE-2006-6009","summary":"Unspecified vulnerability in the Java Runtime Environment (JRE) Swing library in JDK and JRE 5.0 Update 7 and earlier allows attackers to obtain certain information via unknown attack vectors, related to an untrusted applet accessing data in other applets.","score":"5.0","vector":"NETWORK","access_complexity":"LOW","vulnerability_authentication":"NONE","confidentiality_impact":"PARTIAL","integrity_impact":"NONE","availability_impact":"NONE","vulnerability_source":"http://nvd.nist.gov","assessment_check":null,"scanner":null,"recommendation":"","modified_at":"2011-03-07T21:44:31.750Z","published_at":"2006-11-21T18:07:00.000Z","created_at":"2017-02-10T23:51:59.945Z","updated_at":"2017-02-10T23:51:59.945Z","source_id":1},{"id":22122,"external_id":"CVE-2006-6731","title":"CVE-2006-6731","summary":"Multiple buffer overflows in Sun Java Development Kit (JDK) and Java Runtime Environment (JRE) 5.0 Update 7 and earlier, Java System Development Kit (SDK) and JRE 1.4.2_12 and earlier 1.4.x versions, and SDK and JRE 1.3.1_18 and earlier allow attackers to develop Java applets that read, write, or execute local files, possibly related to (1) integer overflows in the Java_sun_awt_image_ImagingLib_convolveBI, awt_parseRaster, and awt_parseColorModel functions; (2) a stack overflow in the Java_sun_awt_image_ImagingLib_lookupByteRaster function; and (3) improper handling of certain negative values in the Java_sun_font_SunLayoutEngine_nativeLayout function.  NOTE: some of these details are obtained from third party information.","score":"9.3","vector":"NETWORK","access_complexity":"MEDIUM","vulnerability_authentication":"NONE","confidentiality_impact":"COMPLETE","integrity_impact":"COMPLETE","availability_impact":"COMPLETE","vulnerability_source":"http://nvd.nist.gov","assessment_check":{"system":[],"href":[],"name":[]},"scanner":{"system":[],"href":[],"name":[]},"recommendation":"","modified_at":"2011-03-07T21:46:52.610Z","published_at":"2006-12-26T18:28:00.000Z","created_at":"2017-02-10T23:52:00.256Z","updated_at":"2017-02-10T23:52:00.256Z","source_id":1},{"id":22127,"external_id":"CVE-2006-6736","title":"CVE-2006-6736","summary":"Unspecified vulnerability in Sun Java Development Kit (JDK) and Java Runtime Environment (JRE) 5.0 Update 6 and earlier, Java System Development Kit (SDK) and JRE 1.4.2_12 and earlier 1.4.x versions, and SDK and JRE 1.3.1_18 and earlier allows attackers to use untrusted applets to \"access data in other applets,\" aka \"The second issue.\"","score":"4.3","vector":"NETWORK","access_complexity":"MEDIUM","vulnerability_authentication":"NONE","confidentiality_impact":"PARTIAL","integrity_impact":"NONE","availability_impact":"NONE","vulnerability_source":"http://nvd.nist.gov","assessment_check":{"system":[],"href":[],"name":[]},"scanner":{"system":[],"href":[],"name":[]},"recommendation":"","modified_at":"2011-04-27T00:00:00.000Z","published_at":"2006-12-26T18:28:00.000Z","created_at":"2017-02-10T23:52:00.256Z","updated_at":"2017-02-10T23:52:00.256Z","source_id":1},{"id":22128,"external_id":"CVE-2006-6737","title":"CVE-2006-6737","summary":"Unspecified vulnerability in Sun Java Development Kit (JDK) and Java Runtime Environment (JRE) 5.0 Update 5 and earlier, Java System Development Kit (SDK) and JRE 1.4.2_10 and earlier 1.4.x versions, and SDK and JRE 1.3.1_18 and earlier allows attackers to use untrusted applets to \"access data in other applets,\" aka \"The first issue.\"","score":"4.3","vector":"NETWORK","access_complexity":"MEDIUM","vulnerability_authentication":"NONE","confidentiality_impact":"PARTIAL","integrity_impact":"NONE","availability_impact":"NONE","vulnerability_source":"http://nvd.nist.gov","assessment_check":{"system":[],"href":[],"name":[]},"scanner":{"system":[],"href":[],"name":[]},"recommendation":"","modified_at":"2011-03-07T21:46:53.673Z","published_at":"2006-12-26T18:28:00.000Z","created_at":"2017-02-10T23:52:00.256Z","updated_at":"2017-02-10T23:52:00.256Z","source_id":1},{"id":22875,"external_id":"CVE-2007-0243","title":"CVE-2007-0243","summary":"Buffer overflow in Sun JDK and Java Runtime Environment (JRE) 5.0 Update 9 and earlier, SDK and JRE 1.4.2_12 and earlier, and SDK and JRE 1.3.1_18 and earlier allows applets to gain privileges via a GIF image with a block with a 0 width field, which triggers memory corruption.","score":"6.8","vector":"NETWORK","access_complexity":"MEDIUM","vulnerability_authentication":"NONE","confidentiality_impact":"PARTIAL","integrity_impact":"PARTIAL","availability_impact":"PARTIAL","vulnerability_source":"http://nvd.nist.gov","assessment_check":{"system":[],"href":[],"name":[]},"scanner":{"system":[],"href":[],"name":[]},"recommendation":"","modified_at":"2011-03-07T21:49:04.110Z","published_at":"2007-01-17T17:28:00.000Z","created_at":"2017-02-10T23:52:20.971Z","updated_at":"2017-02-10T23:52:20.971Z","source_id":1},{"id":25356,"external_id":"CVE-2007-2788","title":"CVE-2007-2788","summary":"Integer overflow in the embedded ICC profile image parser in Sun Java Development Kit (JDK) before 1.5.0_11-b03 and 1.6.x before 1.6.0_01-b06, and Sun Java Runtime Environment in JDK and JRE 6, JDK and JRE 5.0 Update 10 and earlier, SDK and JRE 1.4.2_14 and earlier, and SDK and JRE 1.3.1_20 and earlier, allows remote attackers to execute arbitrary code or cause a denial of service (JVM crash) via a crafted JPEG or BMP file that triggers a buffer overflow.","score":"6.8","vector":"NETWORK","access_complexity":"MEDIUM","vulnerability_authentication":"NONE","confidentiality_impact":"PARTIAL","integrity_impact":"PARTIAL","availability_impact":"PARTIAL","vulnerability_source":"http://nvd.nist.gov","assessment_check":{"system":[],"href":[],"name":[]},"scanner":{"system":[],"href":[],"name":[]},"recommendation":"","modified_at":"2011-03-07T00:00:00.000Z","published_at":"2007-05-21T20:30:00.000Z","created_at":"2017-02-10T23:52:22.405Z","updated_at":"2017-02-10T23:52:22.405Z","source_id":1},{"id":25357,"external_id":"CVE-2007-2789","title":"CVE-2007-2789","summary":"The BMP image parser in Sun Java Development Kit (JDK) before 1.5.0_11-b03 and 1.6.x before 1.6.0_01-b06, and Sun Java Runtime Environment in JDK and JRE 6, JDK and JRE 5.0 Update 10 and earlier, SDK and JRE 1.4.2_14 and earlier, and SDK and JRE 1.3.1_19 and earlier, when running on Unix/Linux systems, allows remote attackers to cause a denial of service (JVM hang) via untrusted applets or applications that open arbitrary local files via a crafted BMP file, such as /dev/tty.","score":"4.3","vector":"NETWORK","access_complexity":"MEDIUM","vulnerability_authentication":"NONE","confidentiality_impact":"NONE","integrity_impact":"NONE","availability_impact":"PARTIAL","vulnerability_source":"http://nvd.nist.gov","assessment_check":{"system":[],"href":[],"name":[]},"scanner":{"system":[],"href":[],"name":[]},"recommendation":"","modified_at":"2011-03-07T00:00:00.000Z","published_at":"2007-05-21T20:30:00.000Z","created_at":"2017-02-10T23:52:22.405Z","updated_at":"2017-02-10T23:52:22.405Z","source_id":1},{"id":26058,"external_id":"CVE-2007-3503","title":"CVE-2007-3503","summary":"The Javadoc tool in Sun JDK 6 and JDK 5.0 Update 11 can generate HTML documentation pages that contain cross-site scripting (XSS) vulnerabilities, which allows remote attackers to inject arbitrary web script or HTML via unspecified vectors.","score":"4.3","vector":"NETWORK","access_complexity":"MEDIUM","vulnerability_authentication":"NONE","confidentiality_impact":"NONE","integrity_impact":"PARTIAL","availability_impact":"NONE","vulnerability_source":"http://nvd.nist.gov","assessment_check":{"system":[],"href":[],"name":[]},"scanner":{"system":[],"href":[],"name":[]},"recommendation":"","modified_at":"2012-10-30T22:38:56.030Z","published_at":"2007-06-29T21:30:00.000Z","created_at":"2017-02-10T23:52:22.878Z","updated_at":"2017-02-10T23:52:22.878Z","source_id":1},{"id":26059,"external_id":"CVE-2007-3504","title":"CVE-2007-3504","summary":"Directory traversal vulnerability in the PersistenceService in Sun Java Web Start in JDK and JRE 5.0 Update 11 and earlier, and Java Web Start in SDK and JRE 1.4.2_13 and earlier, for Windows allows remote attackers to perform unauthorized actions via an application that grants file overwrite privileges to itself.  NOTE: this can be leveraged to execute arbitrary code by overwriting a .java.policy file.","score":"9.3","vector":"NETWORK","access_complexity":"MEDIUM","vulnerability_authentication":"NONE","confidentiality_impact":"COMPLETE","integrity_impact":"COMPLETE","availability_impact":"COMPLETE","vulnerability_source":"http://nvd.nist.gov","assessment_check":null,"scanner":null,"recommendation":"","modified_at":"2012-10-30T22:38:57.357Z","published_at":"2007-06-29T21:30:00.000Z","created_at":"2017-02-10T23:52:22.878Z","updated_at":"2017-02-10T23:52:22.878Z","source_id":1},{"id":26251,"external_id":"CVE-2007-3698","title":"CVE-2007-3698","summary":"The Java Secure Socket Extension (JSSE) in Sun JDK and JRE 6 Update 1 and earlier, JDK and JRE 5.0 Updates 7 through 11, and SDK and JRE 1.4.2_11 through 1.4.2_14, when using JSSE for SSL/TLS support, allows remote attackers to cause a denial of service (CPU consumption) via certain SSL/TLS handshake requests.","score":"7.8","vector":"NETWORK","access_complexity":"LOW","vulnerability_authentication":"NONE","confidentiality_impact":"NONE","integrity_impact":"NONE","availability_impact":"COMPLETE","vulnerability_source":"http://nvd.nist.gov","assessment_check":{"system":[],"href":[],"name":[]},"scanner":{"system":[],"href":[],"name":[]},"recommendation":"","modified_at":"2012-10-30T22:39:36.937Z","published_at":"2007-07-11T18:30:00.000Z","created_at":"2017-02-10T23:52:23.029Z","updated_at":"2017-02-10T23:52:23.029Z","source_id":1},{"id":26462,"external_id":"CVE-2007-3922","title":"CVE-2007-3922","summary":"Unspecified vulnerability in the Java Runtime Environment (JRE) Applet Class Loader in Sun JDK and JRE 5.0 Update 11 and earlier, 6 through 6 Update 1, and SDK and JRE 1.4.2_14 and earlier, allows remote attackers to violate the security model for an applet's outbound connections by connecting to certain localhost services running on the machine that loaded the applet.","score":"6.8","vector":"NETWORK","access_complexity":"MEDIUM","vulnerability_authentication":"NONE","confidentiality_impact":"PARTIAL","integrity_impact":"PARTIAL","availability_impact":"PARTIAL","vulnerability_source":"http://nvd.nist.gov","assessment_check":{"system":[],"href":[],"name":[]},"scanner":{"system":[],"href":[],"name":[]},"recommendation":"","modified_at":"2011-03-07T21:57:22.970Z","published_at":"2007-07-20T20:30:00.000Z","created_at":"2017-02-10T23:52:23.183Z","updated_at":"2017-02-10T23:52:23.183Z","source_id":1}],"meta":{"copyright":"Copyright 2017 Selection Pressure LLC www.selectpress.net","authors":["Dan Hess","John Clemson"],"version":"v1","last_update":"2017-04-17T16:06:00.055Z","total_count":21,"limit":100,"offset":0},"links":{"self":"https://bunsen.ionchannel.testing/v1/vulnerability/getVulnerabilities?limit=100&offset=0&product=jdk"}}`
	SampleVulnerabilityResponse   = `{"data":{"id":59031,"external_id":"CVE-2013-4164","title":"CVE-2013-4164","summary":"Heap-based buffer overflow in Ruby 1.8, 1.9 before 1.9.3-p484, 2.0 before 2.0.0-p353, 2.1 before 2.1.0 preview2, and trunk before revision 43780 allows context-dependent attackers to cause a denial of service (segmentation fault) and possibly execute arbitrary code via a string that is converted to a floating point value, as demonstrated using (1) the to_f method or (2) JSON.parse.","score":"6.8","vector":"NETWORK","access_complexity":"MEDIUM","vulnerability_authentication":"NONE","confidentiality_impact":"PARTIAL","integrity_impact":"PARTIAL","availability_impact":"PARTIAL","vulnerability_source":"http://nvd.nist.gov","assessment_check":null,"scanner":null,"recommendation":"","references":[{"type":"PATCH","source":"CONFIRM","url":"https://www.ruby-lang.org/en/news/2013/11/22/ruby-2-0-0-p353-is-released","text":"https://www.ruby-lang.org/en/news/2013/11/22/ruby-2-0-0-p353-is-released"},{"type":"UNKNOWN","source":"CONFIRM","url":"https://www.ruby-lang.org/en/news/2013/11/22/ruby-1-9-3-p484-is-released","text":"https://www.ruby-lang.org/en/news/2013/11/22/ruby-1-9-3-p484-is-released"},{"type":"UNKNOWN","source":"CONFIRM","url":"https://www.ruby-lang.org/en/news/2013/11/22/heap-overflow-in-floating-point-parsing-cve-2013-4164","text":"https://www.ruby-lang.org/en/news/2013/11/22/heap-overflow-in-floating-point-parsing-cve-2013-4164"},{"type":"UNKNOWN","source":"CONFIRM","url":"https://support.apple.com/kb/HT6536","text":"https://support.apple.com/kb/HT6536"},{"type":"UNKNOWN","source":"DEBIAN","url":"http://www.debian.org/security/2013/dsa-2810","text":"DSA-2810"},{"type":"UNKNOWN","source":"DEBIAN","url":"http://www.debian.org/security/2013/dsa-2809","text":"DSA-2809"},{"type":"UNKNOWN","source":"REDHAT","url":"http://rhn.redhat.com/errata/RHSA-2014-0215.html","text":"RHSA-2014:0215"},{"type":"UNKNOWN","source":"REDHAT","url":"http://rhn.redhat.com/errata/RHSA-2014-0011.html","text":"RHSA-2014:0011"},{"type":"UNKNOWN","source":"REDHAT","url":"http://rhn.redhat.com/errata/RHSA-2013-1767.html","text":"RHSA-2013:1767"},{"type":"UNKNOWN","source":"REDHAT","url":"http://rhn.redhat.com/errata/RHSA-2013-1764.html","text":"RHSA-2013:1764"},{"type":"UNKNOWN","source":"REDHAT","url":"http://rhn.redhat.com/errata/RHSA-2013-1763.html","text":"RHSA-2013:1763"},{"type":"UNKNOWN","source":"OSVDB","url":"http://osvdb.org/100113","text":"100113"},{"type":"UNKNOWN","source":"SUSE","url":"http://lists.opensuse.org/opensuse-updates/2013-12/msg00028.html","text":"openSUSE-SU-2013:1835"},{"type":"UNKNOWN","source":"SUSE","url":"http://lists.opensuse.org/opensuse-updates/2013-12/msg00027.html","text":"openSUSE-SU-2013:1834"},{"type":"UNKNOWN","source":"APPLE","url":"http://archives.neohapsis.com/archives/bugtraq/2014-10/0103.html","text":"APPLE-SA-2014-10-16-3"},{"type":"UNKNOWN","source":"APPLE","url":"http://archives.neohapsis.com/archives/bugtraq/2014-04/0134.html","text":"APPLE-SA-2014-04-22-1"}],"modified_at":"2014-10-24T02:55:22.250Z","published_at":"2013-11-23T14:55:03.517Z","created_at":"2017-02-10T23:55:34.426Z","updated_at":"2017-02-10T23:55:34.426Z","source_id":1,"dependencies":[{"id":96317,"name":"ruby","org":"ruby-lang","version":"1.8","up":"","edition":"","aliases":null,"created_at":"2017-02-10T23:50:43.134Z","updated_at":"2017-02-10T23:50:43.134Z","title":"ruby-lang Ruby 1.8","references":[{"Version information":"ftp://ftp.ruby-lang.org/pub/ruby/1.8/"}],"part":"/a","language":"","source_id":1,"external_id":"cpe:/a:ruby-lang:ruby:1.8"},{"id":96345,"name":"ruby","org":"ruby-lang","version":"1.9","up":"","edition":"","aliases":null,"created_at":"2017-02-10T23:50:43.134Z","updated_at":"2017-02-10T23:50:43.134Z","title":"ruby-lang Ruby 1.9","references":[],"part":"/a","language":"","source_id":1,"external_id":"cpe:/a:ruby-lang:ruby:1.9"},{"id":96347,"name":"ruby","org":"ruby-lang","version":"1.9.2","up":"","edition":"","aliases":null,"created_at":"2017-02-10T23:50:43.134Z","updated_at":"2017-02-10T23:50:43.134Z","title":"ruby-lang Ruby 1.9.2","references":[],"part":"/a","language":"","source_id":1,"external_id":"cpe:/a:ruby-lang:ruby:1.9.2"},{"id":96348,"name":"ruby","org":"ruby-lang","version":"1.9.3","up":"","edition":"","aliases":null,"created_at":"2017-02-10T23:50:43.134Z","updated_at":"2017-02-10T23:50:43.134Z","title":"ruby-lang Ruby 1.9.3","references":[],"part":"/a","language":"","source_id":1,"external_id":"cpe:/a:ruby-lang:ruby:1.9.3"},{"id":96368,"name":"ruby","org":"ruby-lang","version":"2.1","up":"preview1","edition":"","aliases":null,"created_at":"2017-02-10T23:50:43.134Z","updated_at":"2017-02-10T23:50:43.134Z","title":"Ruby-lang Ruby 2.1preview1","references":[{"Version information ":"https://www.ruby-lang.org/en/news/2013/09/23/ruby-2-1-0-preview1-is-released/"}],"part":"/a","language":"","source_id":1,"external_id":"cpe:/a:ruby-lang:ruby:2.1:preview1"},{"id":96346,"name":"ruby","org":"ruby-lang","version":"1.9.1","up":"","edition":"","aliases":null,"created_at":"2017-02-10T23:50:43.134Z","updated_at":"2017-02-10T23:50:43.134Z","title":"ruby-lang Ruby 1.9.1","references":[],"part":"/a","language":"","source_id":1,"external_id":"cpe:/a:ruby-lang:ruby:1.9.1"},{"id":96359,"name":"ruby","org":"ruby-lang","version":"2.0.0","up":"","edition":"","aliases":null,"created_at":"2017-02-10T23:50:43.134Z","updated_at":"2017-02-10T23:50:43.134Z","title":"Ruby-lang Ruby 2.0.0","references":[],"part":"/a","language":"","source_id":1,"external_id":"cpe:/a:ruby-lang:ruby:2.0.0"}],"source":{"id":1,"name":"NVD","description":"National Vulnerability Database","created_at":"2017-02-08T01:14:34.835Z","updated_at":"2017-02-13T19:58:27.912Z","attribution":"Copyright © 1999–2017, The MITRE Corporation. CVE and the CVE logo are registered trademarks and CVE-Compatible is a trademark of The MITRE Corporation.","license":"Submissions: For all materials you submit to the Common Vulnerabilities and Exposures (CVE®), you hereby grant to The MITRE Corporation (MITRE) and all CVE Numbering Authorities (CNAs) a perpetual, worldwide, non-exclusive, no-charge, royalty-free, irrevocable copyright license to reproduce, prepare derivative works of, publicly display, publicly perform, sublicense, and distribute such materials and derivative works. Unless required by applicable law or agreed to in writing, you provide such materials on an \"AS IS\" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied, including, without limitation, any warranties or conditions of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A PARTICULAR PURPOSE.\n\nCVE Usage: MITRE hereby grants you a perpetual, worldwide, non-exclusive, no-charge, royalty-free, irrevocable copyright license to reproduce, prepare derivative works of, publicly display, publicly perform, sublicense, and distribute Common Vulnerabilities and Exposures (CVE®). Any copy you make for such purposes is authorized provided that you reproduce MITRE's copyright designation and this license in any such copy.\n","copyright_url":"http://cve.mitre.org/about/termsofuse.html"}},"meta":{"copyright":"Copyright 2017 Selection Pressure LLC www.selectpress.net","authors":["Dan Hess","John Clemson"],"version":"v1","last_update":"2017-04-18T16:06:17.245Z","limit":10,"offset":0},"links":{"self":"https://bunsen.ionchannel.testing/v1/vulnerability/getVulnerability?external_id=CVE%2d2013%2d4164"}}`
)
