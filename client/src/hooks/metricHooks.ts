
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




export const useGetMetrics = ({ owner, repo, role }: ApiParams) => {
    const [data, setData] = useState<any>(null);
    const [loading, setLoading] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        if (!owner || !repo) return;

        let endpoint =  `/api/v1/metrics`;

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
                }else{
                    endpoint += "/general"
                }

                endpoint+=`/${owner}/${repo}`
                const response = await axiosInstance.get(endpoint);
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


    return { data, loading, error };
};
