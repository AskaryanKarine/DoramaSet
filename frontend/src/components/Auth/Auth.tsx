import React, {useState} from 'react';
import {useAppDispatch, useAppSelector} from "../../hooks/redux";
import {Modal} from "../Modal/Modal";
import styles from './Auth.module.css'
import {AuthState, getAuthTitle} from "../../models/AuthState";
import {AuthForm} from "./Form/AuthForm";
// import {resetError, resetState} from "../../store/reducers/UserSlice";


export function Auth() {
    const dispatch = useAppDispatch()
    const {token} = useAppSelector(state => state.userReducer)

    const [modalVisibility, setModalVisibility] = useState(false)
    const [authState, setAuthState] = useState<AuthState>(AuthState.SIGN_IN)

    const modalTitle = getAuthTitle(authState)
    const btnTitle: string = token.length === 0 ? 'Войти' : 'Выйти'
    const iconName: string = token.length === 0 ? "fa-right-to-bracket" : "fa-right-from-bracket"
    const iconClasses = ['fa pr-1.5 fa-lg', iconName]

    const closeModalHandler = () => {
        setModalVisibility(false)
        setAuthState(AuthState.SIGN_IN)
        // dispatch(resetError())
    }

    const escFunc = (event: KeyboardEvent) => {
        if (event.key === "escape") {
            setModalVisibility(false)
        }
    }
    document.addEventListener("keydown", escFunc)

    const loginHandler = () => {
        if (token.length === 0) {
            setModalVisibility(true)
        } else {
            // dispatch(resetState())
        }
    }

    return (
        <div className={styles.menu}>
            <button onClick={loginHandler}>
                <i className={iconClasses.join( " ")} ></i>
                {btnTitle}
            </button>
            {modalVisibility &&
                <Modal title={modalTitle} onClose={closeModalHandler}>
                    {token.length === 0
                        &&
                        <AuthForm
                         authState={authState}
                         onClose={closeModalHandler}
                         onRegistration={() => {setAuthState(AuthState.SIGN_IN)}}
                         onSignIn={() => {setAuthState(AuthState.REGISTRATION)}}/>
                    }
                </Modal>
            }
        </div>
    )
}