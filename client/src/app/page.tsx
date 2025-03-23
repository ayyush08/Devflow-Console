'use client';
import Dashboard from '@/components/Dashboard';

import { PlaceholdersAndVanishInput } from '@/components/ui/placeholders-and-vanish-input';

import React, { useState } from 'react'



const placeholders = [
    "Enter user/repository name (e.g., facebook/react)"
];


const MetricsDashboard = () => {
    const [repo, setRepo] = useState<string>("");

    const [repoOwner, setRepoOwner] = useState<string | null>(null);
    const [repoName, setRepoName] = useState<string | null>(null);
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        e.preventDefault();
    };
    const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget);
        const repoValue = formData.get("repo") as string;
        if (repoValue.includes('/')) {
            const [owner, name] = repoValue.split('/');
            setRepoOwner(owner);
            setRepoName(name);
            setRepo(repoValue);
        } else {
            alert("Invalid repository format. Use owner/repo.");
        }
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

                    <Dashboard repoName={repoName} repoOwner={repoOwner} setRepo={setRepo} />
                )
            }
        </div>
    )
}

export default MetricsDashboard