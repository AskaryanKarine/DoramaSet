import style from "./Nickname.module.css"
import {useAppSelector} from "../../../hooks/redux";
import {EmojiStatus} from "../../EmojiStatus/EmojiStatus";
import {Avatar} from "../Avatar/Avatar";
import React from "react";
import {ErrorMessage} from "../../ErrorMessage/ErrorMessage";

export function Nickname() {
    const {user, error} = useAppSelector(state => state.userReducer)

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
                <p>{user.isAdmin ? "Администратор" : "Пользователь"}</p>
                <p>Баллы: {user.points}</p>
            </div>
            <ErrorMessage error={error as string}/>
        </div>
    )
}