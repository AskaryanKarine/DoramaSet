import {IDorama} from "../../../models/IDorama";
import styles from "./DoramaPreview.module.css"
import {upToFirst} from "../../../hooks/upToFirst";
import {useState} from "react";
import {Modal} from "../../Modal/Modal";
import {Dorama} from "../Dorama";
import {UpdateForm} from "../Form/UpdateForm";
import {AdminPanel} from "../../Admin/Panel/AdminPanel";

interface DoramaPreviewProps {
    dorama: IDorama
}

export function DoramaPreview({dorama}:DoramaPreviewProps) {
    const [modalVisible, setModalVisible] = useState(false)
    const [editVisible, setEditVisible] = useState(false)

    return (
        <>{dorama && <>
            <div className={styles.container}>
                <div className="flex justify-center items-center h-full max-w-[150px] w-full">
                    {dorama.posters ?
                        <img src={dorama.posters[0].url} alt={dorama.name}></img> :
                        <i className="fa-regular fa-image fa-5x" style={{color: "#787d87"}}></i>
                    }
                </div>
                <div className={styles.info}>
                    <div>
                        <p className="text-3xl">{dorama.name}</p>
                        <p>{upToFirst(dorama.genre)}</p>
                        <p>Год выхода {dorama.release_year}, {dorama.status}</p>
                        <p>Количество эпизодов: {dorama.episodes ? dorama.episodes.length : "0"}</p>
                    </div>
                    <button
                        className={styles.more}
                        onClick={()=>{setModalVisible(true)}}
                    >Подробнее</button>
                </div>
                <AdminPanel
                    onDelete={()=>{}}
                    onEdit={()=>{setEditVisible(true)}}
                />
            </div>
            {modalVisible &&
                <Modal
                    title={dorama.name}
                    onClose={()=>{setModalVisible(false)}}
                >
                <Dorama dorama={dorama}/>
            </Modal>}
            {editVisible &&
                <Modal
                    title="Изменить дораму"
                    onClose={() => {setEditVisible(false)
                }}
            >
                <UpdateForm dorama={dorama}/>
            </Modal>}
        </>}</>
    )
}