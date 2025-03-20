import axios from 'axios';
import { useEffect, useState } from 'react';


const BACKEND_URL =  process.env.NEXT_PUBLIC_BACKEND_URL 

console.log("BACKEND_URL", BACKEND_URL);

const axiosInstance = axios.create({
    baseURL: BACKEND_URL,
    withCredentials: true,
    headers:{
        'Content-Type': 'application/json',
    }
});


interface ApiParams {
    owner: string | null;
    repo: string | null;
    role?: "developer" | "qa" | "manager" | "general" ;
}

export const useGetMetrics = ({ owner, repo, role }: ApiParams) => {
    const [data, setData] = useState<any>(null);
    const [loading, setLoading] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        if (!owner || !repo) return;

        const fetchMetrics = async () => {
            console.log("Fetching metrics...");
            
            setLoading(true);
            setError(null);
            try {
                let endpoint = `/api/v1/metrics/${owner}/${repo}`;
                if (role === "developer") {
                    endpoint += "/developer";
                } else if (role === "qa") {
                    endpoint += "/qa";
                } else if (role === "manager") {
                    endpoint += "/manager";
                }
                else{
                    endpoint += "/"
                }

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
    }, [owner, repo]); 

    return { data, loading, error };
};