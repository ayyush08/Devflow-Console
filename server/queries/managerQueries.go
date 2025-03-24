package queries

const ManagerMetricsQuery = `
query ManagerMetrics($owner: String!, $repo: String!, $since: GitTimestamp!) {
  repository(owner: $owner, name: $repo) {
    # Tile Data
    totalCommits: defaultBranchRef {
      target {
        ... on Commit {
          history(since: $since) {
            totalCount
          }
        }
      }
    }
    totalPRs: pullRequests {
      totalCount
    }
    totalBugsOpen:  issues(states: OPEN, labels: ["bug"]) {
      totalCount
    }
    totalTestsRun: defaultBranchRef {
      target {
        ... on Commit {
          history(since: $since) {
            totalCount
          }
        }
      }
    }

    # Area Graph Data (Commits & Bugs Reported per Day)
    commitsHistory: defaultBranchRef {
      target {
        ... on Commit {
          history(since: $since, first: 100) {
            edges {
              node {
                committedDate
              }
            }
          }
        }
      }
    }
    bugsReportedHistory: issues(first: 100, orderBy: {field: CREATED_AT, direction: DESC}, labels: ["bug"]) {
      nodes {
        createdAt
      }
    }

    # Bar Graph Data (PRs Merged & Bugs Fixed per Month)
    prsMergedHistory: pullRequests(states: MERGED, first: 100) {
      nodes {
        mergedAt
      }
    }
    bugsFixedHistory: issues(states: CLOSED, first: 100, labels: ["bug"]) {
      nodes {
        closedAt
      }
    }

    # Donut Chart Data (Open/Merged PRs & Open/Resolved Bugs)
    openPRs: pullRequests(states: OPEN) {
      totalCount
    }
    mergedPRs: pullRequests(states: MERGED) {
      totalCount
    }
    openBugs: issues(states: OPEN, labels: ["bug"]) {
      totalCount
    }
    resolvedBugs: issues(states: CLOSED, labels: ["bug"]) {
      totalCount
    }
  }
}

`
