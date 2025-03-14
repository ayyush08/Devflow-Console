"use client";
import {  useState } from "react";
import { PlaceholdersAndVanishInput } from "@/components/ui/placeholders-and-vanish-input";
import { BackgroundLines } from "@/components/ui/background-lines";
import CustomLineChart from "@/components/charts/CustomLineChart";

import ChartLoader from "@/components/loaders/ChartLoader";
import CustomPieChart from "@/components/charts/CustomPieChart";
import { useGetMetrics } from "@/hooks/metricHooks";



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

    if(data){
        console.log(data)
    }


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
                <div className="flex flex-col items-center justify-center w-full max-w-7xl mx-auto py-8">
                    {/* 🔥 Role Switcher */}
                    <div className="flex gap-4 mb-4">
                        {["developer", "qa", "manager"].map((r) => (
                            <button
                                key={r}
                                onClick={() => setRole(r as "developer" | "qa" | "manager")}
                                className={`px-4 py-2 border rounded ${role === r ? "bg-blue-500 text-white" : "bg-gray-300"}`}
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
                                <CustomLineChart
                                    // dataPoints={data?.lineDataPoints || []}
                                    labels={["data.labels"]}
                                    values={[2]}
                                    total={43}
                                    title="Commits Over Time"
                                    yAxisLabel="Commits"
                                    backgroundColor="purple"
                                    borderColor="cyan"
                                />
                            </div>
                            <div className="w-full h-96 md:h-[28rem] flex justify-center items-center">
                                <CustomPieChart
                                    dataPoints={[]}
                                    total={2}
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