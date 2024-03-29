import {IList} from "../../../models/IList";
import styles from "./ListPreview.module.css";
import {Modal} from "../../Modal/Modal";
import {useState} from "react";
import {List} from "../List";
import {useAppSelector} from "../../../hooks/redux";

interface ListPreviewProps {
    list:IList
    inFav?:boolean
}

export function ListPreview({list, inFav}:ListPreviewProps) {
    const [listLen, setListLen] = useState(list.doramas ? list.doramas.length : 0)
    const [modalVisible, setModalVisible] = useState(false)

    return (<>{list && <>
        <div className={styles.container}>
            <div className={styles.info}>
                <div>
                    <p className="text-3xl">{list.name}</p>
                    <p>Создатель: {list.creator_name}</p>
                    <p>Тип: {list.type}</p>
                    <p>Количество дорам: {listLen}</p>
                </div>
                <button
                    className={styles.more}
                    onClick={()=>{setModalVisible(true)}}
                >Подробнее</button>
            </div>
        </div>
        {modalVisible &&
            <Modal
                title={list.name}
                onClose={()=>{setModalVisible(false)}}
            >
                <List list={list} onDelete={()=>setListLen(listLen - 1)} inFav={inFav}/>
            </Modal>}
    </>}</>)
}