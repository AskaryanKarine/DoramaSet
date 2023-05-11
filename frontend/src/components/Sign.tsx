import React, {useState} from 'react';
import './styles.css'
import {useAppSelector} from "../hooks/redux";
import {Modal} from "./Modal";
import {SignForm} from "./SignForm";
import {Link} from "react-router-dom";
import CircumIcon from "@klarr-agency/circum-icons-react";

export function Sign() {
    const [modal, setModal] = useState(false)
    const {token, user} = useAppSelector(state => state.userReducer)
    const [login, setLogin] = useState(true)

    const modalTitle: string = login ? "Вход" : "Регистрация"
    const iconName: string = !login ? "fa-right-to-bracket" : "fa-right-from-bracket"
    const iconClases = ['fa pr-1 fa-xl', iconName]

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
        <div className='signHeader'>
            {token.length > 0 &&
                <Link to="/account">
                    <button className="w-[110px] h-[40px]">
                        <i className="fa-regular fa-user pr-1 fa-lg"></i>
                        {user.username}
                    </button>
                </Link>
            }

            <button onClick={loginHandler}>
                <i className={iconClases.join( " ")} ></i>
                {token.length === 0 ? 'Войти' : 'Выйти'}

            </button>
            {modal &&
                <Modal title={modalTitle} onClose={closeModalHandler}>
                    {token.length === 0 &&
                        <SignForm
                            onRegistration={()=>setLogin(false)}
                            onLogin={()=>setLogin(true)}
                            login={login}
                            onClose={closeModalHandler}
                        />
                    }
                </Modal>
            }
        </div>
    )
}