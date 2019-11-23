package buildinfo

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func (info BuildInfo) BuildId() string {
	var buildId = os.Getenv("BUILD_ID")

	if buildId == "" {
		buildPath, _ := os.Getwd()
		uid := os.Getuid()
		buildId = fmt.Sprintf("%s:%d:%s:%s", buildPath, uid, info.GitInfo.Sha, info.GitInfo.Repository)
	}

	hash := sha256.Sum256([]byte(buildId))

	return fmt.Sprintf("%x", hash)
}
