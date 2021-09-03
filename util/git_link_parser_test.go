package util

import (
	"testing"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestGitLinkParser(t *testing.T) {
	g := goblin.Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("ParseGitLink", func() {
		g.It("should return the correct source URL and branch for an SSH Git URL with branch name", func() {
			url := "git@github.com:ianling/some_repo.git@my-branch"
			url, branch := ParseGitURL(url)

			Expect(url).To(Equal("git@github.com:ianling/some_repo.git"))
			Expect(branch).To(Equal("my-branch"))
		})

		g.It("should return the correct source URL and branch for an HTTPS Git URL with branch name", func() {
			url := "https://github.com/ianling/some_repo.git@my-branch"
			url, branch := ParseGitURL(url)

			Expect(url).To(Equal("https://github.com/ianling/some_repo.git"))
			Expect(branch).To(Equal("my-branch"))
		})

		g.It("should return the correct source URL and branch for an SSH Git URL with commit hash", func() {
			url := "git@github.com:ianling/some_repo.git@aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
			url, branch := ParseGitURL(url)

			Expect(url).To(Equal("git@github.com:ianling/some_repo.git"))
			Expect(branch).To(Equal("HEAD"))
		})

		g.It("should return the correct source URL and branch for an HTTPS Git URL with commit hash", func() {
			url := "https://github.com/ianling/some_repo.git@aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
			url, branch := ParseGitURL(url)

			Expect(url).To(Equal("https://github.com/ianling/some_repo.git"))
			Expect(branch).To(Equal("HEAD"))
		})

		g.It("should return the correct source URL and branch for a plain SSH Git URL", func() {
			url := "git@github.com:ianling/some_repo.git"
			url, branch := ParseGitURL(url)

			Expect(url).To(Equal("git@github.com:ianling/some_repo.git"))
			Expect(branch).To(Equal("HEAD"))
		})

		g.It("should return the correct source URL and branch for a plain HTTPS Git URL", func() {
			url := "https://github.com/ianling/some_repo.git"
			url, branch := ParseGitURL(url)

			Expect(url).To(Equal("https://github.com/ianling/some_repo.git"))
			Expect(branch).To(Equal("HEAD"))
		})

		g.It("should return the correct source URL and branch when branch name contains a valid commit hash", func() {
			// 40 a's is a valid commit hash, but the presence of h's means it must be a branch name, not a commit hash
			url := "git@github.com:ianling/some_repo.git@aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaahhhhhhhhh"
			url, branch := ParseGitURL(url)

			Expect(url).To(Equal("git@github.com:ianling/some_repo.git"))
			Expect(branch).To(Equal("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaahhhhhhhhh"))
		})
	})
}
