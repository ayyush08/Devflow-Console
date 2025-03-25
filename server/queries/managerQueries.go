package queries

const ManagerMetricsQuery = `
query ManagerMetrics($owner: String!, $repo: String!) {
  repository(owner: $owner, name: $repo) {
    # Tile Data
    totalCommits: defaultBranchRef {
      target {
        ... on Commit {
          history {
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
    totalIssuesOpen: issues(states: OPEN) {
    totalCount
  }

    # Area Graph Data (Commits & Bugs Reported per Day)
    commitsHistory: defaultBranchRef {
      target {
        ... on Commit {
          history( first: 100) {
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
    prsMergedHistory: pullRequests(states: MERGED, first: 100,orderBy: {field: UPDATED_AT,direction: DESC}) {
      nodes {
        mergedAt
      }
    }
    bugsFixedHistory: issues(states: CLOSED, first: 100, labels: ["bug"],orderBy: {field: UPDATED_AT,direction: DESC}) {
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
