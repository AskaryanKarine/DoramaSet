import React, {useCallback, useEffect, useReducer, useRef} from 'react';
import './styles.css'
import {useAppSelector} from "../hooks/redux";

export function Menu() {
    const {token} = useAppSelector(state => state.userReducer)
    return (
        <div className="Menu">
            <nav className="Menu Nav">
                <div style={{marginTop:'30px'}}>Публичные списки</div>
                <div>Дорамы</div>
                <div>Стафф</div>
                {token.length > 0 && <>
                    <div>Мои списки</div>
                    <div>Мое избранное</div>
                </>}
            </nav>
        </div>
    )
}