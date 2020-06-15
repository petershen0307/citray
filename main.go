package main

import (
	"context"
	"fmt"
)

func main() {
	setting := parseSetting("setting.yaml")
	githubClient := githubClientWrapper{}
	githubClient.init(context.Background(), setting)
	fmt.Println(githubClient.queryPRFromSetting())
}
