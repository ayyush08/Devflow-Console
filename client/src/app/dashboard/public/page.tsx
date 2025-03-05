"use client";

import { useEffect, useState } from "react";
import { PlaceholdersAndVanishInput } from "@/components/ui/placeholders-and-vanish-input";
import { BackgroundLines } from "@/components/ui/background-lines";

import CustomLineChart from "@/components/charts/CustomLineChart";
import { getRepoCommits } from "@/utils/test";
import ChartLoader from "@/components/loaders/ChartLoader";
import { transformCommitByAuthor, transformCommitData } from "@/utils/transformations";
import CustomPieChart from "@/components/charts/CustomPieChart";



const placeholders = [
    "Enter user/repository name (e.g., facebook/react)"
];

export default function Dashboard() {
    const [repo, setRepo] = useState<string>('');
    const [totalCommits, setTotalCommits] = useState<number>(0);
    const [loading, setLoading] = useState<boolean>(true);
    const [lineDataPoints, setLineDataPoints] = useState<{ label: string; value: number; details?: string[] }[]>([]);
    const [pieDataPoints, setPieDataPoints] = useState<{ label: string; value: number; details?: string[] }[]>([]);
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
                const rawData = await getRepoCommits(repo);
                const { totalCommits, data: lineTransformedData } = transformCommitData(rawData);
                setTotalCommits(totalCommits);
                setLineDataPoints(lineTransformedData)
                const { data: pieTransformedData } = transformCommitByAuthor(rawData);
                setPieDataPoints(pieTransformedData);
                console.log("Line Data:", lineTransformedData);
                console.log("Pie Data:", pieTransformedData);
            } catch (error) {
                console.error("Error fetching commits:", error);
            } finally {
                setLoading(false);
            }
        };
        fetchCommits();
    }, [repo]);


    return (
        <BackgroundLines className="flex justify-center items-center min-h-screen w-full flex-col px-4">
            {repo === "" ? (
                <div className="flex flex-col justify-center items-center px-4">
                    <h2 className="bg-clip-text text-transparent text-center bg-gradient-to-b dark:from-teal-700 dark:to-gray-200 text-lg md:text-2xl lg:text-6xl font-sans py-2 md:py-10 relative z-20 font-bold tracking-tight">
                        Explore metrics for your repositories
                    </h2>
                    <PlaceholdersAndVanishInput
                        placeholders={placeholders}
                        onChange={handleChange}
                        onSubmit={onSubmit}
                        name="repo"
                    />
                </div>
            ) : (
                <div className="flex flex-col items-center justify-center w-full max-w-7xl mx-auto py-8">
                    {loading ? (
                        <div className="flex justify-center items-center w-full h-96">
                            <ChartLoader color="rgba(75,192,192,1)" />
                        </div>
                    ) : (
                        <div className="grid grid-cols-1 md:grid-cols-2 gap-6 w-full">
                            <div className="w-full h-96 md:h-[28rem] flex justify-center z-10 items-center">
                                <CustomLineChart
                                    dataPoints={lineDataPoints}
                                    total={totalCommits}
                                    title="Commits Over Time"
                                    yAxisLabel="Commits"
                                    backgroundColor="purple"
                                    borderColor="cyan"
                                />
                            </div>
                            <div className="w-full h-96 md:h-[28rem] flex z-10 justify-center items-center">
                                <CustomPieChart
                                    dataPoints={pieDataPoints}
                                    total={totalCommits}
                                    title="Commits by Author"
                                />
                            </div>
                        </div>
                    )}
                </div>
            )}
        </BackgroundLines>

    );
}