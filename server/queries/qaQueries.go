package queries

const QaMetricsQuery = `query QaMetrics($owner: String!, $repo: String!) {
  repository(owner: $owner, name: $repo) {

    # Tile Data
    totalBugsReported: issues(labels: "bug") {
      totalCount
    }
    totalBugsResolved: issues(labels: "bug", states: CLOSED) {
      totalCount
    }
    totalTestSuites: discussions {
      totalCount
    }

    # Area Graph Data (Bugs Reported Over Time)
    issues(first: 100, labels: "bug", orderBy: {field: CREATED_AT, direction: DESC}) {
      nodes {
        createdAt
      }
    }

    # Bar Graph Data (Bugs Fixed Over Time)
    closedIssues: issues(first: 100, labels: "bug", states: CLOSED, orderBy: {field: UPDATED_AT, direction: DESC}) {
      nodes {
        closedAt
      }
    }

    # Donut Chart Data (Success, Failed, Skipped Tests from GitHub CI/CD)
    defaultBranchRef {
      target {
        ... on Commit {
    checkSuites(first: 100) {
            nodes {
              conclusion # Success, Failure, or Skipped
            }
          }
        }
      }
    }
  }
}

`
