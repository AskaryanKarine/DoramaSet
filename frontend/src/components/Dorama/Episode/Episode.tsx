import {IEpisode} from "../../../models/IEpisode";
import {useEpisode} from "../../../hooks/episode";
import {useEffect, useState} from "react";
import styles from  "./Episode.module.css"
import {instance} from "../../../http-common";
import {AxiosError} from "axios/index";
import {IError} from "../../../models/IError";


interface EpisodeProps {
    ep:IEpisode
    flag:boolean
}

export function Episode({ep, flag}:EpisodeProps) {
    const [watched, setWatched] = useState(flag)
    const [epErr, setEpErr] = useState("")

    async function markEpisode(idEp:number) {
        try {
            setEpErr('')

            console.log(idEp)
            await instance.post('/user/episode', {
                id: idEp,
            })
        } catch (e: unknown) {
            const error = e as AxiosError<IError>
            if (error.response) {
                setEpErr(error.response.data.error)
            } else {
                setEpErr(error.message)
            }
        }
    }

    const clickSubmit = () => {
        if (!watched) {
            markEpisode(ep.id)
            if (epErr.length === 0) {
                setWatched(true)
            }
        }
    }

    return (
        <>
            <div className={styles.container}>
                <button onClick={()=>{clickSubmit()}}>
                    {watched ? <i className="fa-regular fa-eye"></i> : <i className="fa-regular fa-eye-slash"></i>}
                </button>

                <p>
                    Сезон {ep.num_season} серия {ep.num_episode}
                </p>
            </div>
        </>
    )
}