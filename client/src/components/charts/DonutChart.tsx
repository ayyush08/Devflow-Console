"use client"

import * as React from "react"
import { Label, Pie, PieChart } from "recharts"
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import {
    ChartContainer,
    ChartTooltip,
    ChartTooltipContent,
    ChartConfig,
} from "@/components/ui/chart"

// Define the expected chart data format
export interface DonutChartProps {
    chartData: Record<string, number>;  // Generic key-value pair data
    labelKey?: string; // Custom label for category (default: "Category")
    valueKey?: string; // Custom label for value (default: "Value")
}

export function DonutChart({ chartData, labelKey = "Category", valueKey = "Value" }: DonutChartProps) {
    // Convert object data to array format
    const formattedData = React.useMemo(() =>
        Object.entries(chartData).map(([key, value]) => ({
            name: key,
            value
        })),
        [chartData]);

    // Calculate total value for the center label
    const totalValue = React.useMemo(() =>
        formattedData.reduce((acc, curr) => acc + curr.value, 0),
        [formattedData]);

    // Generate config dynamically
    const config = React.useMemo(() => {
        return formattedData.reduce((acc, curr) => {
            acc[curr.name.toLowerCase()] = {
                label: curr.name,
                color: `#${Math.floor(Math.random() * 16777215).toString(16)}`, // Random color
            };
            return acc;
        }, {} as ChartConfig);
    }, [formattedData]);

    return (
        <Card className="flex flex-col bg-transparent text-white border-none w-full">
            <CardHeader className="items-center pb-0">
                <CardTitle>{labelKey} Distribution</CardTitle>
                <CardDescription>Data Representation</CardDescription>
            </CardHeader>
            <CardContent className="flex-1 pb-0">
                <ChartContainer config={config} className="mx-auto aspect-square max-h-[400px]">
                    <PieChart>
                        <ChartTooltip cursor={false} content={<ChartTooltipContent hideLabel />} />
                        <Pie
                            data={formattedData}
                            dataKey="value"
                            nameKey="name"
                            innerRadius={90}
                            strokeWidth={5}
                        >
                            <Label
                                content={({ viewBox }) => {
                                    if (viewBox && "cx" in viewBox && "cy" in viewBox) {
                                        return (
                                            <text
                                                x={viewBox.cx}
                                                y={viewBox.cy}
                                                textAnchor="middle"
                                                dominantBaseline="middle"
                                                style={{ fill: "white" }}
                                            >
                                                <tspan
                                                    x={viewBox.cx}
                                                    y={viewBox.cy}
                                                    className="fill-foreground text-3xl font-bold"
                                                    style={{ fill: "white" }}
                                                >
                                                    {totalValue.toLocaleString()}
                                                </tspan>
                                                <tspan
                                                    x={viewBox.cx}
                                                    y={(viewBox.cy || 0) + 24}
                                                    className="fill-muted-foreground"
                                                >
                                                    {valueKey}
                                                </tspan>
                                            </text>
                                        )
                                    }
                                }}
                            />
                        </Pie>
                    </PieChart>
                </ChartContainer>
            </CardContent>
        </Card>
    )
}
