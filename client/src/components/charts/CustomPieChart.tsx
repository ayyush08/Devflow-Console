"use client";
import React from "react";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { Pie } from "react-chartjs-2";

ChartJS.register(ArcElement, Tooltip, Legend);


interface ChartDataPoint {
    label: string;
    value: number;
    details?: string[];
}

interface CustomPieChartProps {
    dataPoints: ChartDataPoint[];
    total?: number;
    title?: string;
    colors?: string[];
}

const CustomPieChart: React.FC<CustomPieChartProps> = ({
    dataPoints,
    total,
    title = "Commit Distribution by Author",
    colors = [
        "#FF9500", // Neon Orange
        "#FF10F0", // Neon Pink
        "#00D4FF", // Neon Blue
        "#FFFF00", // Neon Yellow
        "#39FF14", // Neon Green
        "#D900FF", // Neon Purple
        "#F5F6FF",
    ],
}) => {

    const MAX_SLICES = 6;
    const topDataPoints = dataPoints
        .sort((a, b) => b.value - a.value)
        .slice(0, MAX_SLICES);
    const othersSum = dataPoints.slice(MAX_SLICES).reduce((sum, point) => sum + point.value, 0);
    const finalDataPoints =
        othersSum > 0 ? [...topDataPoints, { label: "Others", value: othersSum }] : topDataPoints;

    const data = {
        labels: finalDataPoints.map((point) => point.label),
        datasets: [
            {
                data: finalDataPoints.map((point) => point.value),
                backgroundColor: colors.slice(0, finalDataPoints.length),
                borderColor: "#000",
                borderWidth: 2,
            },
        ],
    };

    const options = {
        responsive: true,
        plugins: {
            legend: {
                position: "right" as const,
                labels: { color: "#fff", font: { size: 14 } },
            },
            tooltip: {
                callbacks: {
                    label: (context: any) => {
                        const label = context.label || "";
                        const value = context.raw || 0;
                        const percentage = total ? ((value / total) * 100).toFixed(1) : "N/A";
                        return `${label}: ${value} commits (${percentage}%)`;
                    },
                },
            },
        },
    };

    return (
        <div className="mx-auto w-full max-w-sm p-5 rounded-xl shadow-md ">
            <h2 className="text-xl font-bold text-center text-white">{title}</h2>
            <Pie data={data} options={options} />
        </div>
    );
};

export default CustomPieChart;