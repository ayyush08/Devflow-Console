'use client';
import Dashboard from '@/components/Dashboard';
import { Button } from '@/components/ui/button';
import { PlaceholdersAndVanishInput } from '@/components/ui/placeholders-and-vanish-input';
import React, { useState } from 'react'


// const sampleTileData = [
//     { title: "Total Stars", value: 1000, icon: <StarIcon /> },
//     { title: "Total Commits", value: 500, icon: <GitCommitIcon /> },
//     { title: "Total PRs", value: 350, icon: <GitPullRequest /> },
//     { title: "Total Issues", value: 200, icon: <Bug /> },
// ]

const placeholders = [
    "Enter user/repository name (e.g., facebook/react)"
];

// type Role = "developer" | "qa" | "manager";

const MetricsDashboard = () => {
    const [repo, setRepo] = useState<string>('ayush');
    // const [role, setRole] = useState<Role>("developer");
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        e.preventDefault();
    };
    const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget);
        const repoValue = formData.get("repo") as string;
        setRepo(repoValue);
    };
    return (
        <div className='flex flex-col justify-center items-center px-4 min-h-screen'>
          
            {repo === "" ? (
                <div className="flex flex-col justify-center items-center px-4">
                    <h2 className="text-lg md:text-2xl lg:text-6xl font-sans bg-gradient-to-b text-transparent bg-clip-text from-orange-500 to-teal-400 py-2 md:py-10 font-bold">
                        Explore metrics for your repositories
                    </h2>
                    <PlaceholdersAndVanishInput
                        placeholders={placeholders}
                        onChange={handleChange}
                        onSubmit={onSubmit}
                        name="repo"

                    />
                </div>
            ) :
                (
                    <div className="w-full">
                        <Button variant={'secondary'} className='bg-slate-400/30 text-white hover:bg-slate-300 z-20 hover:text-black font-semibold px-4 absolute top-5 right-5' onClick={() => setRepo('')}>
                            Reset
                        </Button>
                        <Dashboard repo={repo} />
                    </div>
                )}
        </div>
    )
}

export default MetricsDashboard