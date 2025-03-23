import { Bug, GitCommitIcon, GitPullRequest, Star } from 'lucide-react'
import React from 'react'

const IconsConfig = {
    "Stars": <Star />,
    "Issues": <Bug/>,
    "PRs" : <GitPullRequest/>,
    "Commits": <GitCommitIcon/>

}

const Icons = ({ name }: { name: string }): React.ReactNode => {

    const type = name.split("total")[1].trim() as keyof typeof IconsConfig;

    const icon = IconsConfig[type];

    if(!icon){
        return null;
    }

    return (
        <>
        {
            icon
        }
        </>
    )
}

export default Icons