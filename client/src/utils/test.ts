export const getRepoData = async (repo: string) => {
    const response = await fetch(`https://api.github.com/repos/${repo}`);
    const data = await response.json();
    return data;
}