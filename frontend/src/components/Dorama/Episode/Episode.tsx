import {IEpisode} from "../../../models/IEpisode";
import {useState} from "react";
import styles from  "./Episode.module.css"
import {useAppSelector} from "../../../hooks/redux";
import {useEpisode} from "../../../hooks/episode";


interface EpisodeProps {
    ep:IEpisode
    flag:boolean
}

export function Episode({ep, flag}:EpisodeProps) {
    const {isAuth} = useAppSelector(state => state.userReducer)
    const [watched, setWatched] = useState(flag)
    const [error, setError] = useState<string[]>([])
    const {epErr, markEpisode} = useEpisode()

    const clickSubmit = () => {
        if (!watched && ep.id) {
            markEpisode(ep.id).then(_ => setError(prevState => [...prevState, epErr]))
            if (epErr.length === 0) {
                setWatched(true)
            }
        }
    }

    return (
        <>
            <div className={styles.container}>
                <button onClick={()=>{clickSubmit()}} disabled={!isAuth || watched}>
                    {watched ? <i className="fa-regular fa-eye"></i> : <i className="fa-regular fa-eye-slash"></i>}
                </button>

                <p>
                    Сезон {ep.num_episode} серия {ep.num_season}
                </p>
            </div>
        </>
    )
}