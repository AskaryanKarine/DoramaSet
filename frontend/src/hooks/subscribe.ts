import {useEffect, useState} from "react";
import {ISubscribe} from "../models/ISubscribe";
import axios, {AxiosError} from "axios";
import {instance} from "../http-common";
import {IError} from "../models/IError";


interface SubscribeResponse {
    data: ISubscribe[]
}

export function useSubscribe() {
    const [subscribes, setSubscribes] = useState<ISubscribe[]>([])
    const [loading, setLoading] = useState(false)
    const [subErr, setSubErr] = useState('')

    async function fetchSubscribe() {
        try {
            setSubErr('')
            setLoading(true)
            const response = await instance.get<SubscribeResponse>('/subscription/')
            setSubscribes(response.data.data)
            setLoading(false)
        } catch (e: unknown) {
            const error = e as AxiosError<IError>
            setLoading(false)
            if (error.response) {
                setSubErr(error.response.data.error)
            } else {
                setSubErr(error.message)
            }
        }
    }

    useEffect(()=>{
        fetchSubscribe()
    }, [])

    return {subscribes, subErr, loading}
}