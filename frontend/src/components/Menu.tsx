import React from 'react';
import './styles.css'


export function Menu() {
    return (
        <nav className="Nav">
            <div style={{marginTop:'30px'}}>Публичные списки</div>
            <div>Дорамы</div>
            <div>Стафф</div>
            <div>Мои списки</div>
            <div>Мое избранное</div>
        </nav>
    )
}