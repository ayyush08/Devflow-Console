import {
    Card,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"




interface TileProps {
    title: string;
    value: number;
    icon: React.ReactNode;
}

const Tile = ({
    title,
    value,
    icon
}:TileProps) => {
    return (
        <Card className="min-w-64 px-2 bg-transparent shadow-md shadow-slate-500 ">
            <CardHeader>
                <div className="flex items-center justify-between text-white">
                <CardTitle className="text-2xl font-bold tracking-tight text-gray-200">{title}</CardTitle>
                    {icon}
                </div>
                <CardDescription className="text-4xl font-extrabold tracking-tight text-white">{value}</CardDescription>
            </CardHeader>
            </Card>

            )
}

            export default Tile