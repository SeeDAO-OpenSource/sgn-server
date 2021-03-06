package main

import "github.com/SeeDAO-OpenSource/sgn/pkg/app"

var (
	version string
	date    string
	commit  string
	builtBy string
)

func GetVersion() *app.VersionInfo {
	return &app.VersionInfo{
		Version: version,
		Date:    date,
		Commit:  commit,
		BuiltBy: builtBy,
	}
}
