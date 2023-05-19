import {IList} from "../../models/IList";
import styles from "./List.module.css";
import React from "react";
import {DoramaShort} from "../Dorama/Short/Short";
import {useAppSelector} from "../../hooks/redux";
import {IPhoto} from "../../models/IPhoto";
import {instance} from "../../http-common";
import {errorHandler} from "../../hooks/errorHandler";

interface ListProps {
    list:IList
}

export function List({list}:ListProps) {
    const {isAuth, user} = useAppSelector(state => state.userReducer)
    const canAddToFav = user.username !== list.creator_name


    const addToFav = async () => {
        try {
            await instance.post<void>("/user/favorite", {
                id: list.id
            })
        } catch (e:unknown) {
            errorHandler(e)
        }
    }

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
                        <button className={styles.FavButton} onClick={addToFav}>
                            <i className="fa-regular fa-2x fa-heart" style={{color: "#a52727"}}></i>
                        </button>
                    }
                </div>
            </div>
        </div>
        <h2>Дорамы</h2>
        {console.log(list.doramas)}
        {list.doramas ? [...list.doramas].map(drm =>
            <DoramaShort
                dorama={drm}
                isEdit={isAuth && !canAddToFav}
                idList={list.id}
                key={drm.id}
            />) :
            <p className="text-center text-xl">
                Нет дорам
            </p>}
    </>}</>)
}