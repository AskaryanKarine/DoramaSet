import {ISubscribe} from "../models/ISubscribe";
import {useEffect, useMemo, useState} from "react";
import {instance} from "../http-common";
import {AxiosError} from "axios";
import {IError} from "../models/IError";
import {IDorama} from "../models/IDorama";
import {errorHandler} from "./errorHandler";
import {Simulate} from "react-dom/test-utils";
import copy = Simulate.copy;

interface DoramaResponse {
    data: IDorama[]
}

export function useAllDorama() {
    const [allDorama, setAllDorama] = useState<IDorama[]>([])
    const [loading, setLoading] = useState(false)
    const [doramaErr, setDoramaErr] = useState('')

    async function fetchDorama() {
        try {
            setDoramaErr('')
            setLoading(true)
            const response = await instance.get<DoramaResponse>('/dorama/')
            setAllDorama(response.data.data)
            setLoading(false)
        } catch (e: unknown) {
            setLoading(false)
            setDoramaErr(errorHandler(e))
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
            setAllDorama(response.data.data)
            setLoading(false)
        } catch (e: unknown) {
            const error = e as AxiosError<IError>
            setLoading(false)
            if (error.response) {
                if (error.response.status === 400) {
                    console.log("error", error)
                    setAllDorama([])
                }
                setDoramaErr(error.response.data.error)
            }
            if (!error.response) {
                console.error(e)
                throw e;
            }
        }
    }

    function addDorama(newDorama:IDorama) {
        if (newDorama) {
            setAllDorama(prev=>[...prev, newDorama])
        } else {
            setAllDorama([newDorama])
        }
    }

    const resetAllDorama = () => {
        fetchDorama().then(_ => {})
    }

    useEffect(()=>{
        fetchDorama().then()
    }, [])

    return {allDorama, doramaErr, loading, addDorama, findDorama, resetDorama: resetAllDorama}
}