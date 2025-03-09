package models



type DeveloperTemplate struct {
	Merged       int     `json:"merged"`
	Pending      int     `json:"pending"`
	AvgMergeTime float64 `json:"avg_merge_time"`
}

type QATemplate struct {
	Rejected    int     `json:"rejected"`
	TestsPassed int     `json:"tests_passed"`
	TestsFailed int     `json:"tests_failed"`
	PassRate    float64 `json:"pass_rate"`
}

type ManagerTemplate struct {
	PRCount    int `json:"pr_count"`
	Stars      int `json:"stars"`
	OpenIssues int `json:"open_issues"`
}
