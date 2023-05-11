import React, {useState} from "react";
import {LoginQuestion, RegistrationQuestion} from "./Question";
import {ErrorMessage} from "./ErrorMessage";
import {useAppDispatch, useAppSelector} from "../hooks/redux";
import {fetchUser} from "../store/reducers/UserSlice";

interface SignFormProps {
    login: boolean
    onRegistration: () => void
    onLogin: () => void
    onClose: () => void
}

export function SignForm({login, onRegistration, onLogin, onClose}:SignFormProps) {
    const dispatch = useAppDispatch()
    const [valueLogin, setValueLogin] = useState('')
    const [valuePassword, setValuePassword] = useState('')
    const [valueEmail, setValueEmail] = useState('')
    const [errorForm, setErrorForm] = useState('')
    const {error} = useAppSelector(state => state.userReducer)

    const btnName = login ? "Войти" : "Зарегистрироваться"

    const submitHandler = async (event: React.FormEvent) => {
        setErrorForm('')
        event.preventDefault()

        if (valueLogin.trim().length === 0 || valuePassword.trim().length === 0 ||
            (!login && valueEmail.length === 0)) {
            setErrorForm('Пожалуйста введите валидные данные')
            return
        }
        dispatch(fetchUser({
            login: valueLogin,
            password: valuePassword,
            email: valueEmail,
            isLogin: login
        }))
        if (!(error as string)) {
            onClose()
        }
    }

    const loginHandler = () => {
        if (login) {
            onRegistration()
        } else {
            onLogin()
        }
        setErrorForm('')
    }

    return (
        <form onSubmit={submitHandler}>
            <input
                type="text"
                placeholder='Введите логин'
                value={valueLogin}
                onChange={(event)=>{setValueLogin(event.target.value)}}
            />
            {!login &&
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
            <button
                type='submit'
            >
                {btnName}
            </button>
            {errorForm && <ErrorMessage error={errorForm}/>}
            {error as string && <ErrorMessage error={error as string}/>}
            {!login ?
                <LoginQuestion onLogin={loginHandler}/> :
                <RegistrationQuestion onRegistration={loginHandler}/>}

        </form>
    )
}