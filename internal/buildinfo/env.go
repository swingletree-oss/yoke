package buildinfo

import "fmt"

func (info BuildInfo) DotEnv() string {
	return fmt.Sprintf(
		"BUILD_ID=%s\n"     +
		"COMMIT_ID=%s\n"    +
		"BRANCH=%s\n"       +
		"ORGANIZATION=%s\n" +
		"REPOSITORY=%s\n",
		info.BuildId(),
		info.GitInfo.Sha,
		info.GitInfo.BranchName,
		info.GitInfo.Organization,
		info.GitInfo.Repository,
	)
}
