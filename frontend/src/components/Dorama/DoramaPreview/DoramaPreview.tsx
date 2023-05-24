import {IDorama} from "../../../models/IDorama";
import styles from "./DoramaPreview.module.css"
import {upToFirst} from "../../../hooks/upToFirst";
import {useEffect, useState} from "react";
import {Modal} from "../../Modal/Modal";
import {Dorama} from "../Dorama";
import {DoramaUpdate} from "../Form/DoramaUpdate";
import {AdminPanel} from "../../Admin/Panel/AdminPanel";
import {useAllDorama} from "../../../hooks/dorama";
import {IEpisode} from "../../../models/IEpisode";

interface DoramaPreviewProps {
    dorama: IDorama
}

export function DoramaPreview({dorama}:DoramaPreviewProps) {
    const [modalVisible, setModalVisible] = useState(false)
    const [editVisible, setEditVisible] = useState(false)
    const [current, setCurrent] = useState(dorama)

    const onUpdateDorama = (dorama:IDorama) => {
        dorama.episodes = current.episodes
        setCurrent(dorama)
        setEditVisible(false)
    }

    const addEpisode = (ep:IEpisode) => {
        if (current.episodes) {
            current.episodes = [...current.episodes, ep]
        } else {
            current.episodes = [ep]
        }
        setCurrent(prev => {
            return {
                ...prev,
                episodes: current.episodes
            }
        })
    }

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
                        <p className="text-3xl">{current.name}</p>
                        <p>{upToFirst(current.genre)}</p>
                        <p>Год выхода {current.release_year}, {current.status}</p>
                        <p>Количество эпизодов: {current.episodes ? current.episodes.length : 0}</p>
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
                <DoramaUpdate dorama={current} onClose={onUpdateDorama} addEpisode={addEpisode}/>
            </Modal>}
        </>}</>
    )
}