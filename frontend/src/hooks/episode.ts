import {IDorama} from "../models/IDorama";
import {useEffect, useState} from "react";
import {instance} from "../http-common";
import {AxiosError} from "axios";
import {IError} from "../models/IError";
import {IEpisode} from "../models/IEpisode";

interface EpisodeResponse {
    Data: IEpisode[]
}

export function useEpisode() {
    const [episode, setEpisode] = useState<IEpisode[]>([])
    const [loading, setLoading] = useState(false)
    const [epErr, setEpErr] = useState('')

    function addEpisode(episode:IEpisode) {
        setEpisode(prev=>[...prev, episode])
    }

    async function fetchWatchingEpisode(idDorama:number) {
        try {
            setEpErr('')
            setLoading(true)
            const response = await instance.get<EpisodeResponse>('/user/episode', {
                params: {
                    id: idDorama,
                }
            })
            setEpisode(response.data.Data)
            setLoading(false)
        } catch (e: unknown) {
            const error = e as AxiosError<IError>
            setLoading(false)
            if (error.response) {
                setEpErr(error.response.data.error)
            } else {
                setEpErr(error.message)
            }
        }
    }

    return {episode, epErr, loading, addEpisode, fetchWatchingEpisode}
}