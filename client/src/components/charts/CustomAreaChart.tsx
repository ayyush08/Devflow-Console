import { ResponsiveContainer, AreaChart, Area, XAxis, YAxis, Tooltip, Legend, CartesianGrid } from 'recharts';

interface CustomAreaChartProps {
    data: any[];
    xKey: string;
    yKeys: { key: string; color: string }[];
    strokeColor?: string;
    xAxisFill?: string;
    yAxisFill?: string;
    tooltipBackgroundColor?: string;
    tooltipColor?: string;
    
}

const CustomAreaChart: React.FC<CustomAreaChartProps> = ({ data, xKey, yKeys,strokeColor,xAxisFill,yAxisFill,tooltipBackgroundColor,tooltipColor }) => {
    return (
        <ResponsiveContainer width="100%" height={400}>
            <AreaChart data={data} margin={{ top: 20, right: 30, left: 20, bottom: 10 }}>
                <CartesianGrid strokeDasharray="2 3"  />
                <XAxis dataKey={xKey} tick={{ fill: xAxisFill}} />
                <YAxis tick={{ fill: yAxisFill }} />
                <Tooltip contentStyle={{ backgroundColor:tooltipBackgroundColor, color:tooltipColor }}  />
                <Legend wrapperStyle={{ color:"white" }} />
                {yKeys.map(({ key, color }) => (
                    <Area
                        key={key}
                        type="monotone"
                        dataKey={key}
                        stroke={strokeColor}
                        fill={color}
                        fillOpacity={0.6}
                    />
                ))}
            </AreaChart>
        </ResponsiveContainer>
    );
};

export default CustomAreaChart;
