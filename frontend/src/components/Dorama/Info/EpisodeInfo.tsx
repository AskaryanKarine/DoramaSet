import styles from "../Dorama.module.css";
import {Episode} from "../Episode/Episode";
import React, {useState} from "react";
import {useEpisodeWithStatus} from "../../../hooks/episodeWithStatus";
import {Loading} from "../../Loading/Loading";
import {Modal} from "../../Modal/Modal";
import {instance} from "../../../http-common";
import {IEpisode} from "../../../models/IEpisode";

interface EpisodeInfoProps {
    id?:number
    isEdit:boolean
}

export function EpisodeInfo({id, isEdit}:EpisodeInfoProps) {
    const {episodeWithStatus, loading, epErr, createEpisode} = useEpisodeWithStatus(id)
    const [modal, setModal] = useState(false)
    const [ep, setEp] = useState("")
    const [season, setSeason] = useState("")

    console.log(episodeWithStatus)

    const submitHandler = (event: React.FormEvent) => {
        event.preventDefault()
        if (id) {
            createEpisode(id, ep, season).then(_=>{setModal(false)})
        }
    }
    
    return (<>
        {loading && <Loading/>}
        <div>
            <div className={styles.addHeader}>
                <h2>Эпизоды</h2>
                {isEdit && <button onClick={()=>{setModal(true)}}>
                    <i className="fa-solid fa-plus fa-border border-2 rounded-full bg-white border-black"></i>
                </button>}
            </div>
            {episodeWithStatus ? [...episodeWithStatus].map(ep =>
                <Episode
                    ep={ep.episode}
                    flag={ep.watching}
                    key={ep.episode.id}
                />) : "Нет эпизодов"}
        </div>
        {modal && <Modal title={"Добавить эпизод"} onClose={()=>{setModal(false)}}>
            <form onSubmit={submitHandler}>
                <input
                    type="text"
                    placeholder="Сезон"
                    value={season}
                    onChange={(event)=>{setSeason(event.target.value)}}
                />
                <input
                    type="text"
                    placeholder="Серия"
                    value={ep}
                    onChange={(event)=>{setEp(event.target.value)}}
                />
                <button type="submit">Добавить</button>
            </form>
        </Modal>}
    </>)
}