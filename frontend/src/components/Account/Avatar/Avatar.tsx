import style from "./Avatar.module.css";
import {useAppDispatch, useAppSelector} from "../../../hooks/redux";
import React, {useEffect, useState} from "react";
import {HexColorPicker} from "react-colorful";
import {IUser} from "../../../models/IUser";
import {setColor} from "../../../store/reducers/UserSlice";

export function Avatar() {
    const dispatch = useAppDispatch()
    const {user, loading, error} = useAppSelector(state => state.userReducer)
    const [colorVisibly, setColorVisibly] = useState(false)
    const colorStatus = user.sub && (user.sub.access_lvl >= 3) ? user.color : "#000000"
    const [color, setCurColor] = useState(colorStatus);


    const iconStyle = "fa-solid fa-user fa-5x fa-border border-0 " +
        "rounded-full bg-white border-black"


    useEffect(() => {
        if ((user === {} as IUser)) {}
        else {
            if (user.sub && user.sub.access_lvl >= 3) {
                setCurColor(user.color)
                console.log(user, color)
            }
        }
    }, [user]);

    return (
        <>
            <div className={style.dropdown}>
                <div className={style.iconName} onClick={
                    ()=>{
                        if (user.sub && user.sub.access_lvl >= 3)
                        {
                            setColorVisibly(!colorVisibly)
                        }
                    }
                }>
                <span style={{color: color, border:5, borderColor: color, borderStyle:"solid", borderRadius: 99999999}}>
                    <i className={iconStyle}></i>
                </span>
                </div>
                <div className={!colorVisibly ? style.dropdown_content : style.show}>
                    <HexColorPicker color={color} onChange={setCurColor}/>
                    <button
                        onClick={() => {
                             dispatch(setColor({color: color}))
                            if (!loading && !error) {
                                setColorVisibly(false)
                            }
                        }
                    }
                    >Выбрать</button>
                </div>
            </div>
        </>
    )
}