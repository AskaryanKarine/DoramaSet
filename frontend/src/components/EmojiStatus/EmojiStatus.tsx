import {useAppDispatch, useAppSelector} from "../../hooks/redux";
import EmojiPicker, {Emoji, EmojiClickData} from "emoji-picker-react";
import React, {useState} from "react";
import style from "./EmojiStatus.module.css";
import classNames from "classnames";
import {setEmoji} from "../../store/reducers/UserSlice";

export function EmojiStatus() {
    const [pickerVisibly, setPickerVisibly] = useState(false)
    const [chosenEmoji, setChosenEmoji] = useState({} as EmojiClickData);
    const {user} = useAppSelector(state => state.userReducer)
    const emoji = chosenEmoji.unified && pickerVisibly ? chosenEmoji.unified : user.emoji
    const dispatch = useAppDispatch()

    // const submitHandler: async (event: React.>) {

    // }

    return(
        <div className={style.dropdown}>
            <button
                className="pl-3 pt-1.5 border-0 w-auto h-auto"
                onClick={()=>{
                    if (pickerVisibly) {
                        setChosenEmoji({} as EmojiClickData)
                    }
                    setPickerVisibly(!pickerVisibly)}}
            >
                <Emoji unified={emoji} size={50}/></button>
            <div className={!pickerVisibly ? style.dropdown_content : style.show}>
                {<EmojiPicker onEmojiClick={(emoji:EmojiClickData, event) => {setChosenEmoji(emoji)}}/>}
                <button
                    onClick={() => {
                        if (chosenEmoji && chosenEmoji.unified) {
                            dispatch(setEmoji({emoji: chosenEmoji.unified}))
                    }}
                        }
                >Выбрать</button>
            </div>
        </div>
    )
}