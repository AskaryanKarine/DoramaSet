import React, {useState} from "react";
import {AuthQuestion} from "../Question/Question";
import {ErrorMessage} from "../../ErrorMessage/ErrorMessage";
import {useAppDispatch, useAppSelector} from "../../../hooks/redux";
import {AuthState} from "../../../models/AuthState";
import {fetchUser} from "../../../store/reducers/UserSlice";

interface SignFormProps {
    authState: AuthState
    onRegistration: () => void
    onSignIn: () => void
    onClose: () => void
}

export function AuthForm({authState, onRegistration, onSignIn, onClose}:SignFormProps) {
    const dispatch = useAppDispatch()

    const [valueLogin, setValueLogin] = useState('')
    const [valuePassword, setValuePassword] = useState('')
    const [valueEmail, setValueEmail] = useState('')
    const [errorValue, setErrorValue] = useState('')
    const {error, loading} = useAppSelector(state => state.userReducer)

    const btnName = authState === AuthState.SIGN_IN ? "Войти" : "Зарегистрироваться"

    const submitHandler = async (event: React.FormEvent) => {
        setErrorValue('')
        event.preventDefault()

        if (valueLogin.trim().length === 0 || valuePassword.trim().length === 0 ||
            (authState === AuthState.REGISTRATION && valueEmail.length === 0)) {
            setErrorValue('Пожалуйста введите валидные данные')
            return
        }
        await dispatch(fetchUser({
            login: valueLogin,
            password: valuePassword,
            email: valueEmail,
            authState: authState
        }))

        console.log("error in form", loading, error, (error as string).length)
        // TODO
        if ((loading == false) && ((error as string).length === 0)) {
            onClose()
        }
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
            {(errorValue) && <ErrorMessage error={errorValue}/>}
            {(error as string).length != 0 && <ErrorMessage error={error as string}/>}
        </>
    )
}