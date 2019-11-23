package git

import (
	"fmt"
)

type person struct {
	Name string
	Mail string
}

type remote struct {
	Url string
}

type GitInfo struct {
	Remote       remote
	Sha          string
	BranchName   string
	Author       person
	Committer    person
	Organization string
	Repository   string
}

func (p person) String() string {
	return fmt.Sprintf("%s <%s>", p.Name, p.Mail)
}

func (info GitInfo) String() string {
	return fmt.Sprintf(
		"  Remote origin: %s\n"+
			"  Branch name:   %s\n"+
			"  Sha:           %s\n"+
			"  Remote origin: %s\n"+
			"  Organization:  %s\n"+
			"  Repository:    %s\n"+
			"  Committer:     %s\n"+
			"  Author:        %s",
		info.Remote.Url,
		info.BranchName,
		info.Sha,
		info.Remote.Url,
		info.Organization,
		info.Repository,
		info.Committer,
		info.Author,
	)
}
