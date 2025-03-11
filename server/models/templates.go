package models

type DeveloperTemplate struct {
	Merged       int     `json:"merged"`
	Pending      int     `json:"pending"`
	Closed       int     `json:"closed"`
	AvgMergeTime float64 `json:"avg_merge_time"`
}

type QATemplate struct {
	Rejected    int     `json:"rejected"`
	Conclusions map[TestConclusion]int `json:"conclusions"` 
	PassRate    float64 `json:"pass_rate"`
}

type ManagerTemplate struct {
	PRCount    int `json:"pr_count"`
	Stars      int `json:"stars"`
	ClosedPRs  int `json:"closed_prs"` 
	OpenIssues int `json:"open_issues"`
	ClosedIssues int `json:"closed_issues"`
}
