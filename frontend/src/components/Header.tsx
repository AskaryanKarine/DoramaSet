import React from 'react';
import './styles.css'

interface HeaderProps {
    children: React.ReactNode
}

export function Header({children}:HeaderProps) {
    return (
        <div className="header">
            <h1 className='text-3xl font-serif font-bold mx-20'>DoramaSet</h1>
            {children}
        </div>
    )
}