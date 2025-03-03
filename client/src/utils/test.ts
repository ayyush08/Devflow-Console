import { transformCommitData } from "./transformations";

export const getRepoData = async (repo: string) => {
    const response = await fetch(`https://api.github.com/repos/${repo}`);
    const data = await response.json();
    return data;
}

export const getRepoCommits = async (repo: string) => {
    const response = await fetch(`https://api.github.com/repos/${repo}/commits`)
    const data = await response.json();
    const transformedData = transformCommitData(data);
    return transformedData;
}