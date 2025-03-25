package queries

const GeneralMetricsQuery = `
query RepoStats($owner: String!, $repo: String!, $since: GitTimestamp!) {
    repository(owner: $owner, name: $repo) {
    # Tiles: Total Stars, Issues, PRs, Commits
    stargazerCount
    issues {
        totalCount
    }
    pullRequests(first: 100,orderBy: {field: CREATED_AT, direction: DESC}) {
        totalCount
        edges {
        node {
            createdAt
        }
        }
    }
	totalCommits: defaultBranchRef {
        target {
        ... on Commit {
            history {
            totalCount
            }
        }
        }
    }
    recentCommits: defaultBranchRef {
        target {
        ... on Commit {
            history(since: $since) {
            totalCount
            edges {
                node {
                committedDate
                }
            }
            }
        }
        }
    }

    # Donut Chart: PRs Status Breakdown
    mergedPRs: pullRequests(states: MERGED) {
        totalCount
    }
    closedPRs: pullRequests(states: CLOSED) {
        totalCount
    }
    openPRs: pullRequests(states: OPEN) {
        totalCount
    }

    # Bar Chart: Additions & Deletions
    barData: defaultBranchRef {
        target {
        ... on Commit {
            history(first: 100) {
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
}`
