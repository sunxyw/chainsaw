package jobs

import (
	"encoding/json"
	"gohub/app/models/news"
	"gohub/pkg/logger"
	"net/http"
)

type FetchMCVersionsAsNews struct {
}

type apiResponse struct {
	Versions []version `json:"versions"`
}

type version struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	Time        string `json:"time"`
	ReleaseTime string `json:"releaseTime"`
}

func (job *FetchMCVersionsAsNews) Name() string {
	return "fetch_mc_versions_as_news"
}

func (job *FetchMCVersionsAsNews) Run() {
	endpoint := "https://launchermeta.mojang.com/mc/game/version_manifest.json"
	resp, err := http.Get(endpoint)
	if err != nil {
		logger.LogIf(err)
		return
	}
	resp.Body.Close()

	var apiResponse apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		logger.LogIf(err)
		return
	}

	for _, version := range apiResponse.Versions {
		title := "Minecraft " + version.ID + " 现已发布！"
		news.AddNews(title, version.URL)
	}
}

func (job *FetchMCVersionsAsNews) ShouldRunAtStartup() bool {
	return true
}

func (job *FetchMCVersionsAsNews) CronSpec() string {
	return "@every 1h"
}
