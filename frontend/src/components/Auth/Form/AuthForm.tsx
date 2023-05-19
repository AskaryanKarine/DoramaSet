import React, {useEffect, useState} from "react";
import {AuthQuestion} from "../Question/Question";
import {ErrorMessage} from "../../ErrorMessage/ErrorMessage";
import {useAppDispatch, useAppSelector} from "../../../hooks/redux";
import {AuthState} from "../../../models/AuthState";
import {authUser} from "../../../store/reducers/UserSlice";
import {Loading} from "../../Loading/Loading";

interface SignFormProps {
    authState: AuthState
    onRegistration: () => void
    onSignIn: () => void
    onClose: () => void
}

export function AuthForm({authState, onRegistration, onSignIn, onClose}:SignFormProps) {
    const dispatch = useAppDispatch()
    let {error, loading, user, isAuth} = useAppSelector(state => state.userReducer)
    const btnName = authState === AuthState.SIGN_IN ? "Войти" : "Зарегистрироваться"

    const [valueLogin, setValueLogin] = useState('')
    const [valuePassword, setValuePassword] = useState('')
    const [valueEmail, setValueEmail] = useState('')
    const [errorValue, setErrorValue] = useState('')

    const submitHandler = async (event: React.FormEvent) => {
        event.preventDefault()

        if (valueLogin.trim().length === 0 || valuePassword.trim().length === 0 ||
            (authState === AuthState.REGISTRATION && valueEmail.length === 0)) {
            setErrorValue('Пожалуйста введите валидные данные')
            return
        }

        await dispatch(authUser({
                login: valueLogin,
                password: valuePassword,
                email: valueEmail,
                authState: authState})
        )
    }

    const authHandler = () => {
        setErrorValue('')
        switch (authState) {
            case AuthState.SIGN_IN:
                onSignIn();
                break;
            case AuthState.REGISTRATION:
                onRegistration()
                break
        }
    }

    useEffect(() => {
        if ((loading == false) && user.username && isAuth) {
            onClose()
        }
    }, [loading, user])

    const errMsg = errorValue.length > 0 ? errorValue : error as string

    return (
        <>
            <form onSubmit={submitHandler} onChange={() => {setErrorValue("")}}>
                <input
                    type="text"
                    placeholder='Введите логин'
                    value={valueLogin}
                    onChange={(event)=>{setValueLogin(event.target.value)}}
                />
                {authState === AuthState.REGISTRATION &&
                    <input
                        type="text"
                        placeholder="Введите email"
                        value={valueEmail}
                        onChange={(event) => {setValueEmail(event.target.value)}}
                    />}
                <input
                    type="password"
                    placeholder="Введите пароль"
                    value={valuePassword}
                    onChange={(event)=>{setValuePassword(event.target.value)}}
                />
                <button type='submit'>
                    {btnName}
                </button>
                <AuthQuestion onAction={authHandler} status={authState}/>
            </form>
            {(errMsg) && <ErrorMessage error={errMsg}/>}
            {loading && <Loading/>}
        </>
    )
}