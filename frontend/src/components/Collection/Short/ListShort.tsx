import {IDorama} from "../../../models/IDorama";
import styles from "./ListShort.module.css";
import {IList} from "../../../models/IList";
import {IPhoto} from "../../../models/IPhoto";
import {instance} from "../../../http-common";
import {errorHandler} from "../../../hooks/errorHandler";
import {useState} from "react";

interface ListShortProps {
    list: IList
    isEdit: boolean
    idDorama?:number
}

export function ListShort({list, isEdit, idDorama}:ListShortProps) {
    const [lenList, setLenList] = useState(list.doramas ? list.doramas.length : 0)
    const onCreate = async () => {
        const url = ["/list/", list.id].join("")
        try {
            await instance.post<void>(url, {
                id: idDorama
            }).then(_ => {setLenList(lenList + 1)})
        } catch (e:unknown) {
            errorHandler(e)
        }
    }

    let disable = false
    for (let i = 0; list.doramas && i < list.doramas.length; i++) {
        if (list.doramas[i].id === idDorama) {
            disable = true
            break
        }
    }

    return (
        <>
            <div className={styles.container}>
                <p>{list.name}</p>
                <p>Количество: {lenList}</p>
                {isEdit &&
                    <button disabled={disable} onClick={onCreate}>Добавить в список</button>}
            </div>
        </>
    )
}