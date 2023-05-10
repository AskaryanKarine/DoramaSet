import React, {useState} from "react";
import {RegistrationQuestion} from "./RegistrationQuestion";
import {LoginQuestion} from "./LoginQuestion";

interface RegistrationFormProps {
    onLogin: () => void
}

export function RegistrationForm({onLogin}:RegistrationFormProps) {
    const [valueLogin, setValueLogin] = useState('')
    const [valuePassword, setValuePassword] = useState('')
    const [valueEmail, setValueEmail] = useState('')

    const [error, setError] = useState('')

    const submitHandler = async (event: React.FormEvent) => {
        setError('')
        event.preventDefault()

        if (valueLogin.trim().length === 0 || valuePassword.trim().length === 0
            || valueEmail.length === 0) {
            setError('Please enter valid title')
        }
    }

    return (
        <form onSubmit={submitHandler}>
            <input
                type="text"
                placeholder='Введите имя пользователя'
                value={valueLogin}
                onChange={(event)=>{setValueLogin(event.target.value)}}
            />
            <input
                type="password"
                placeholder="Введите пароль"
                value={valuePassword}
                onChange={(event)=>{setValuePassword(event.target.value)}}
            />
            <input
                type="text"
                placeholder="Введите email"
            />
            <button
                type='submit'
                onClick={submitHandler}
            >
                Зарегистрироваться
            </button>
            <LoginQuestion onLogin={onLogin}/>
            {/*{error && <ErrorMessage error={error}/>}*/}
        </form>
    )
}