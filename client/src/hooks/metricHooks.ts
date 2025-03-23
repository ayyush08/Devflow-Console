import { processGeneralMetrics } from '@/utils/transformations';
import { GeneralMetricsType } from '@/utils/type';
import axios from 'axios';
import { useEffect, useState } from 'react';

const BACKEND_URL = process.env.NEXT_PUBLIC_BACKEND_URL;

const axiosInstance = axios.create({
    baseURL: BACKEND_URL,
    withCredentials: true,
});

interface ApiParams {
    owner: string | null;
    repo: string | null;
    role?: "developer" | "qa" | "manager" | "general";
}



// Define response type for non-general metrics (adjust as needed)
type MetricsData = GeneralMetricsType | Record<string, any>;

export const useGetMetrics = ({ owner, repo, role }: ApiParams) => {
    const [data, setData] = useState<MetricsData | null>(null);
    const [loading, setLoading] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        if (!owner || !repo) return;

        let endpoint = role === "general" ? `/api/v1/metrics/general/${owner}/${repo}` : `/api/v1/metrics/${owner}/${repo}`;

        const fetchMetrics = async () => {
            setLoading(true);
            setError(null);
            try {
                if (role === "developer") {
                    endpoint += "/developer";
                } else if (role === "qa") {
                    endpoint += "/qa";
                } else if (role === "manager") {
                    endpoint += "/manager";
                }

                const response = await axiosInstance.get<MetricsData>(endpoint);
                setData(response.data);
            } catch (err) {
                setError("Failed to fetch data");
                console.error(err);
            } finally {
                setLoading(false);
            }
        };

        fetchMetrics();
    }, [owner, repo, role]);

    if (role === "general" && data) {
        const generalMetrics: GeneralMetricsType = processGeneralMetrics(data as GeneralMetricsType);
        return { generalMetrics, loading, error };
    }

    return { data, loading, error };
};
