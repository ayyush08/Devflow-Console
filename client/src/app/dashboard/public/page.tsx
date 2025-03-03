"use client";

import { useEffect, useState } from "react";
import { PlaceholdersAndVanishInput } from "@/components/ui/placeholders-and-vanish-input";
import { BackgroundLines } from "@/components/ui/background-lines";

import CustomLineChart from "@/components/charts/CustomLineChart";
import { getRepoCommits } from "@/utils/test";
import ChartLoader from "@/components/loaders/ChartLoader";



const placeholders = [
    "Enter user/repository name (e.g., facebook/react)"
];

export default function Dashboard() {
    const [repo, setRepo] = useState<string>('');
    const [commits, setCommits] = useState<{ label: string; value: number; details: string[] }[]>([]);
    const [totalCommits, setTotalCommits] = useState<number>(0);
    const [loading, setLoading] = useState<boolean>(true);
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        e.preventDefault();
    };
    const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget);
        const repoValue = formData.get("repo") as string;
        setRepo(repoValue);
    };


    useEffect(() => {
        const fetchCommits = async () => {
            if (!repo) return;
            try {
                setLoading(true);
                const data = await getRepoCommits(repo);
                const { totalCommits, data: transformedData } = data;
                setTotalCommits(totalCommits);
                setCommits(transformedData);
                console.log("Fetched commits:", commits);
            } catch (error) {
                console.error("Error fetching commits:", error);
            } finally {
                setLoading(false);
            }
        };
        fetchCommits();
    }, [repo]);


    return (
        <BackgroundLines className="flex justify-center items-center  h-screen w-full flex-col px-4">

            {repo === '' ?
                <div className="flex flex-col justify-center items-center px-4">
                    <h2 className="bg-clip-text text-transparent text-center bg-gradient-to-b  dark:from-teal-700 dark:to-gray-200 text-lg md:text-2xl lg:text-6xl font-sans py-2 md:py-10 relative z-20 font-bold tracking-tight ">
                        Explore metrics for your repositories
                    </h2>
                    <PlaceholdersAndVanishInput
                        placeholders={placeholders}
                        onChange={handleChange}
                        onSubmit={onSubmit}
                        name="repo"
                    />
                </div>
                : <div className="flex flex-col items-center justify-center w-full">
                    {
                        loading ? <div className="mx-auto w-full  p-4 rounded-xl shadow-md">
                            <ChartLoader color="rgba(75,192,192,1)" />
                        </div> : 
                        <div className="scale-125 p-5 mx-auto  overflow-visible z-20">

                        <CustomLineChart
                            dataPoints={commits}
                            total={totalCommits}
                            title="Commits "
                            yAxisLabel="Commits"
                            borderColor="yellow"
                            backgroundColor="red"
                            />
                            </div>
                    }
                </div>}


        </BackgroundLines>
    );
}