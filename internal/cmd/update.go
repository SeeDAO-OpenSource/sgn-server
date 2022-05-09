package cmd

import (
	"context"
	"errors"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"github.com/google/go-github/v43/github"
	"github.com/spf13/cobra"
	"github.com/waite-lee/sgn/pkg/mvc"
	"github.com/waite-lee/sgn/pkg/utils"
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

func NewUpdateCmd(hoptions *mvc.HttpClientOptions) *UpdateCmd {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update to latest version",
		Long:  "update to latest version",
		RunE:  func(cmd *cobra.Command, args []string) error { return update(cmd, hoptions, args) },
	}
	return (*UpdateCmd)(cmd)
}

func update(cmd *cobra.Command, hoptions *mvc.HttpClientOptions, args []string) error {
	client := github.NewClient(mvc.NewHttpClient(hoptions))
	currentVersion := getVersion(cmd)
	if currentVersion == "" {
		return errors.New("未找到版本信息")
	}
	latest, _, err := client.Repositories.GetLatestRelease(context.TODO(), "SeeDAO-OpenSource", "nft-server")
	if err != nil {
		return err
	}
	log.Println("latest tag: " + latest.GetTagName())
	log.Println("current version: " + currentVersion)
	if strings.Compare(currentVersion, latest.GetTagName()) < 0 {
		for _, asset := range latest.Assets {
			if strings.Contains(*asset.Name, "Linux") && strings.Contains(*asset.Name, "amd64") {
				reader, _, err := client.Repositories.DownloadReleaseAsset(context.TODO(), "SeeDAO-OpenSource", "nft-server", *asset.ID, client.Client())
				if err != nil {
					return err
				}
				defer reader.Close()
				filePath := "./temp_" + asset.GetName()
				file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
				_, err = io.Copy(file, reader)
				file.Close()
				err = utils.Uncompress(filePath, path.Dir(os.Args[0])+"/test")
				if err != nil {
					return err
				}
				err = os.Remove(filePath)
				if err != nil {
					return err
				}
				break
			}
		}
	}
	return nil
}

func getVersion(cmd *cobra.Command) string {
	root := cmd.Root()
	return root.Version
}
