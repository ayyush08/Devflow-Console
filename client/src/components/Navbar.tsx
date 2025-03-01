"use client";
import { APP_NAME } from "@/lib/constants";
import Link from "next/link";


export default function Navbar() {
    return (
        <div className="w-full justify-between p-5 top-0 left-0 right-0 z-10 fixed flex">
            <Link className=" text-2xl font-bold " href='/'>{APP_NAME}</Link>
            <div className="flex space-x-3">
                <button className="bg-slate-300/10 hover:bg-slate-400/10 px-4 py-2 rounded-lg font-sans font-semibold">Login</button>
                <button className="bg-slate-300/10 hover:bg-slate-400/10 px-4 py-2 rounded-lg font-sans font-semibold">Login</button>
                <button className="bg-slate-300/10 hover:bg-slate-400/10 px-4 py-2 rounded-lg font-sans font-semibold">Login</button>

            </div>
        </div>
    );
}