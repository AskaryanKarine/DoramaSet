import {IList} from "../../models/IList";
import styles from "./List.module.css";
import React, {useState} from "react";
import {DoramaShort} from "../Dorama/Short/Short";
import {useAppSelector} from "../../hooks/redux";
import {IPhoto} from "../../models/IPhoto";
import {instance} from "../../http-common";
import {errorHandler} from "../../hooks/errorHandler";
import {IDorama} from "../../models/IDorama";

interface ListProps {
    list:IList
    onDelete?:()=>void
    inFav?:boolean
}

export function List({list, onDelete, inFav}:ListProps) {
    const [disable, setDisable] = useState(inFav)
    const [listDorama, setListDorama] = useState(list.doramas)
    const {isAuth, user} = useAppSelector(state => state.userReducer)
    const canAddToFav = user.username !== list.creator_name

    const deleteDorama = (dorama:IDorama) => {
        if (onDelete) {
            onDelete()
        }
        if (listDorama) {
            const index = listDorama.indexOf(dorama)
            if (index !== -1) {
                listDorama.splice(index, 1)
            }
        }
        setListDorama(listDorama)
    }

    const addToFav = async () => {
        try {
            await instance.post<void>("/user/favorite", {
                id: list.id
            }).then(_ => {setDisable(!disable)})
        } catch (e:unknown) {
            errorHandler(e)
        }
    }
    console.log(list, disable)
    return(<>{list && <>
        <div className={styles.container}>
            <div className={styles.info}>
                <div>
                    <p>Создатель: {list.creator_name}</p>
                    <p>Тип: {list.type}</p>
                    <p>Описание: {list.description}</p>
                </div>
                <div>
                    {isAuth && canAddToFav &&
                        <button className={styles.FavButton} onClick={addToFav} disabled={disable}>
                            {!disable ? <i className="fa-regular fa-2x fa-heart" style={{color: "#a52727"}}></i> :
                                <i className="fa-solid fa-2x fa-heart" style={{color: "#a52727"}}></i>
                            }
                        </button>
                    }
                </div>
            </div>
        </div>
        <h2>Дорамы</h2>
        {listDorama && listDorama.length > 0 ? [...listDorama].map(drm =>
            <DoramaShort
                dorama={drm}
                isEdit={isAuth && !canAddToFav}
                idList={list.id}
                key={drm.id}
                onDelete={deleteDorama}
            />) :
            <p className="text-center text-xl">
                Нет дорам
            </p>}
    </>}</>)
}