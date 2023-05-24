import {useEffect, useMemo, useState} from "react";
import {instance} from "../http-common";
import {AxiosError} from "axios";
import {IError} from "../models/IError";
import {IStaff} from "../models/IStaff";
import {IDorama} from "../models/IDorama";

interface staffResponse {
    data: IStaff[]
}

export function useStaff(idDorama?:number) {
    const [staff, setStaff] = useState<IStaff[]>([])
    const [staffLoading, setStaffLoading] = useState(false)
    const [staffErr, setStaffErr] = useState('')
    const [staffDorama, setStaffDorama] = useState<IStaff[]>([])


    async function fetchStaff() {
        try {
            setStaffErr('')
            setStaffLoading(true)
            const response = await instance.get<staffResponse>('/staff/')
            setStaff(response.data.data)
            setStaffLoading(false)
        } catch (e: unknown) {
            const error = e as AxiosError<IError>
            setStaffLoading(false)
            if (error.response) {
                setStaffErr(error.response.data.error)
            } else {
                setStaffErr(error.message)
            }
        }
    }

    async function getStaffByDoramaId (idDorama:number) {
        try {
            setStaffLoading(true)
            const url = ["/dorama", idDorama, "staff"].join("/")
            const response = await instance.get<staffResponse>(url)
            setStaffDorama(response.data.data)
            setStaffLoading(false)
        } catch (e: unknown) {
            setStaffLoading(false)
            const error = e as AxiosError<IError>
            if (error.response) {
                setStaffErr(error.response.data.error)
            } else {
                throw new Error(error.message)
            }
        }
    }

    async function findStaff(name:string) {
        try {
            setStaffErr("")
            setStaffLoading(true)
            const response = await instance.get<staffResponse>("/find/staff/", {
                params: {
                    name: name
                }
            })
            setStaff(response.data.data)
            setStaffLoading(false)
        } catch (e: unknown) {
            const error = e as AxiosError<IError>
            setStaffLoading(false)
            if (error.response) {
                if (error.response.status === 400) {
                    setStaff([])
                }
                setStaffErr(error.response.data.error)
            } else {
                setStaffErr(error.message)
            }
        }
    }

    function addAllStaff(staff:IStaff) {
        if (staff) {
            setStaff(prev => [...prev, staff])
        } else {
            setStaff([staff])
        }
    }

    function addStaff(staff:IStaff) {
        if (staffDorama) {
            setStaffDorama(prev=>[...prev, staff])
        } else {
            setStaffDorama(()=>[staff])
        }
    }

    const resetStaff = () => {
        fetchStaff().then(_ => {})
    }

    useEffect(()=>{
        if (idDorama) {
            getStaffByDoramaId(idDorama).then(()=>{setStaffErr(staffErr)})
            return
        }
        fetchStaff().then()
    }, [])

    return {staff, staffErr, staffLoading, addStaff, findStaff, resetStaff, staffDorama, addAllStaff}
}