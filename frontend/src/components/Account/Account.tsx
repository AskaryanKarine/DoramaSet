import {useAppDispatch, useAppSelector} from "../../hooks/redux";
import styles from "./Account.module.css"
import {Emoji} from "emoji-picker-react";
import {useEffect, useState} from "react";
import {getUser} from "../../store/reducers/UserSlice";
import {AvatarInfo} from "./AvatarInfo/AvatarInfo";
import {Modal} from "../Modal/Modal";
import {Subscribe} from "../ManageSubscribe/Subscribe/Subscribe";
import {ManageSubscribe} from "../ManageSubscribe/ManageSubscribe";
import {PurgeForm} from "../PurgeForm/PurgeForm";

export function Account() {
    const [subscribeVisible, setSubscribeVisible] = useState(false)
    const [purgeVisible, setPurgeVisible] = useState(false)
    const { user} = useAppSelector(state => state.userReducer)

    const term = user.sub && user.sub.cost > 0 ? ["до", user.lastSub].join(" ") : ""

    return (
        <>
            <h1 className="text-5xl font-bold text-center">Профиль</h1>
            <div>
                <div className="flex items-center m-3">
                    <div className={styles.leftContainer}>
                        <AvatarInfo/>
                        <button
                            onClick={()=>{setSubscribeVisible(true)}}
                        >Управление подпиской</button>
                        <button
                            onClick={()=>{setPurgeVisible(true)}}
                        >Пополнение баллов</button>
                    </div>
                    <div className={styles.rightContainer}>
                        <h2>Информация</h2>
                        <p>Подписка {user.sub && user.sub.name} {term} </p>
                        <p>Почта: {user.email}</p>
                        <p>Дата регистрации: {user.regData}</p>
                    </div>
                </div>
            </div>
            {subscribeVisible &&
                <Modal onClose={() => {setSubscribeVisible(false)}}
                       title={"Управление подпиской"}>
                   <ManageSubscribe/>
                </Modal>}
            {purgeVisible &&
            <Modal onClose={()=>{setPurgeVisible(false)}} title={"Пополнение баланса"}>
                <PurgeForm onClose={()=>{setPurgeVisible(false)}}/>
            </Modal>}
        </>
    )
}