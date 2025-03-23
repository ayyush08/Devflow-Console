export interface GeneralMetricsType {
    tileData: {
        totalPRs: number;
        totalCommits: number;
        totalIssues: number;
        totalStars: number;
    };
    areaGraphData:any[];
    barGraphData:any[];
    donutChartData: any
}

export interface TilePropsType {
    title: string;
    value: number;
    icon: React.ReactNode;
}
