package git

import (
	"regexp"

	"gopkg.in/src-d/go-git.v4"
)

var originRegex = regexp.MustCompile("(http[s]?:\\/\\/|git@)[^\\/:]+[:\\/](.+)\\/(.+)(\\.git)")

func CollectGitInfo() (*GitInfo, error) {
	var info = GitInfo{}

	// Open git repository
	repo, err := git.PlainOpen(".")
	if err != nil {
		return nil, err
	}

	// Retrieve HEAD information
	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	// Collect remote configuration
	remoteOrigin, err := repo.Remote("origin")
	if err != nil {
		return nil, err
	}

	// Retrieve commit info
	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	remoteConfig := remoteOrigin.Config()
	info.Remote = remote{Url: remoteConfig.URLs[0]}

	originSplit := originRegex.FindStringSubmatch(info.Remote.Url)
	info.Organization = originSplit[2]
	info.Repository = originSplit[3]

	info.Sha = ref.Hash().String()
	info.BranchName = ref.Name().Short()

	info.Author.Name = commit.Author.Name
	info.Author.Mail = commit.Author.Email

	info.Committer.Name = commit.Committer.Name
	info.Committer.Mail = commit.Committer.Email

	return &info, nil
}
