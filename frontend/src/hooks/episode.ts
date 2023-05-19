import {useState} from "react";
import {instance} from "../http-common";
import {AxiosError} from "axios";
import {IError} from "../models/IError";

export function useEpisode() {
    const [epErr, setEpEpr] = useState("")

    async function markEpisode(idEp:number) {
        try {
            setEpEpr("")
            await instance.post('/user/episode', {
                id: idEp,
            })
        } catch (e: unknown) {
            const error = e as AxiosError<IError>
            if (error.response) {
                setEpEpr(error.response.data.error)
            } else {
                setEpEpr(error.message)
            }
        }
    }

    return {epErr, markEpisode}

}