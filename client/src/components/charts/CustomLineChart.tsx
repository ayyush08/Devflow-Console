'use client';
import React, { useMemo } from "react";
import { Line } from "react-chartjs-2";
import {
    Chart as ChartJS,
    LineElement,
    Tooltip,
    Legend,
    CategoryScale,
    LinearScale,
    PointElement,
    Title,
} from "chart.js";

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

interface CustomLineChartProps {
    labels: string[];
    values: number[];
    title?: string;
    yAxisLabel?: string;
    borderColor?: string;
    backgroundColor?: string;
    total?: number;
}

const CustomLineChart: React.FC<CustomLineChartProps> = ({
    labels,
    values,
    title = "Data Over Time",
    yAxisLabel = "Value",
    borderColor = "#4bc0c0", // Neon cyan
    backgroundColor = "rgba(75,192,192,0.2)",
    total,
}) => {
    const data = useMemo(() => ({
        labels,
        datasets: [
            {
                label: title,
                data: values,
                borderColor,
                backgroundColor,
                tension: 0.4,
            },
        ],
    }), [labels, values, borderColor, backgroundColor, title]);

    const options = useMemo(() => ({
        responsive: true,
        plugins: {
            legend: { display: false },
            tooltip: { enabled: true },
            title: {
                display: !!total,
                text: total ? `Total ${title}: ${total}` : "",
                color: "#fff",
                font: { size: 16 },
                padding: { top: 10, bottom: 20 },
            },
        },
        scales: {
            x: { ticks: { color: "#fff" }, grid: { color: "rgba(255,255,255,0.2)" } },
            y: {
                ticks: { color: "#fff" },
                grid: { color: "rgba(255,255,255,0.2)" },
                title: { display: true, text: yAxisLabel, color: "#fff" },
            },
        },
    }), [total, title, yAxisLabel]);

    return <Line data={data} options={options} />;
};

export default CustomLineChart;
