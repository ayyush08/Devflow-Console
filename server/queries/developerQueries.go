package queries

const DeveloperMetricsQuery = `
query DeveloperMetrics($owner: String!, $repo: String!, $since: GitTimestamp!) {
  repository(owner: $owner, name: $repo) {
    # Total Commits
    defaultBranchRef {
      target {
        ... on Commit {
          history(since: $since) {
            totalCount
          }
        }
      }
    }

    # Pull Requests (Total, Open, Closed, Merged)
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

    # Commit history (for daily commits, additions, deletions)
    ref(qualifiedName: "refs/heads/main") {
      target {
        ... on Commit {
          history(first: 100, since: $since) {
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
