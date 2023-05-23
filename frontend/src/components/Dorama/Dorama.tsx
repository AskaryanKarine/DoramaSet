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
import {Review} from "../Review/Review";

interface DoramaProps {
    dorama: IDorama
}

export function Dorama({dorama}:DoramaProps) {
    const [rate, setRate] = useState(dorama.rate ? dorama.rate : 0)
    const [cnt, setCnt] = useState(dorama.cnt_rate ? dorama.cnt_rate : 0)
    const [visible, setVisible] = useState(false)
    const {isAuth} = useAppSelector(state => state.userReducer)
    const {userCollection} = useCollection()
    const [currentReviews, setCurrentReviews] = useState(dorama.reviews)
    const updateRateAdd = (mark:number) => {
        if (cnt) {
            setCnt(cnt + 1)
        } else {
            setCnt(1)
        }
        if (rate) {
            const tmp = (rate + mark)/(cnt+1)
            setRate(tmp)
        } else {
            setRate(mark)
        }
    }

    const updateRateDel = (mark:number) => {
        if (cnt) {
            setCnt(cnt - 1)
        } else {
            setCnt(1)
        }
        if (rate) {
            const tmp = (rate * cnt - mark)/(cnt-1)
            setRate(tmp)
        } else {
            setRate(mark)
        }
    }
    return (<>{dorama && <>
        <div className={styles.container}>
            <Slider photo={dorama.posters}/>
            <div className={styles.info}>
                <p>Год выхода {dorama.release_year}, {dorama.status}</p>
                <p>Жанр: {dorama.genre}</p>
                {cnt > 0 &&
                    <p>Рейтинг: {rate}/5, {cnt}</p>
                }
                <p>Количество эпизодов: {dorama.episodes ? dorama.episodes.length : "0"}</p>
                <p>Описание: {dorama.description}</p>
                {isAuth && <button className="w-[100%]" onClick={()=>{setVisible(true)}}>Добавить в список</button>}
            </div>
        </div>
        <Review
            id={dorama.id}
            reviews={currentReviews}
            add={updateRateAdd}
            del={updateRateDel}
        />
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

