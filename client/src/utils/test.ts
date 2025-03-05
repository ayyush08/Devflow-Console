

export const getRepoData = async (repo: string) => {
    const response = await fetch(`https://api.github.com/repos/${repo}`);
    const data = await response.json();
    return data;
}

export const getRepoCommits = async (repo: string) => {
    const accessToken : string = process.env.NEXT_PUBLIC_GITHUB_ACCESS_TOKEN as string
    const maxCommits = 500
    const commits = [];
    let url: string | null = `https://api.github.com/repos/${repo}/commits?per_page=100`;

    try {
        while (url && commits.length < maxCommits) {
            const response = await fetch(url, {
                headers: {
                    Accept: "application/vnd.github+json",
                    Authorization: `Bearer ${accessToken}`,
                    
                },
            });

            if (!response.ok) {
                if (response.status === 403) {
                    throw new Error(
                        `Rate limit exceeded. Remaining: ${response.headers.get("X-RateLimit-Remaining")}, Reset: ${new Date(
                            Number(response.headers.get("X-RateLimit-Reset")) * 1000
                        )}`
                    );
                }
                throw new Error(`HTTP error! Status: ${response.status}`);
            }

            const data: any[] = await response.json();
            const newCommits = data.slice(0, maxCommits - commits.length);
            commits.push(...newCommits);

            const linkHeader = response.headers.get("Link");
            url = null;
            if (linkHeader) {
                const links = parseLinkHeader(linkHeader);
                if (links.next) {
                    url = links.next;
                }
            }
        }

        
        return commits;
    } catch (error) {
        console.error(`Error fetching commits for ${repo}:`, error);
        throw error;
    }
};

interface LinkHeader {
    [key: string]: string | undefined;
    next?: string;
    last?: string;
}
function parseLinkHeader(header: string | null): LinkHeader {

    if (!header) return {};

    const links: LinkHeader = {};
    const parts = header.split(",");

    parts.forEach((part) => {
        const [urlPart, relPart] = part.split(";").map((p) => p.trim());
        const urlMatch = urlPart.match(/<(.*)>/);
        const relMatch = relPart.match(/rel="(.*)"/);

        if (urlMatch && relMatch) {
            links[relMatch[1]] = urlMatch[1];
        }
    });

    return links;
}