interface Commit {
    commit: {
        committer: {
            date: string;
        };
        message: string;
    };
}
export const transformCommitData = (commits: Commit[]): {
    totalCommits: number;
    data: { label: string; value: number; details: string[] }[];
} => {
    const totalCommits = commits.length;
    const commitCountByDate = new Map<string, number>();
    const commitMessages = new Map<string, string[]>();
    commits.forEach((commit: any) => {
        const date = new Date(commit.commit.committer.date).toLocaleDateString();
        commitCountByDate.set(date, (commitCountByDate.get(date) || 0) + 1);
        const message = commit.commit.message;
        const existingMessages = commitMessages.get(date) || [];
        commitMessages.set(date, [...existingMessages, message]);
    });


    const transformedData = Array.from(commitCountByDate.entries()).map(([label, value]) => ({
        label,
        value,
        details: commitMessages.get(label) || [],
    })).sort((a, b) => {
        const dateA: Date = new Date(a.label.split('/').reverse().join('-')); // Convert to YYYY-MM-DD
        const dateB: Date = new Date(b.label.split('/').reverse().join('-'));
        return dateA.getTime() - dateB.getTime(); // Ascending order
    });


    return {
        totalCommits,
        data: transformedData,
    };

}

// utils/transformCommitByAuthor.ts
interface RawCommit {
    commit: { committer: { name: string } };
    author?: { login: string };
}

export const transformCommitByAuthor = (commits: RawCommit[]): {
    total: number;
    data: { label: string; value: number; details?: string[] }[];
} => {
    const total = commits.length;
    const commitsByAuthor = new Map<string, number>();

    commits.forEach((commit) => {
        const author = commit.author?.login || commit.commit.committer.name || "Unknown";
        commitsByAuthor.set(author, (commitsByAuthor.get(author) || 0) + 1);
    });

    const data = Array.from(commitsByAuthor.entries()).map(([label, value]) => ({
        label,
        value,
    }));

    return { total, data };
};