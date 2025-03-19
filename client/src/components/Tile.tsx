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
        <Card className="bg-transparent shadow-md shadow-slate-500 ">
            <CardHeader className="bg-transparent">
                <div className="flex gap-5  items-center justify-between text-white">
                <CardTitle className="text-md font-bold tracking-tight text-gray-200">{title}</CardTitle>
                <span>
                    {icon}
                </span>
                </div>
                <CardDescription className="text-2xl font-extrabold tracking-tight text-white">{value}</CardDescription>
            </CardHeader>
            </Card>

            )
}

            export default Tile