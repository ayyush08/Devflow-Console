package models


type PRMetrics struct {
	Merged  int `json:"merged"`
	Pending int `json:"pending"`
	Rejected int `json:"rejected"`
}


type PRDetails struct {
	TopContributors []Contributor `json:"top_contributors"`
	PRCounts   int `json:"pr_counts"`
}


type Contributor struct {
	Name         string `json:"name"`
	Contributions int   `json:"contributions"`
}


type Changes struct {
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
}


type ReviewMetrics struct {
	Approved        int `json:"approved"`
	ChangesRequested int `json:"changes_requested"`
	Comments        int `json:"comments"`
}
