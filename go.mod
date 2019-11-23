module github.com/error418/yoke

require (
	github.com/go-resty/resty/v2 v2.1.0
	github.com/urfave/cli v1.22.1
	gopkg.in/src-d/go-git.v4 v4.13.1
	gopkg.in/yaml.v2 v2.2.4
)

replace github.com/error418/yoke/internal v0.0.0 => ./internal

go 1.13
