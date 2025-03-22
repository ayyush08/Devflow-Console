export interface GeneralMetricsType {
    tileData: { title: string, value: number }[];
    areaGraphData: any;
    barGraphData: any;
    donutChartData: any;
    totalCommits: number;
    totalPRs: number;
    totalIssues: number;
    totalStars: number;
}

export interface TilePropsType {
    title: string;
    value: number;
    icon: React.ReactNode;
}
