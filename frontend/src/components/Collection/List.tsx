import {IList} from "../../models/IList";
import styles from "../Dorama/Dorama.module.css";
import React from "react";
import {DoramaShort} from "../Dorama/Short/Short";

interface ListProps {
    list:IList
    isPublic:boolean
}

export function List({list, isPublic}:ListProps) {
    console.log(list)
    return(<>{list && <>
        <div className={styles.container}>
            <div className={styles.info}>
                <p>Создатель: {list.creator_name}</p>
                <p>Тип: {list.type}</p>
                <p>Описание: {list.description}</p>
                <button className="w-[150%] mt-3 mb-3">Добавить в избранное</button>
            </div>
        </div>
        <h2 className="text-center font-bold text-xl" >Дорамы</h2>
        {list.doramas ? [...list.doramas].map(lst =>
            <DoramaShort
                dorama={lst}
                isEdit={true}
                idList={lst.id}
            />) : "Нет дорам"}
    </>}</>)
}