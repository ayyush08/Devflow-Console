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
    //Commits by date
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
    })).sort((a, b) => new Date(a.label).getTime() - new Date(b.label).getTime());

    return {
        totalCommits,
        data: transformedData,
    };

}