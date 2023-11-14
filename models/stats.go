package models

type Stats struct {
	Path       string   `json:"path"`
	Count      int      `json:"count"`
	UniqueUser []string `json:"unique_user_agent"`
}

type StatsPublic struct {
	Path            string `json:"path"`
	Count           int    `json:"count"`
	UniqueUserCount int    `json:"unique_user_agent"`
}
