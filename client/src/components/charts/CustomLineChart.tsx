import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from "recharts";

interface CustomLineChartProps {
    data: { time: string; commits: number }[];
    title: string;
    yAxisLabel: string;
    lineColor?: string;
    tooltipColor?: string;
    tooltipBackgroundColor?: string;

}

const CustomLineChart: React.FC<CustomLineChartProps> = ({
    data,
    title,
    yAxisLabel,
    lineColor = "#8884d8",
    tooltipBackgroundColor = "#333",
    tooltipColor = "#fff"   
}) => {
    return (
        <div className="w-full h-full p-4 rounded-lg shadow-lg  ">
            <h2 className="text-lg font-bold text-center mb-2">{title}</h2>
            <ResponsiveContainer width="100%" height="90%">
                <LineChart data={data}>
                    <CartesianGrid strokeDasharray="3 3" strokeOpacity={0.3} />
                    <XAxis dataKey="time" stroke="#ccc" />
                    <YAxis label={{ value: yAxisLabel, angle: -90, position: "insideLeft" }} stroke="#ccc" />
                    <Tooltip contentStyle={{ backgroundColor: tooltipBackgroundColor, color: tooltipColor, borderRadius: "5px" }} />
                    <Line type="natural" dataKey="commits" stroke={lineColor} strokeWidth={2} dot={{ r: 4 }} />
                </LineChart>
            </ResponsiveContainer>
        </div>
    );
};

export default CustomLineChart;
