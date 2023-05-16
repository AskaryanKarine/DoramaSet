import {IDorama} from "../../../models/IDorama";
import styles from "./DoramaPreview.module.css"
import {toUpperFirst} from "../../../hooks/toUpperFirst";
import {useState} from "react";
import {AdminPanel} from "../../AdminPanel/AdminPanel";
import {Modal} from "../../Modal/Modal";
import {Dorama} from "../Dorama";

interface DoramaPreviewProps {
    dorama: IDorama
}

export function DoramaPreview({dorama}:DoramaPreviewProps) {
    const [modalVisible, setModalVisible] = useState(false)

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
                        <p>{toUpperFirst(dorama.genre)}</p>
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
                    onEdit={()=>{}}
                />
            </div>
            {modalVisible &&
                <Modal
                    title={dorama.name}
                    onClose={()=>{setModalVisible(false)}}
                >
                <Dorama dorama={dorama}/>
            </Modal>}
        </>}</>
    )
}