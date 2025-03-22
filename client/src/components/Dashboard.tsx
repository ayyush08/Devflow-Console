import React, { useEffect, useState } from 'react'
import Tile from './Tile'
import {
    Bug,
    GitCommitIcon,
    GitPullRequest,
    Star
} from 'lucide-react'
import { AreaGraph } from './charts/AreaGraph'
import { DonutChart } from './charts/DonutChart'
import { BarGraph } from './charts/BarGraph'
import { Tabs, TabsList, TabsTrigger } from './ui/tabs'
import { Button } from './ui/button'
import ChartLoader from './loaders/ChartLoader'
import { useGetMetrics } from '@/hooks/metricHooks'

const sampleTileData = [
    { title: "Total Stars", value: 1000, icon: <Star /> },
    { title: "Total Commits", value: 500, icon: <GitCommitIcon /> },
    { title: "Total PRs", value: 350, icon: <GitPullRequest /> },
    { title: "Total Issues", value: 200, icon: <Bug /> },
]

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

type template = "general" | "developer" | "qa" | "manager";

const Dashboard = ({
    repoOwner,
    repoName,
    setRepo
}: DashboardProps) => {

    const [template, setTemplate] = React.useState<template>(tabs[0].value as template)
    const [metrics, setMetrics] = useState<any>(null);

    useEffect(() => {
        if (repoOwner && repoName) {
            console.log("Fetching metrics for:", repoOwner, repoName);
        }
    }, [setRepo]);

    const { data, loading, error } = useGetMetrics({
        owner: repoOwner || '',
        repo: repoName || '',
        role: template
    });

    
    
    useEffect(() => {
        console.log("Current template:", template);
        if (data) {
            setMetrics(data);
        }
        if(metrics) console.log("Fetched data:", metrics);
    }, [metrics]);


    if (error) {
        return <div>{error}</div>
    }

    if(loading){
        return <ChartLoader color='orange'  />
    }

    return (
        <div className="w-full">
            <Button variant={'secondary'} className='bg-slate-400/30 text-white hover:bg-slate-300 z-20 hover:text-black font-semibold px-4 absolute top-5 right-5' onClick={() => setRepo('')}>
                Reset
            </Button>
            <div className="bg-transparent text-white w-full h-full flex flex-col items-center justify-center p-14 z-10 relative">

                <h1 className="text-2xl md:text-5xl  font-extrabold mx-auto p-5 text-white">
                    {repoName} Dashboard

                </h1>
                <Tabs className='flex justify-center items-center w-full gap-6'>
                    <h2 className="text-xl md:text-2xl font-mono font-semibold my-4 text-center">
                        Choose a Dashboard Template
                    </h2>
                    <TabsList className='dark p-5 ' defaultValue={tabs[0].value} >
                        {tabs.map((tab, i) => (
                            <TabsTrigger
                                className={`text-md font-bold ${template === tab.value ? 'text-white bg-black' : 'text-gray-400'}`}
                                onClick={
                                    () => setTemplate(tab.value as template)
                                } key={i} value={tab.value}>{tab.name}</TabsTrigger>
                        ))}
                    </TabsList>
                </Tabs>
                <div className="flex  gap-4 mt-3 justify-center w-full items-center">
                    <div className="flex gap-3 p3">
                        {sampleTileData.map((tile, i) => (
                            <Tile key={i} title={tile.title} value={tile.value} icon={tile.icon} />
                        ))}
                    </div>
                </div>
                <div className="flex flex-col items-center gap-5 w-full overflow-hidden mt-2 flex-grow">
                    <AreaGraph />
                    <div className=" flex justify-center  min-w-full ">
                        <DonutChart />
                        <BarGraph />
                    </div>
                </div>
            </div>
        </div>
    )
}


export default Dashboard