import React from 'react';
import styles from './SideMenu.module.css'
import {useAppSelector} from "../../hooks/redux";
import {Link} from "react-router-dom";

export function SideMenu() {
    const {isAuth} = useAppSelector(state => state.userReducer)
    return (
        <div className={styles.menu}>
            <nav className={`${styles.navigation} ${styles.menu}`}>

                <div className={styles.firstDiv}>
                    <Link to="/list/public" className={styles.linkDiv}>
                        Публичные коллекции
                    </Link></div>

                <div><Link to="/dorama" className={styles.linkDiv}>
                        Дорамы
                </Link></div>

                <div><Link to="/staff" className={styles.linkDiv}>
                        Стафф
                </Link></div>

                {isAuth &&
                    <>
                        <div><Link to="/list" className={styles.linkDiv}>
                                Мои коллекции
                        </Link></div>

                        <div><Link to="/list/favorite" className={styles.linkDiv}>
                                Избранное
                        </Link></div>
                    </>
                }
            </nav>
        </div>
    )
}