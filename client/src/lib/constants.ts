export const APP_NAME: string = "DevFlow Console";

export const isAuthenticated: boolean = false;


export const sampleDashboardData: any = {
    "repositories": [
        {
            "name": "ayyush08/my-project",
            "stars": 150,
            "open_issues": 5,
            "prs": {
                "open": 3,
                "merged": 10,
                "avg_merge_time": "2.5 days"
            },
            "tests": {
                "passed": 45,
                "failed": 5,
                "pass_rate": "90%"
            }
        }
    ]
}