import React, {useState} from 'react';
import './styles.css'
import {useAppSelector} from "../hooks/redux";
import {Modal} from "./Modal";
import {LoginForm} from "./LoginForm";
import {log} from "util";
import {RegistrationQuestion} from "./RegistrationQuestion";
import {RegistrationForm} from "./RegistrationForm";

export function Login() {
    const [modal, setModal] = useState(false)
    const {token, user} = useAppSelector(state => state.userReducer)
    const [login, setLogin] = useState(true)

    const modalTitle: string = login ? "Вход" : "Регистрация"

    const closeModalHandler = () => {
        setModal(false)
        setLogin(true)
    }

    const escFunc = (event: KeyboardEvent) => {
        if (event.key === "Escape") {
            setModal(false)
        }
    }
    document.addEventListener("keypress", escFunc)

    const loginHandler = () => {
        if (token.length === 0) {
            setModal(true)
        }
    }


    return (
        <div className='userHeader'>
            {token.length > 0 &&
                <button>
                    {user.username}
                </button>
            }
            <button onClick={loginHandler}>
                {token.length == 0 ? 'Войти' : 'Выйти'}
            </button>
            {modal &&
                <Modal title={modalTitle} onClose={closeModalHandler}>
                    {token.length === 0 && login &&
                        <LoginForm onRegistration={()=>setLogin(false)}/>
                    }
                    {token.length === 0 && !login &&
                        <RegistrationForm onLogin={()=>setLogin(true)}/>
                    }
                </Modal>
            }
        </div>
    )
}