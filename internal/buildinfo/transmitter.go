package buildinfo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/error418/yoke/internal/prettylog"
	"github.com/error418/yoke/internal/swingletree"
	"github.com/go-resty/resty/v2"
)

func sendReport(apiBase string, report swingletree.Report, client *resty.Client) error {
	file, err := os.Open(report.Report)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	webhookPart := strings.Replace(report.Plugin, ".", "", -1)
	if splitIndex := strings.LastIndex(webhookPart, "/"); splitIndex >= 0 {
		webhookPart = webhookPart[splitIndex:]
	}
	webhookPart = strings.Replace(webhookPart, "/", "", -1)

	endpoint := fmt.Sprintf("%s/report/%s", apiBase, webhookPart)

	reportFile, err := os.Open(report.Report)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to open report file %s", report.Report))
	}
	reportData, err := ioutil.ReadAll(reportFile)
	if err != nil {
		return errors.New("Failed to read report data for upload.")
	}

	defer reportFile.Close()

	// set default value of content type
	contentType := report.ContentType
	if contentType == "" {
		contentType = "application/json"
	}

	prettylog.Info("Sending report data of %s to report endpoint %s", report.Report, endpoint)
	res, err := client.R().
		SetBody(reportData).
		SetResult(swingletree.Response{}).
		SetError(swingletree.Response{}).
		SetHeader("Content-Type", contentType).
		Post(endpoint)

	if err != nil {
		return err
	}

	if res.IsError() {
		reqErrors := res.Error().(*swingletree.Response).Errors
		return errors.New(fmt.Sprintf("Failed to send report. Caused by: %v", reqErrors))
	} else if res.IsSuccess() {
		prettylog.Check("'%v' successfully sent\n", reportFile.Name())
	}

	return err
}

func (info BuildInfo) Transmit(apiBase, token string, config swingletree.Config) (PublishReport, error) {
	client := resty.New()

	client.SetDisableWarn(true)

	uuid := info.BuildId()
	prettylog.Info("Publishing using uuid %s", uuid)

	client.
		SetBasicAuth("yoke", token).
		SetHeader("X-swingletree-origin", info.GitInfo.Remote.Url).
		SetHeader("X-swingletree-org", info.GitInfo.Organization).
		SetHeader("X-swingletree-repo", info.GitInfo.Repository).
		SetHeader("X-swingletree-sha", info.GitInfo.Sha).
		SetHeader("X-swingletree-branch", info.GitInfo.BranchName).
		SetHeader("X-swingletree-uid", uuid)

	var missing, failures int

	for _, report := range config.Yoke.Reports {
		if _, err := os.Stat(report.Report); os.IsNotExist(err) {
			prettylog.Warn("Missing file %s. Skipping report for plugin %s", report.Report, report.Plugin)
			missing++
		} else {
			err := sendReport(apiBase, report, client)
			if err != nil {
				prettylog.Fail("Failed to send report %s. Caused by: %v", report.Plugin, err)
				failures++
			}
		}
	}

	publishReport := PublishReport{}
	publishReport.Target = apiBase
	publishReport.Ok = len(config.Yoke.Reports) - missing - failures
	publishReport.Missing = missing
	publishReport.Failures = failures
	return publishReport, nil
}

type PublishReport struct {
	Target   string
	Ok       int
	Missing  int
	Failures int
}

func (p PublishReport) String() string {
	return fmt.Sprintf("  Uploaded to:   %s\n", p.Target) +
		fmt.Sprintf("  Ok:            %v\n", p.Ok) +
		fmt.Sprintf("  Missing:       %v\n", p.Missing) +
		fmt.Sprintf("  Failures:      %v\n", p.Failures)
}
