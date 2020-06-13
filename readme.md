# Swingletree Yoke

Publish reports to Swingletree without pain.

## Description

Yoke enables you to integrate your CI-pipeline with ease.

* Publish build reports for processing to Swingletree

## Usage

This section covers yoke usage.

### Publishing build reports

Print the publish help of your yoke version with `yoke publish -h`

Yoke should be invoked after your build has generated the reports you want to upload to Swingletree.
Reports intended for upload are listed in the repository `.swingletree.yml`. A report needs to be processable by a
Swingletree Plugin running in your Swingletree installation.

```yaml
yoke:
  reports:
    - plugin: twistlock
      contenttype: application/json
      report: build/twistlock/report.json
    - plugin: zap
      contenttype: application/json
      report: build/zap/report.json
```

`plugin` specifies the Swingletree Plugin to send the report to. `contenttype` specifies the content type of the report. Defaults to `application/json`.
`report` points to the report file.

Yoke needs a token to authenticate. Set it using the environment variable `YOKE_TOKEN`.
