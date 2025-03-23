"use client"

import { Bar, BarChart, CartesianGrid, XAxis } from "recharts"

import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import {
    ChartConfig,
    ChartContainer,
    ChartTooltip,
    ChartTooltipContent,
} from "@/components/ui/chart"
import { useMemo } from "react";




export function BarGraph(
    { chartData }: { chartData: any[] }
) {

    const keys = useMemo(() => {
        return Object.keys(chartData?.[0] || {}).filter((key) => key !== "month");
    }, [chartData]);

    
    
    const chartConfig = useMemo(() => {
        return keys.reduce((config, key, index) => {
            config[key] = {
                label: key.charAt(0).toUpperCase() + key.slice(1),
                color: `hsl(var(--chart-${index + 1}))`,
            };
            return config;
        }, {} as ChartConfig);
    }, [keys]);
    
    if (!chartData || chartData.length === 0) return null;

    console.log("BarGraph data:", chartData);




    return (
        <Card className="text-white bg-transparent w-full border-none">
            <CardHeader className="text-center">
                <CardTitle>Bar Chart - Multiple</CardTitle>
                <CardDescription>January - June 2024</CardDescription>
            </CardHeader>
            <CardContent>
                <ChartContainer config={chartConfig} className="p-4">
                    <BarChart accessibilityLayer data={chartData}>
                        <CartesianGrid vertical={false} />
                        <XAxis
                            dataKey="month"
                            tickLine={false}
                            tickMargin={10}
                            axisLine={false}
                            tickFormatter={(value) => value.slice(0, 3)}
                        />
                        <ChartTooltip
                            cursor={false}
                            content={<ChartTooltipContent indicator="dot" className=" text-black" />}
                        />
                        <Bar dataKey="desktop" fill="red" radius={4} />
                        <Bar dataKey="mobile" fill="teal" radius={4} />
                    </BarChart>
                </ChartContainer>
            </CardContent>
        </Card>
    )
}
