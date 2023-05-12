import {useAppSelector} from "../../hooks/redux";
import {Link} from "react-router-dom";
import React from "react";
import styles from './TopMenu.module.css'
import {Auth} from "../Auth/Auth";

export function TopMenu() {
    const {token, user} = useAppSelector(state => state.userReducer)

    return (
        <div className={styles.menu}>
            {token.length > 0 &&
                <Link to="/account">
                    <button className="w-[110px] h-[40px]">
                        <i className="fa-regular fa-user pr-1 fa-lg"></i>
                        {user.username}
                    </button>
                </Link>
            }
            <Auth/>
        </div>
    )
}