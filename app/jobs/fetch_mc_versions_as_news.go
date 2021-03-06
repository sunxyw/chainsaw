package jobs

import (
	"encoding/json"
	"gohub/app/models/news"
	"gohub/pkg/logger"
	"net/http"
	"time"
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
	defer resp.Body.Close()

	var apiResponse apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		logger.LogIf(err)
		return
	}

	for _, version := range apiResponse.Versions {
		title := "Minecraft " + version.ID + " 现已发布！"
		releaseTime, err := time.Parse("2006-01-02T15:04:05-07:00", version.ReleaseTime)
		if err != nil {
			logger.LogIf(err)
			continue
		}

		// 舍弃过旧的版本
		if releaseTime.Before(time.Now().Add(-24 * time.Hour)) {
			continue
		}

		news.AddNews(title, version.URL, releaseTime)
	}
}

func (job *FetchMCVersionsAsNews) ShouldRunAtStartup() bool {
	return true
}

func (job *FetchMCVersionsAsNews) CronSpec() string {
	return "@every 1h"
}
