import { processGeneralMetrics } from '@/utils/transformations';
import axios from 'axios';
import{ useEffect, useState } from 'react';


const BACKEND_URL =  process.env.NEXT_PUBLIC_BACKEND_URL 

console.log("BACKEND_URL", BACKEND_URL);

const axiosInstance = axios.create({
    baseURL: BACKEND_URL,
    withCredentials: true,
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

        let endpoint = role === "general" ? `/api/v1/metrics/general/${owner}/${repo}` : `/api/v1/metrics/${owner}/${repo}`;

        const fetchMetrics = async () => {
            console.log("Fetching metrics...");
            
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
                else{
                    endpoint += ""
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
    }, [owner, repo, role]); 

    

    if(role === "general"){
        console.log(data);
        
        const generalMetrics = processGeneralMetrics(data);
        return { generalMetrics, loading, error };
    }

    console.log("returning data:", data);
    

    return { data, loading, error };
};




