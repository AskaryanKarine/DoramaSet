import styles from "./Account.module.css"
import {useState} from "react";
import {Nickname} from "./Nickname/Nickname";
import {Modal} from "../Modal/Modal";
import {ManageSubscribe} from "../ManageSubscribe/ManageSubscribe";
import {PurgeForm} from "../PurgeForm/PurgeForm";

export function Account() {
    const [subscribeVisible, setSubscribeVisible] = useState(false)
    const [purgeVisible, setPurgeVisible] = useState(false)

    return (
        <>
            <div className={styles.leftContainer}>
                    <Nickname/>
                    <button
                        onClick={()=>{setSubscribeVisible(true)}}
                    >Управление подпиской</button>
                    <button
                        onClick={()=>{setPurgeVisible(true)}}
                    >Пополнение баллов</button>
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