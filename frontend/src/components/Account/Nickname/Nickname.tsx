import style from "./Nickname.module.css"
import {useAppSelector} from "../../../hooks/redux";
import {EmojiStatus} from "../../EmojiStatus/EmojiStatus";
import {Avatar} from "../Avatar/Avatar";
import React from "react";

export function Nickname() {
    const {user} = useAppSelector(state => state.userReducer)

    return (
        <div>
            <div className={style.iconName}>
                <Avatar/>
                <h1>{user.username}</h1>
                {user.sub && user.sub.access_lvl >= 2 &&
                    <EmojiStatus/>
                }
            </div>
            <div className={style.info}>
                <p>Администратор</p>
                <p>Баллы: {user.points}</p>
            </div>
        </div>
    )
}