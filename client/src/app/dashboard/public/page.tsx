"use client";
import { Bar, Pie } from "react-chartjs-2";
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    BarElement,
    ArcElement,
    Title,
    Tooltip,
    Legend,
} from "chart.js";
import { useEffect, useState } from "react";
import { PlaceholdersAndVanishInput } from "@/components/ui/placeholders-and-vanish-input";
import { BackgroundLines } from "@/components/ui/background-lines";
import { getRepoData } from "@/utils/test";
import { useDebounceCallback } from "usehooks-ts";

// Register Chart.js components
ChartJS.register(CategoryScale, LinearScale, BarElement, ArcElement, Title, Tooltip, Legend);

const placeholders = [
    "Enter user/repository name (e.g., facebook/react)"
];

export default function Dashboard() {
    const [inputValue, setInputValue] = useState<string>('');
    const [repo, setRepo] = useState<string>('');
    const [data, setData] = useState<any>(null);

    const debouncedInputChange = useDebounceCallback(setInputValue, 500);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        e.preventDefault(); // Prevent form submission on change
        console.log(e.target.value);
        debouncedInputChange(e.target.value);
    };
    const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget);
        const repo = formData.get("repo") as string;
        console.log(repo);
        setRepo(repo);
    };


    useEffect(() => {
        if (repo !== "") {
            getRepoData(repo).then((data) => {
                setData(data);
                console.log("Fetched data:", data);
            });
        }
    }, [repo]);

    // Bar chart data for PRs
    // const prChartData = {
    //     labels: ["Open", "Merged"],
    //     datasets: [
    //         {
    //             label: "Pull Requests",
    //             data: [sampleData.repositories[0].prs.open, sampleData.repositories[0].prs.merged],
    //             backgroundColor: ["rgba(255, 99, 132, 0.6)", "rgba(54, 162, 235, 0.6)"],
    //             borderColor: ["rgba(255, 99, 132, 1)", "rgba(54, 162, 235, 1)"],
    //             borderWidth: 1,
    //         },
    //     ],
    // };

    // const prChartOptions = {
    //     responsive: true,
    //     maintainAspectRatio: false,
    //     plugins: { legend: { position: "top" as const }, title: { display: true, text: "Pull Request Status" } },
    // };

    // Pie chart data for tests
    // const testChartData = {
    //     labels: ["Passed", "Failed"],
    //     datasets: [
    //         {
    //             label: "Test Outcomes",
    //             data: [sampleData.repositories[0].tests.passed, sampleData.repositories[0].tests.failed],
    //             backgroundColor: ["rgba(75, 192, 192, 0.6)", "rgba(255, 99, 132, 0.6)"],
    //             borderColor: ["rgba(75, 192, 192, 1)", "rgba(255, 99, 132, 1)"],
    //             borderWidth: 1,
    //         },
    //     ],
    // };

    // const testChartOptions = {
    //     responsive: true,
    //     maintainAspectRatio: false,
    //     plugins: { legend: { position: "top" as const }, title: { display: true, text: "Test Outcomes" } },
    // };

    return (
        <BackgroundLines className="flex items-center justify-center h-screen w-full flex-col px-4">

            <div className="h-fit flex flex-col justify-center  items-center px-4">
                <h2 className="bg-clip-text text-transparent text-center bg-gradient-to-b  dark:from-teal-700 dark:to-gray-200 text-lg md:text-2xl lg:text-6xl font-sans py-2 md:py-10 relative z-20 font-bold tracking-tight ">
                            Explore metrics for your repositories
                            </h2>
                {repo === '' ?
                    <PlaceholdersAndVanishInput
                        placeholders={placeholders}
                        onChange={handleChange}
                        onSubmit={onSubmit}
                        name="repo"
                    /> : null}
            </div>

            {
                repo !== '' ? (
                    <div className="flex flex-col items-center justify-center w-full">
                        {repo}
                    </div>) : null
            }


        </BackgroundLines>
    );
}