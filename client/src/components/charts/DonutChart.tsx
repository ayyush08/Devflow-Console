"use client"

import * as React from "react"
import { Label, Pie, PieChart } from "recharts"
import {
    Card,
    CardContent,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import {
    ChartContainer,
    ChartTooltip,
    ChartTooltipContent,
    ChartConfig,
} from "@/components/ui/chart"
import { colorPalette } from "@/lib/constants"



export interface DonutChartProps {
    chartData: Record<string, number>;  
    labelKey?: string; 
    valueKey?: string;
}

export function DonutChart({ chartData, labelKey = "Category", valueKey = "Value" }: DonutChartProps) {

    
    const formattedData = React.useMemo(() =>
        Object.entries(chartData).map(([key, value],index) => ({
            name: key[0].toUpperCase() + key.slice(1)+"",
            value,
            fill: colorPalette[index % colorPalette.length],
        })),
        [chartData]);
        console.log("Donut Chart Data: ", formattedData);
        
    let isEmpty = false
    isEmpty = formattedData.every((data) => data.value === 0);
    
    
    const totalValue = React.useMemo(() =>
        formattedData.reduce((acc, curr) => acc + curr.value, 0),
    [formattedData]);
    
    if (!chartData || chartData.length === 0 || !formattedData) return null;
    
    if(isEmpty) return null
    const config:ChartConfig = {}

    formattedData.forEach((item) => {
        config[item.name] = {
            label: item.name,
            color: item.fill,
        }
    })
    

    return (
        <Card className="flex flex-col bg-transparent text-white border-none w-full">
            <CardHeader className="items-center pb-0 text-xl">
                <CardTitle>{labelKey} Distribution </CardTitle>
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
                                                    className="fill-foreground text-3xl font-bold "
                                                    style={{ fill: "white" }}
                                                >
                                                    {totalValue.toLocaleString()}
                                                </tspan>
                                                <tspan
                                                    x={viewBox.cx}
                                                    y={(viewBox.cy || 0) + 24}
                                                    className="fill-foreground font-semibold"
                                                    style={{ fill: "white" }}
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
