




export const processGeneralMetrics = (data: any) => {


    return {
        tileData: {
            totalStars: data?.tileData.totalStars ?? 0,
            totalCommits: data?.tileData.totalCommits ?? 0,
            totalPRs: data?.tileData.totalPRs ?? 0,
            totalIssues: data?.tileData.totalIssues ?? 0,
        },
        areaGraphData: data?.areaGraphData ?? [],
        barGraphData: data?.barGraphData ?? [],
        donutChartData: {
            mergedPRs: data?.donutChartData?.mergedPRs ?? 0,
            closedPRs: data?.donutChartData?.closedPRs ?? 0,
            openPRs: data?.donutChartData?.openPRs ?? 0,
        },
    };
};


