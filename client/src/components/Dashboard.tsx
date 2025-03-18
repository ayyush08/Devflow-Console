import React from 'react'
import Tile from './Tile'
import {
    Bug,
    GitCommitIcon,
    GitPullRequest
} from 'lucide-react'
import { AreaGraph } from './charts/AreaGraph'
import { DonutChart } from './charts/DonutChart'
import { BarGraph } from './charts/BarGraph'

const sampleTileData = [
    { title: "Total Requests", value: 1000, icon: <GitPullRequest /> },
    { title: "Total Commits", value: 500, icon: <GitCommitIcon /> },
    { title: "Total PRs", value: 350, icon: <GitPullRequest /> },
    { title: "Total Issues", value: 200, icon: <Bug /> },
]
const Dashboard = () => {
    return (
        <div className=' bg-black text-white  w-full p-5 '>
            <h1 className='text-4xl font-bold text-white px-4 my-5 mx-2'>
                Dashboard
            </h1>
            <div className="flex flex-wrap gap-6 justify-around items-center">
                {sampleTileData.map((tile, i) => (
                    <Tile key={i} title={tile.title} value={tile.value} icon={tile.icon} />
                ))}
                <div className="flex flex-col px-6 gap-5 w-full  mt-6">
                    <AreaGraph />
                    <div className="px-4 flex gap-6">
                    <DonutChart  />
                    <BarGraph />
                    </div>
                </div>
            </div>
        </div>
    )
}


export default Dashboard