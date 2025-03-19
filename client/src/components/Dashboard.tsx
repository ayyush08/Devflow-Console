import React from 'react'
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

const sampleTileData = [
    { title: "Total Stars", value: 1000, icon: <Star /> },
    { title: "Total Commits", value: 500, icon: <GitCommitIcon /> },
    { title: "Total PRs", value: 350, icon: <GitPullRequest /> },
    { title: "Total Issues", value: 200, icon: <Bug /> },
]

interface DashboardProps {
    repo: string
}

const Dashboard = ({
    repo
}:DashboardProps) => {
    return (
        <div className="bg-transparent text-white w-full h-full flex flex-col items-center justify-center p-14 z-10 relative">
        
        <h1 className="text-2xl md:text-5xl  font-extrabold mx-auto p-5 text-white">
            {repo} Dashboard
        
        </h1>
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
    )
}


export default Dashboard