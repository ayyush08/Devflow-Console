package queries

const DeveloperMetricsQuery = `
query DeveloperMetrics($owner: String!, $repo: String!, $sinceCommits: GitTimestamp!) {
  repository(owner: $owner, name: $repo) {
    # ✅ Total Commits (All time)
    defaultBranchRef {
      target {
        ... on Commit {
          history {
            totalCount
          }
        }
      }
    }

    # ✅ Pull Requests (All time, filter in Go for last 30 days)
    pullRequests(first: 100, orderBy: { field: CREATED_AT, direction: DESC }) {
      totalCount
      nodes {
        state # OPEN, CLOSED, MERGED
        createdAt
        reviewRequests(first: 10) {
          totalCount
        }
      }
    }

    # ✅ Commit history (For Bar Graph: Last 6 months)
    ref(qualifiedName: "refs/heads/main") {
      target {
        ... on Commit {
          history(first: 100, since: $sinceCommits) {
            edges {
              node {
                committedDate
                additions
                deletions
              }
            }
          }
        }
      }
    }
  }
}


`
