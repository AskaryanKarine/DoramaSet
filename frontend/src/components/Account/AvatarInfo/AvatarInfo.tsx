import style from "./AvatarInvo.module.css"
import {useAppSelector} from "../../../hooks/redux";
import {EmojiStatus} from "../../EmojiStatus/EmojiStatus";

export function AvatarInfo() {
    const {user} = useAppSelector(state => state.userReducer)

    const colorStatus = user.sub && (user.sub.access_lvl >= 3) ? user.color : "#000000"
    const varIconStyle = user.sub && user.sub.access_lvl >= 3 ? `border-[${user.color}]` : "border-black"
    const allIconStyle = ["fa-solid fa-user fa-5x fa-border border-solid border-2 rounded-full", varIconStyle].join(" ")

    return (
        <div>
            <div className={style.iconName}>
                    <span style={{color: colorStatus}}>
                        <i className={allIconStyle}></i>
                    </span>
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