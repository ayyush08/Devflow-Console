import React, { useEffect, useState } from 'react'
import Tile from './Tile'

import { AreaGraph } from './charts/AreaGraph'
import { DonutChart } from './charts/DonutChart'
import { BarGraph } from './charts/BarGraph'
import { Tabs, TabsList, TabsTrigger } from './ui/tabs'
import { Button } from './ui/button'
import ChartLoader from './loaders/ChartLoader'
import { useGetMetrics } from '@/hooks/metricHooks'
import Icons from './Icons'


const tabs = [
    { name: "General", value: "general" },
    { name: "Developer", value: "developer" },
    { name: "QA Engineer", value: "qa" },
    { name: "Manager", value: "manager" }
]

interface DashboardProps {
    repoOwner: string | null;
    repoName: string | null;
    setRepo: (repo: string) => void;
}

type TemplateType = "general" | "developer" | "qa" | "manager";

const Dashboard = ({ repoOwner, repoName, setRepo }: DashboardProps) => {
    const [template, setTemplate] = useState<TemplateType>(tabs[0].value as TemplateType)
    const [metrics, setMetrics] = useState<any>(null);

    useEffect(() => {
        if (repoOwner && repoName) {
            console.log("Fetching metrics for:", repoOwner, repoName);
        }
    }, [repoOwner, repoName]);

    const {  data, loading, error } = useGetMetrics({
        owner: repoOwner || '',
        repo: repoName || '',
        role: template
    });

    useEffect(() => {
        console.log("Current template:", template);
        if (data) {
            setMetrics(data);
        }
    }, [template, data]);

    if (error) {
        return <div className="text-red-500">{error}</div>
    }

    if (loading) {
        return <ChartLoader color='orange' />
    }

    return (
        <div className="w-full">
            <Button
                variant={'secondary'}
                className='bg-slate-400/30 text-white hover:bg-slate-300 z-20 hover:text-black font-semibold px-4 absolute top-5 right-5'
                onClick={() => setRepo('')}
            >
                Reset
            </Button>
            <div className="bg-transparent text-white w-full h-full flex flex-col items-center justify-center p-14 z-10 relative">
                <h1 className="text-2xl md:text-5xl font-extrabold mx-auto p-5 text-white">
                    {repoName} Dashboard
                </h1>
                <Tabs className='flex justify-center items-center w-full gap-6'>
                    <h2 className="text-xl md:text-2xl font-mono font-semibold my-4 text-center">
                        Choose a Dashboard Template
                    </h2>
                    <TabsList className='dark p-5' defaultValue={tabs[0].value}>
                        {tabs.map((tab, i) => (
                            <TabsTrigger
                                key={i}
                                value={tab.value}
                                className={`text-md font-bold ${template === tab.value ? 'text-white bg-black' : 'text-gray-400'}`}
                                onClick={() => setTemplate(tab.value as TemplateType)}
                            >
                                {tab.name}
                            </TabsTrigger>
                        ))}
                    </TabsList>
                </Tabs>
                <div className="flex gap-4 mt-3 justify-center w-full items-center">
                    {metrics?.tileData && (
                        <div className="flex gap-3 p3">
                            {Object.entries(metrics.tileData).map(([key, value], i) => {
                                
                                return (<Tile
                                    key={i}
                                    title={key}
                                    value={value as number}
                                    icon={<Icons name={key} />}
                                />)
                            }
                            )}
                        </div>
                    )}
                </div>
                <div className="flex flex-col items-center gap-5 w-full overflow-hidden mt-2 flex-grow">
                    <AreaGraph chartData={metrics?.areaGraphData || []} />
                    <div className="flex justify-center min-w-full">
                        <DonutChart chartData={metrics?.donutChartData || {}} />
                        <BarGraph chartData={metrics?.barGraphData || []} />
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Dashboard;
