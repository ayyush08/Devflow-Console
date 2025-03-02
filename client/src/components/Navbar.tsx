"use client";
import { APP_NAME } from "@/lib/constants";
import Link from "next/link";

import { isAuthenticated } from "@/lib/constants";
import StyledLink from "./StyledLink";
import { usePathname } from "next/navigation";

export default function Navbar() {
    
    const pathname = usePathname()
    
    return (
        <div className="w-full justify-between p-5 px-12 top-0 left-0 right-0 z-10 fixed flex">
            <Link className=" text-2xl font-bold " href='/'>{APP_NAME}</Link>
            <div className="flex space-x-3">
                {isAuthenticated ? <AuthButtons pathname={pathname} /> : <NonAuthButtons pathname={pathname} />}
            </div>
        </div>
    );
}


const NonAuthButtons: React.FC<{ pathname: string }> = ({ pathname }) => {
    return (
        <>
            {pathname !== '/dashboard/public' ? <StyledLink to="/dashboard/public" color="cyan" name="Explore Metrics" /> : null}

            <StyledLink to="/login" color="blue" name="Login" />
        </>
    );
}

const AuthButtons: React.FC<{ pathname: string }> = ({ pathname }) => {
    return (
        <>
            {pathname !== '/dashboard' ? <StyledLink to="/dashboard" color="cyan" name="Dashboard" /> : null}
            <StyledLink to="/logout" color="red" name="Logout" />
        </>
    );
}