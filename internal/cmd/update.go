package cmd

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-github/v43/github"
	"github.com/spf13/cobra"
)

type UpdateCmd cobra.Command

type Release struct {
	Id          string
	TagName     string `json:"tag_name"`
	Prerelease  bool   `json:"prerelease"`
	CreatedAt   string `json:"created_at"`
	PublishedAt string `json:"published_at"`
}

const (
	githubURL         = "https://api.github.com"
	releasesLatestURL = "/repos/SeeDAO-OpenSource/nft-server/releases/latest"
)

func NewUpdateCmd(client *http.Client) *UpdateCmd {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update to latest version",
		Long:  "update to latest version",
		RunE:  func(cmd *cobra.Command, args []string) error { return update(client) },
	}
	return (*UpdateCmd)(cmd)
}

func update(httpClient *http.Client) error {
	client := github.NewClient(httpClient)
	latest, _, err := client.Repositories.GetLatestRelease(context.TODO(), "SeeDAO-OpenSource", "nft-server")
	if err != nil {
		return err
	}
	for _, asset := range latest.Assets {
		if strings.Contains(*asset.Name, "linux") && strings.Contains(*asset.Name, "amd64") {
			reader, _, err := client.Repositories.DownloadReleaseAsset(context.TODO(), "SeeDAO-OpenSource", "nft-server", *asset.ID, client.Client())
			if err != nil {
				return err
			}
			defer reader.Close()
			file, err := os.OpenFile("", os.O_RDONLY|os.O_CREATE, 0666)
			io.Copy(file, reader)
		}
	}
	return nil
}
