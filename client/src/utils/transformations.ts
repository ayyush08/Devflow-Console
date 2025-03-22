




export const processGeneralMetrics = (data: any) => {

    
    const tileData= [
        { title: "Total Stars", value: data?.totalStars},
        { title: "Total Commits", value: data?.totalCommits,  },
        { title: "Total PRs", value: data?.totalPRs },
        { title: "Total Issues", value: data?.totalIssues},
    ]
    

    return {
        tileData,
        areaGraphData: data?.areaGraphData,
        barGraphData: data?.barGraphData,
        donutChartData: data?.donutChartData,
    };
}
