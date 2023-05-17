import {useEffect, useState} from "react";
import {IList} from "../models/IList";
import {instance} from "../http-common";
import {AxiosError} from "axios";
import {IError} from "../models/IError";
import {useAppSelector} from "./redux";
import {IDorama} from "../models/IDorama";

interface listResponse {
    data: IList[]
}

export function useCollection() {
    const [publicCollection, setPublicCollection] = useState<IList[]>([])
    const [userCollection, setUserCollection] = useState<IList[]>([])
    const [favCollection, setFavCollection] = useState<IList[]>([])
    const [colErr, setColErr] = useState('')
    const [loading, setLoading] = useState(false)

    async function fetchPublic() {
        try {
            setColErr("")
            setLoading(true)
            const response = await instance.get<listResponse>('/list/public')
            setPublicCollection(response.data.data)
            setLoading(false)
        } catch (e: unknown) {
            setLoading(false)
            const error = e as AxiosError<IError>
            if (error.response) {
                setColErr(error.response.data.error)
            } else {
                setColErr(error.message)
            }
        }
    }

    async function fetchFav() {
        try {
            setColErr("")
            setLoading(true)
            const response = await instance.get<listResponse>('/user/favorite')
            setFavCollection(response.data.data)
            setLoading(false)
        } catch (e: unknown) {
            setLoading(false)
            const error = e as AxiosError<IError>
            if (error.response) {
                setColErr(error.response.data.error)
            } else {
                setColErr(error.message)
            }
        }
    }

    async function fetchUserList() {
        try {
            setColErr("")
            setLoading(true)
            const response = await instance.get<listResponse>('/user/list')
            setUserCollection(response.data.data)
            setLoading(false)
        } catch (e: unknown) {
            setLoading(false)
            const error = e as AxiosError<IError>
            if (error.response) {
                setColErr(error.response.data.error)
            } else {
                setColErr(error.message)
            }
        }
    }

    function addPrivateList(list:IList) {
        setUserCollection(prev=>[...prev, list])
    }

    useEffect(()=>{
        fetchPublic()
        fetchFav()
        fetchUserList()
    }, [])

    return {publicCollection, userCollection, favCollection, loading, colErr, addPrivateList}
}