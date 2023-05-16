import {ISubscribe} from "../models/ISubscribe";
import {useEffect, useState} from "react";
import {instance} from "../http-common";
import {AxiosError} from "axios";
import {IError} from "../models/IError";
import {IDorama} from "../models/IDorama";

interface DoramaResponse {
    Data: IDorama[]
}

export function useDorama() {
    const [dorama, setDorama] = useState<IDorama[]>([])
    const [loading, setLoading] = useState(false)
    const [doramaErr, setDoramaErr] = useState('')

    function addDorama(dorama:IDorama) {
        setDorama(prev=>[...prev, dorama])
    }

    function delDorama(id:number) {
        let i = 0
        for (i = 0; i < dorama.length; i++) {
            if (dorama[i].id === id) {
                break
            }
        }
        delete dorama[i]
        setDorama(dorama)
    }

    function updateDorama(upDorama:IDorama) {
        dorama.forEach(x => {
            x = upDorama.id === x.id ? upDorama : x
        })
        setDorama(dorama)
    }

    async function fetchDorama() {
        try {
            setDoramaErr('')
            setLoading(true)
            const response = await instance.get<DoramaResponse>('/dorama/')
            setDorama(response.data.Data)
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
        fetchDorama()
    }, [])

    return {dorama, doramaErr, loading, addDorama, delDorama, updateDorama}
}