import {IDorama} from "../../models/IDorama";
import styles from "./Dorama.module.css"
import SimpleImageSlider from "react-simple-image-slider";
import {Episode} from "./Episode/Episode";
import {useEffect, useState} from "react";
import {instance} from "../../http-common";
import {AxiosError} from "axios";
import {IError} from "../../models/IError";
import {IStaff} from "../../models/IStaff";
import {StaffShort} from "../Staff/Short/StaffShort";
import {Loading} from "../Loading/Loading";
import {useEpisode} from "../../hooks/episode";
import {IEpisode} from "../../models/IEpisode";

interface DoramaProps {
    dorama: IDorama
}

export function Dorama({dorama}:DoramaProps) {

    const [staff, setStaff] = useState<IStaff[]>([])
    const [loading, setLoading] = useState(false)
    const [doramaErr, setDoramaErr] = useState('')
    const [notWatching, setNotWatching] = useState<IEpisode[]>([])
    const {episode, fetchWatchingEpisode} = useEpisode() // watching

    // useEffect(()=>{
    //     if (dorama.id) {
    //         fetchWatchingEpisode(dorama.id)
    //     }
    // }, [episode])

    // if (dorama.episodes) {
    //     for (let i = 0; i < episode.length; i++) {
    //         const result = dorama.episodes.filter(item => item.id !== episode[i].id)
    //         setNotWatching(result)
    //     }
    // }

    console.log(dorama.episodes, episode)

    async function fetchStaff() {
        try {
            setDoramaErr('')
            setLoading(true)
            const url = ["/dorama", dorama.id, "staff"].join("/")
            const response = await instance.get<{Data:IStaff[]}>(url)
            setStaff(response.data.Data)
            setLoading(false)
        } catch (e: unknown) {
            const error = e as AxiosError<IError>
            setLoading(false)
            if (error.response) {
                setDoramaErr(error.response.data.error)
            } else {
                setDoramaErr(error.message)
            }
        }
    }

    useEffect(()=>{
        fetchStaff()
    }, [])
    console.log(staff)

    return (<>{staff && <>
        {loading && <Loading/>}
        <div className={styles.container}>
            <div className="w-auto mr-5 relative">
                {dorama.posters ? <SimpleImageSlider
                    width={300}
                    height={400}
                    autoPlay={true}
                    images={dorama.posters}
                    showBullets={false}
                    showNavs={true}
                    navSize={20}
                /> : <i className="fa-regular fa-image fa-5x" style={{color: "#787d87"}}></i>}
            </div>
            <div className={styles.info}>
                <p>Год выхода {dorama.release_year}, {dorama.status}</p>
                <p>Жанр: {dorama.genre}</p>
                <p>Количество эпизодов: {dorama.episodes ? dorama.episodes.length : "0"}</p>
                <p>Описание: {dorama.description}</p>
            </div>
        </div>
        <div>
            <h2 className="text-center font-bold text-xl mt-5" >Эпизоды</h2>
            {dorama.episodes && dorama.episodes.length > 0 && dorama.id ? [...dorama.episodes].map(
                ep => <Episode ep={ep} flag={false} key={ep.id}/>
            ) : "Нет эпизодов"}
            {/*{episode.length > 0 ? [...watching].map(*/}
            {/*    ep => <Episode ep={ep} flag={true} key={ep.id}/>*/}
            {/*) : "Нет эпизодов"}*/}
        </div>
        <div>
            <h2 className="text-center font-bold text-xl mb-3 mt-3" >Актеры</h2>
            <div className="grid grid-cols-5">
                {staff ? [...staff].map(
                    staff => <StaffShort staff={staff} key={staff.id}/>
                ) : "Нет стаффа"}

            </div>
        </div>

    </>}</>)
}

