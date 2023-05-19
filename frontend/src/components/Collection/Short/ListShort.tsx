import {IDorama} from "../../../models/IDorama";
import styles from "./ListShort.module.css";
import {IList} from "../../../models/IList";
import {IPhoto} from "../../../models/IPhoto";
import {instance} from "../../../http-common";
import {errorHandler} from "../../../hooks/errorHandler";

interface ListShortProps {
    list: IList
    isEdit: boolean
    idDorama?:number
}

export function ListShort({list, isEdit, idDorama}:ListShortProps) {

    let lenList = list.doramas ? list.doramas.length : 0
    const onCreate = async () => {
        const url = ["/list/", list.id].join("")
        try {
            await instance.post<void>(url, {
                id: idDorama
            }).then(_ => {lenList = lenList +1})
        } catch (e:unknown) {
            errorHandler(e)
        }
    }

    return (
        <>
            <div className={styles.container}>
                <p>{list.name}</p>
                <p>Количество: {lenList}</p>
                {isEdit &&
                    <button onClick={onCreate}>Добавить в список</button>}
            </div>
        </>
    )
}