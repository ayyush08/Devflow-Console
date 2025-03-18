import Link from 'next/link'
import React from 'react'



interface StyledLinkProps {
    to: string,
    color: string,
    name: string
}

const StyledLink = ({to, color,name}:StyledLinkProps) => {
    return (
        <Link className={`bg-slate-300/10 text-white shadow-inner shadow-${color}-300  hover:shadow-${color}-400 transition-all  hover:bg-slate-400/10 px-4 py-2 rounded-lg font-sans font-semibold`} href={to}>{name}</Link>
    )
}

export default StyledLink