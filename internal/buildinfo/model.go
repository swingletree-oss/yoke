package buildinfo

import "github.com/error418/yoke/internal/git"

type BuildInfo struct {
	GitInfo *git.GitInfo
}

func NewBuildInfo() (BuildInfo, error) {
	p := BuildInfo{}

	gitInfo, err := git.CollectGitInfo()
	if err != nil {
		return BuildInfo{}, err
	}

	p.GitInfo = gitInfo

	return p, nil
}
