"use client";
import { useEffect, useState } from "react";
import { PlaceholdersAndVanishInput } from "@/components/ui/placeholders-and-vanish-input";
import { BackgroundLines } from "@/components/ui/background-lines";

import ChartLoader from "@/components/loaders/ChartLoader";
import { useGetMetrics } from "@/hooks/metricHooks";
import CustomAreaChart from "@/components/charts/CustomAreaChart";
import CustomLineChart from "@/components/charts/CustomLineChart";

const sample = [
    { time: "00:00", requests: 120, errors: 5 },
    { time: "01:00", requests: 150, errors: 7 },
    { time: "02:00", requests: 90, errors: 3 },
    { time: "03:00", requests: 170, errors: 6 },
    { time: "04:00", requests: 200, errors: 10 },
    { time: "05:00", requests: 230, errors: 8 },
    { time: "06:00", requests: 180, errors: 4 },
];

const sampleCommits = [
    { time: "00:00", commits: 5 },
    { time: "01:00", commits: 8 },
    { time: "02:00", commits: 3 },
    { time: "03:00", commits: 12 },
    { time: "04:00", commits: 7 },
    { time: "05:00", commits: 15 },
    { time: "06:00", commits: 10 },
];


type Role = "developer" | "qa" | "manager";

const placeholders = [
    "Enter user/repository name (e.g., facebook/react)"
];

export default function Dashboard() {
    const [repo, setRepo] = useState<string>('');
    const [role, setRole] = useState<"developer" | "qa" | "manager">("developer");
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        e.preventDefault();
    };
    const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget);
        const repoValue = formData.get("repo") as string;
        setRepo(repoValue);
    };

    const { data, loading, error } = useGetMetrics({
        owner: repo.split("/")[0] || "",
        repo: repo.split("/")[1] || "",
        role,
    });

    if (data) {
        console.log(data)
    }


    useEffect(() => {
        console.log("Role changed to", role);

    }, [role]);


    if (error) {
        return (
            <BackgroundLines className="flex justify-center items-center min-h-screen w-full flex-col px-4">
                <h2 className="text-lg md:text-2xl lg:text-6xl font-sans py-2 md:py-10 font-bold">
                    {error}
                </h2>
            </BackgroundLines>
        )
    }

    if (loading) {
        return (
            <BackgroundLines className="flex justify-center items-center min-h-screen w-full flex-col px-4">
                <ChartLoader color="rgba(75,192,192,1)" />
            </BackgroundLines>
        )
    }



    return (
        <BackgroundLines className="flex justify-center items-center min-h-screen w-full flex-col px-4">
            {repo === "" ? (
                <div className="flex flex-col justify-center items-center px-4">
                    <h2 className="text-lg md:text-2xl lg:text-6xl font-sans py-2 md:py-10 font-bold">
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
                <div className="flex flex-col z-50 items-center justify-center w-full max-w-7xl mx-auto py-8">
                    {/* 🔥 Role Switcher */}
                    <div className="flex gap-4 mb-4">
                        {["developer", "qa", "manager"].map((r) => (
                            <button
                                key={r}
                                onClick={() => setRole((prev) => (prev === r ? prev : (r as Role)))}
                                className={`px-4 py-2 border rounded ${role === r ? "bg-blue-500 text-white dark:text-black" : "bg-gray-600"}`}
                            >
                                {r.toUpperCase()}
                            </button>
                        ))}
                    </div>

                    {loading ? (
                        <ChartLoader color="rgba(75,192,192,1)" />
                    ) : (
                        <div className="grid grid-cols-1 md:grid-cols-2 gap-6 w-full">
                            <div className="w-full h-96 md:h-[28rem] flex justify-center items-center">
                                <CustomAreaChart
                                    data={sample}
                                    xKey="time"
                                    yKeys={[
                                        { key: "requests", color: "green" },
                                        { key: "errors", color: "#ff0000" },
                                    ]}
                                    xAxisFill="cyan"
                                    yAxisFill="cyan"
                                    strokeColor="white"
                                    tooltipBackgroundColor="black"
                                    tooltipColor="red"
                                />
                            </div>
                            <div className="w-full h-96 md:h-[28rem] flex justify-center items-center">
                                <CustomLineChart
                                    data={sampleCommits}
                                    title="Commits Over Time"
                                    yAxisLabel="Commits"
                                    tooltipBackgroundColor="black"
                                    tooltipColor="red"
                                    lineColor="cyan"
                                />
                            </div>
                        </div>
                    )}
                </div>
            )}
        </BackgroundLines>
    );
}