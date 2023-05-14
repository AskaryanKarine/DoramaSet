import EmojiPicker, {Emoji, EmojiClickData} from "emoji-picker-react";
import style from "./AvatarInvo.module.css"
import {useAppSelector} from "../../../hooks/redux";
import {useState} from "react";

export function AvatarInfo() {
    const {user} = useAppSelector(state => state.userReducer)
    const [pickerVisibly, setPickerVisibly] = useState(false)
    const [chosenEmoji, setChosenEmoji] = useState({} as EmojiClickData);


    const varIconStyle = user.sub && user.sub.name === "premium" ? `border-[${user.color}]` : "border-black"
    const allIconStyle = ["fa-solid fa-user fa-5x fa-border border-solid border-2 rounded-full", varIconStyle].join(" ")
    const colorStatus = user.sub && user.sub.name === "premium" ? user.color : "#000000"

    return (
        <div className={style.container}>
            <div className={style.iconName}>
                    <span style={{color: colorStatus}}>
                        <i className={allIconStyle}></i>
                    </span>
                <h1>{user.username}</h1>
                {user.sub && user.sub.name === "premium" &&
                    <button
                        className="pl-3 pt-1.5 border-0 w-auto h-auto"
                        onClick={()=>{setPickerVisibly(!pickerVisibly)}}
                    >
                        <Emoji unified={user.emoji} size={50}/></button>}
            </div>
            <div className={style.info}>
                <p>Администратор</p>
                <p>Баллы: {user.points}</p>
            </div>
            <div className={style.dropdown}>
                {pickerVisibly && <EmojiPicker onEmojiClick={(emoji:EmojiClickData, event) => {setChosenEmoji(emoji)}}/>}
                {chosenEmoji && chosenEmoji.unified}
            </div>
        </div>
    )
}