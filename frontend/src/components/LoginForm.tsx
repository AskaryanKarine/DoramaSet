import React, {useState} from "react";
import {RegistrationQuestion} from "./RegistrationQuestion";

interface LoginFormProps {
    onRegistration: () => void
}

export function LoginForm({onRegistration}:LoginFormProps) {
    const [valueLogin, setValueLogin] = useState('')
    const [valuePassword, setValuePassword] = useState('')

    const [error, setError] = useState('')

    const submitHandler = async (event: React.FormEvent) => {
        setError('')
        event.preventDefault()

        if (valueLogin.trim().length === 0 || valuePassword.trim().length === 0) {
            setError('Please enter valid title')
        }
    }

    return (
        <form onSubmit={submitHandler}>
            <input
                type="text"
                placeholder='Введите логин'
                value={valueLogin}
                onChange={(event)=>{setValueLogin(event.target.value)}}
            />
            <input
                type="password"
                placeholder="Введите пароль"
                value={valuePassword}
                onChange={(event)=>{setValuePassword(event.target.value)}}
            />
            <button
                type='submit'
                onClick={submitHandler}
            >
                Войти
            </button>
            <RegistrationQuestion onRegistration={onRegistration}/>

            {/*{error && <ErrorMessage error={error}/>}*/}
        </form>
    )
}