"use client";

import { PieChart, Pie, Cell, Tooltip, Legend } from "recharts";

interface PieData {
    name: string;
    value: number;
}

interface CustomPieChartProps {
    dataPoints: PieData[];
    title?: string;
    tooltipBackgroundColor?: string;
    tooltipColor?: string;
}



interface CustomTooltipProps {
    active?: boolean;
    payload?: any;
    tooltipBackgroundColor?: string;
    tooltipColor?: string;
}

const COLORS = ["#8884d8", "#82ca9d", "#ffc658", "#ff7f50", "#ff69b4"];

export default function CustomPieChart({ dataPoints, title,tooltipBackgroundColor="black",tooltipColor="white" }: CustomPieChartProps) {
    return (
        <div className="w-full h-full flex flex-col items-center">
            <h2 className="text-lg font-semibold mb-2">{title}</h2>
            <PieChart width={400} height={400}>
                <Pie
                    data={dataPoints}
                    cx="50%"
                    cy="50%"
                    outerRadius={120}
                    fill="#8884d8"
                    dataKey="value"
                    label
                >
                    {dataPoints.map((_, index) => (
                        <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
                    ))}
                </Pie>
                <Tooltip content={<CustomTooltip
                tooltipBackgroundColor={tooltipBackgroundColor}
                tooltipColor={tooltipColor}
                />} />
                <Legend />
            </PieChart>
        </div>
    );
}

const CustomTooltip = ({ active, payload ,tooltipBackgroundColor,tooltipColor}: CustomTooltipProps) => {
    if (active && payload && payload.length) {
        const { name, value } = payload[0]; // Extract data including color

        return (
            <div
                className="p-2 rounded-md shadow-md"
                style={{
                    backgroundColor: tooltipBackgroundColor,
                    color: tooltipColor
                }}
            >
                <p className="font-bold">
                    {name}: {value}
                </p>
            </div>
        );
    }
    return null;
};
