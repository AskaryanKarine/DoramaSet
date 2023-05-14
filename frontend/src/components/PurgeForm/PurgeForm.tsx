import React, {useEffect, useState} from "react";
import {ErrorMessage} from "../ErrorMessage/ErrorMessage";
import {useAppDispatch, useAppSelector} from "../../hooks/redux";
import {earnPoint, getUser} from "../../store/reducers/UserSlice";

interface PurgeFormProps {
    onClose: ()=> void
}

export function PurgeForm({onClose}:PurgeFormProps) {
    const dispatch = useAppDispatch()
    const [points, setPoints] = useState("")
    const [errorValue, setErrorValue] = useState('')
    const {user} = useAppSelector(state => state.userReducer)

    const submitHandler = async (event: React.FormEvent) => {
        event.preventDefault()

        if (points.trim().length === 0) {
            setErrorValue('Пожалуйста введите валидные данные')
            return
        }
        await dispatch(earnPoint({points}))
    }

    return (
        <>
            <form onSubmit={submitHandler} onChange={()=>{setErrorValue("")}}>
                <input
                    type="text"
                    placeholder="Баллы"
                    value={points}
                    onChange={(event) => {setPoints(event.target.value)}}
                />
                <button type="submit">Пополнить</button>
            </form>
            {(errorValue) && <ErrorMessage error={errorValue}/>}
        </>
    )
}