package models

type PRMetrics struct {
    Merged      int     `json:"merged"`
    Pending     int     `json:"pending"`
    Rejected    int     `json:"rejected"`
    AvgMergeTime float64 `json:"avg_merge_time"`
    OpenPRCount     int     `json:"open_pr_count"`
    ClosedPRCount     int     `json:"closed_pr_count"`
    MergedPRCount     int     `json:"merged_pr_count"`
    
}

type RepoMetrics struct {
    Stars      int    `json:"stars"`
    OpenIssues int    `json:"open_issues"`
    ClosedIssues int    `json:"closed_issues"`
    Forks      int    `json:"forks"`
    LastUpdated string `json:"last_updated"`
    CreatedAt string `json:"created_at"`
}

type TestMetrics struct {
    TestsPassed  int     `json:"tests_passed"`
    TestsFailed  int     `json:"tests_failed"`
    TestsSkipped int     `json:"tests_skipped"`
    TestsQueued  int     `json:"tests_queued"`
    TestsRunning int     `json:"tests_running"`
    TestsCanceled int    `json:"tests_canceled"`
    PassRate     float64 `json:"pass_rate"`
}


type DashboardMetrics struct {
    PRMetrics   PRMetrics   `json:"pr_metrics"`
    RepoMetrics RepoMetrics `json:"repo_metrics"`
    TestMetrics TestMetrics `json:"test_metrics"`
}