import {IDorama} from "../models/IDorama";
import {useEffect, useState} from "react";
import {instance} from "../http-common";
import {AxiosError} from "axios";
import {IError} from "../models/IError";
import {IEpisode} from "../models/IEpisode";

interface episodeWithStatus {
    episode: IEpisode
    watching: boolean
}

interface EpisodeResponse {
    data: episodeWithStatus[]
}

export function useEpisodeWithStatus(id?:number) {
    const [episodeWithStatus, setEpisodeWithStatus] = useState<episodeWithStatus[]>([])
    const [epErr, setEpErr] = useState('')
    const [loading, setLoading] = useState(false)

    async function fetchWatchingEpisode(idDorama:number) {
        try {
            setEpErr('')
            setLoading(true)

            const response = await instance.get<EpisodeResponse>('/user/episode', {
                params: {
                    id: idDorama,
                }
            })
            setEpisodeWithStatus(response.data.data)
            setLoading(false)
        } catch (e: unknown) {
            setLoading(false)
            const error = e as AxiosError<IError>
            if (error.response) {
                setEpErr(error.response.data.error)
            } else {
                setEpErr(error.message)
            }
        }
    }

    async function createEpisode(id:number, ep: string, season:string) {
        const url = ["/dorama/", id, "/episode"].join("")
        try {
            const response = await instance.post<{data: IEpisode}>(url, {
                num_episode: parseInt(ep),
                num_season: parseInt(season)
            })
            const epSt: episodeWithStatus = {
                episode: response.data.data,
                watching: false
            }
            console.log(response.data.data, epSt)
            addEpisode(epSt)
        } catch (e:unknown) {

        }
    }

    const addEpisode = (episode:episodeWithStatus) => {
        if (episodeWithStatus) {
            setEpisodeWithStatus(prevState => [...prevState, episode])
        } else {
            setEpisodeWithStatus(()=>[episode])
        }
    }

    useEffect(()=>{
        if (id) {
            fetchWatchingEpisode(id).then(r => {setEpErr(epErr)})
        }
    }, [])


    return {episodeWithStatus, epErr, loading, createEpisode}
}