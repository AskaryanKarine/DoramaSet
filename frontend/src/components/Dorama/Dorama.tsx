import {IDorama} from "../../models/IDorama";
import styles from "./Dorama.module.css"
import {EpisodeInfo} from "./Info/EpisodeInfo";
import {StaffInfo} from "./Info/StaffInfo";
import {Slider} from "../Slider/Slider";
import {useAppSelector} from "../../hooks/redux";
import React, {useState} from "react";
import {Modal} from "../Modal/Modal";
import {useCollection} from "../../hooks/collection";
import {ListShort} from "../Collection/Short/ListShort";

interface DoramaProps {
    dorama: IDorama
}

export function Dorama({dorama}:DoramaProps) {
    const [visible, setVisible] = useState(false)
    const {isAuth} = useAppSelector(state => state.userReducer)
    const {userCollection} = useCollection()
    return (<>{dorama && <>
        <div className={styles.container}>
            <Slider photo={dorama.posters}/>
            <div className={styles.info}>
                <p>Год выхода {dorama.release_year}, {dorama.status}</p>
                <p>Жанр: {dorama.genre}</p>
                <p>Количество эпизодов: {dorama.episodes ? dorama.episodes.length : "0"}</p>
                <p>Описание: {dorama.description}</p>
                {isAuth && <button className="w-[100%]" onClick={()=>{setVisible(true)}}>Добавить в список</button>}
            </div>
        </div>
        <EpisodeInfo id={dorama.id} isEdit={false}/>
        <StaffInfo id={dorama.id} isEdit={false}/>
        {visible && <Modal
            onClose={()=>{setVisible(false)}}
            title={"Выберите список"}
        >
            <div className="grid grid-cols-3">
                {userCollection ? [...userCollection].map(
                    lst => <ListShort list={lst} key={lst.id} isEdit={true} idDorama={dorama.id}/>
                ) : "Нет списков. Создайте!"}

            </div>
        </Modal>}
    </>}</>)
}

