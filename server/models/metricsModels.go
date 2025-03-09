package models

type PRMetrics struct {
    Merged      int     `json:"merged"`
    Pending     int     `json:"pending"`
    Rejected    int     `json:"rejected"`
    AvgMergeTime float64 `json:"avg_merge_time"`
    PRCount     int     `json:"pr_count"`
    
}

type RepoMetrics struct {
    Stars      int    `json:"stars"`
    OpenIssues int    `json:"open_issues"`
    Forks      int    `json:"forks"`
    LastUpdated string `json:"last_updated"`
}

type TestMetrics struct {
    TestsPassed int     `json:"tests_passed"`
    TestsFailed int     `json:"tests_failed"`
    PassRate    float64 `json:"pass_rate"`
}

type DashboardMetrics struct {
    PRMetrics   PRMetrics   `json:"pr_metrics"`
    RepoMetrics RepoMetrics `json:"repo_metrics"`
    TestMetrics TestMetrics `json:"test_metrics"`
}