package models

type PRMetrics struct {
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



type DashboardMetrics struct {
    PRMetrics   PRMetrics   `json:"pr_metrics"`
    RepoMetrics RepoMetrics `json:"repo_metrics"`
    TestMetrics TestMetrics `json:"test_metrics"`
}


type TestStatus string

const (
	StatusQueued    TestStatus = "QUEUED"
	StatusInProgress TestStatus = "IN_PROGRESS"
	StatusCompleted  TestStatus = "COMPLETED"
)

type TestConclusion string

const (
	ConclusionSuccess TestConclusion = "SUCCESS"
	ConclusionFailure TestConclusion = "FAILURE"
	ConclusionNeutral TestConclusion = "NEUTRAL"
	ConclusionSkipped TestConclusion = "SKIPPED"
	ConclusionCanceled TestConclusion = "CANCELED"
	ConclusionTimedOut TestConclusion = "TIMED_OUT"
	ConclusionActionRequired TestConclusion = "ACTION_REQUIRED"
	ConclusionStartupFailure TestConclusion = "STARTUP_FAILURE"
)


type TestSuite struct {
	Conclusion TestConclusion `json:"conclusion"`
	Status     TestStatus     `json:"status"`
	Workflow   string         `json:"workflow"`
}


type TestMetrics struct {
	TotalTests   int         `json:"totalTests"`
	Statuses     map[TestStatus]int `json:"statuses"`
	Conclusions  map[TestConclusion]int `json:"conclusions"`
	TestSuites   []TestSuite  `json:"testSuites,omitempty"`
}
