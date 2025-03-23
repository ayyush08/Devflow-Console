"use client"

import * as React from "react"
import { Area, AreaChart, CartesianGrid, XAxis } from "recharts"

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
    ChartLegend,
    ChartLegendContent,
    ChartTooltip,
    ChartTooltipContent,
} from "@/components/ui/chart"
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select"






export function AreaGraph(
    { chartData }: { chartData: any[] }
) {

    const [timeRange, setTimeRange] = React.useState("30d")
    if (!chartData || chartData.length === 0) return null


    console.log("AreaGraph incoming chartData", chartData);
    const sampleItem = chartData[0] || {}
    const dateKey =
        Object.keys(sampleItem).find(
            (key) =>
                key.toLowerCase().includes("date") ||
                (typeof sampleItem[key] === "string" && !isNaN(Date.parse(sampleItem[key]))),
        ) || "date"

    const dataKeys = Object.keys(sampleItem).filter((key) => key !== dateKey)

    const sortedDates = [...chartData].sort((a, b) => new Date(b[dateKey]).getTime() - new Date(a[dateKey]).getTime())
    const referenceDate = sortedDates.length > 0 ? new Date(sortedDates[0][dateKey]) : new Date()


    const dynamicChartConfig: ChartConfig = {}
    dataKeys.forEach((key, index) => {
        dynamicChartConfig[key] = {
            label: key.charAt(0).toUpperCase() + key.slice(1),
            color: index === 0 ? "hsl(var(--chart-1))" : "hsl(var(--chart-2))",
        }
    })

    const filteredData = chartData.filter((item) => {
        const date = new Date(item[dateKey])
        let daysToSubtract = 90
        if (timeRange === "30d") {
            daysToSubtract = 30
        } else if (timeRange === "7d") {
            daysToSubtract = 7
        }
        const startDate = new Date(referenceDate)
        startDate.setDate(startDate.getDate() - daysToSubtract)
        return date >= startDate
    })


    console.log("AreaGraph chartConfig", dynamicChartConfig);




    console.log("AreaGraph filteredData", filteredData);


    return (
        <Card className="bg-transparent text-white w-full border-none ">
            <CardHeader className="flex   items-center gap-2 space-y-0 py-5 sm:flex-row">
                <div className="grid flex-1 gap-1 text-center sm:text-left">
                    <CardTitle>Area Chart - Interactive</CardTitle>
                    <CardDescription>
                        Showing data for the last 30 days
                    </CardDescription>
                </div>
                <Select value={timeRange} onValueChange={setTimeRange}>
                    <SelectTrigger
                        className="w-[160px] rounded-lg sm:ml-auto"
                        aria-label="Select a value"
                    >
                        <SelectValue className="bg-black text-white" placeholder="Last 3 months" />
                    </SelectTrigger>
                    <SelectContent className="rounded-xl bg-black text-white">
                        <SelectItem value="30d" className="rounded-lg">
                            Last 30 days
                        </SelectItem>
                        <SelectItem value="7d" className="rounded-lg">
                            Last 7 days
                        </SelectItem>
                    </SelectContent>
                </Select>
            </CardHeader>
            <CardContent className="">
                <ChartContainer
                    config={dynamicChartConfig}
                    className="aspect-auto h-[250px] w-full"
                >
                    <AreaChart data={filteredData}>
                        <defs>
                            {dataKeys.map((key, index) => (
                                <linearGradient key={`fill-${index}`} id={`fill-${key}`} x1="0" y1="0" x2="0" y2="1">
                                    <stop offset="5%" stopColor={`var(--color-${key})`} stopOpacity={0.8} />
                                    <stop offset="95%" stopColor={`var(--color-${key})`} stopOpacity={0.1} />
                                </linearGradient>
                            ))}
                        </defs>
                        <CartesianGrid vertical={false} />
                        <XAxis
                            dataKey={dateKey}
                            tickLine={false}
                            axisLine={false}
                            tickMargin={8}
                            minTickGap={32}
                            tickFormatter={(value) => {
                                const date = new Date(value)
                                return date.toLocaleDateString("en-US", {
                                    month: "short",
                                    day: "numeric",
                                })
                            }}
                        />
                        <ChartTooltip
                            cursor={false}
                            content={
                                <ChartTooltipContent
                                    labelFormatter={(value) => {
                                        return new Date(value).toLocaleDateString("en-US", {
                                            month: "short",
                                            day: "numeric",
                                        })
                                    }}
                                    indicator="dot"
                                    className="text-black bg-white"
                                />
                            }
                        />
                        {dataKeys.map((key, index) => (
                            <Area
                                key={index}
                                dataKey={key}
                                type="natural"
                                fill={`url(#fill-${key})`}
                                stroke={`var(--color-${key})`}
                                stackId="a"
                            />
                        ))}
                        <ChartLegend content={<ChartLegendContent />} />
                    </AreaChart>
                </ChartContainer>
            </CardContent>
        </Card>
    )
}
