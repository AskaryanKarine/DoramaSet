import {ISubscribe} from "../models/ISubscribe";
import {useEffect, useState} from "react";
import {instance} from "../http-common";
import {AxiosError} from "axios";
import {IError} from "../models/IError";
import {IDorama} from "../models/IDorama";

interface DoramaResponse {
    data: IDorama[]
}

export function useDorama() {
    const [dorama, setDorama] = useState<IDorama[]>([])
    const [loading, setLoading] = useState(false)
    const [doramaErr, setDoramaErr] = useState('')

    function addDorama(dorama:IDorama) {
        setDorama(prev=>[...prev, dorama])
    }

    function updateDorama(upDorama:IDorama) {
        dorama.forEach(x => {
            x = upDorama.id === x.id ? upDorama : x
        })
        setDorama(dorama)
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

    async function fetchDorama() {
        try {
            setDoramaErr('')
            setLoading(true)
            const response = await instance.get<DoramaResponse>('/dorama/')
            setDorama(response.data.data)
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

    async function findDorama(name:string) {
        try {
            setDoramaErr("")
            setLoading(true)
            const response = await instance.get<DoramaResponse>("/find/dorama/", {
                params: {
                    name: name
                }
            })
            setDorama(response.data.data)
            setLoading(false)
        } catch (e: unknown) {
            const error = e as AxiosError<IError>
            setLoading(false)
            if (error.response) {
                if (error.status === 400) {
                    setDorama([])
                }
                setDoramaErr(error.response.data.error)
            } else {
                setDoramaErr(error.message)
            }
        }
    }

    const resetDorama = () => {
        fetchDorama().then(_ => {})
    }

    useEffect(()=>{
        fetchDorama().then(r => {})
    }, [])

    return {dorama, doramaErr, loading, addDorama, updateDorama, findDorama, resetDorama}
}