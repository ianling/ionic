package util

import (
	"regexp"
	"strings"
)

// ParseGitURL takes a Git URL as a string in one of several forms and returns the repo URL along with the branch or
// commit hash, if one of them was present in the URL.
// The URL must resemble one of these:
//  * https://github.com/org/some-repo.git
//  * https://github.com/org/some-repo.git@branch-name
//  * git@github.com:org/some-repo.git
//  * git@github.com:org/some-repo.git@branch-name
// If the Git URL contains a commit hash instead of a branch name, the branch returned will be "HEAD"
func ParseGitURL(gitURL string) (string, string) {
	var foundBranchName bool
	var branch string
	branchDelimiterIndex := strings.LastIndex(gitURL, "@")
	if branchDelimiterIndex != -1 {
		// if there is a ':' after the last '@', we know we do not have a branch name,
		// because branch names cannot contain colons. A git URL with an '@' that does not denote a branch will
		// always also contain a colon.
		possibleBranch := gitURL[branchDelimiterIndex+1:]
		// this thing could be a branch name, or it could be a commit hash.
		// To determine which it is, check if it looks like a commit hash (exactly 40 lower-case hex characters)
		commitHashRegex := regexp.MustCompile("\\A[a-f0-9]{40}\\z")
		isCommitHash := commitHashRegex.MatchString(possibleBranch)
		if !strings.Contains(possibleBranch, ":") && !isCommitHash {
			// this is a branch name
			branch = possibleBranch
			foundBranchName = true
			// strip the branch name from the url
			gitURL = gitURL[0:branchDelimiterIndex]
		} else if isCommitHash {
			// strip the branch name from the url
			gitURL = gitURL[0:branchDelimiterIndex]
		}
	}
	if !foundBranchName {
		// use the remote's default branch
		branch = "HEAD"
	}

	return gitURL, branch
}
