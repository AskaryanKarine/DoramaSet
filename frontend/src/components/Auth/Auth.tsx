import React, {useEffect, useState} from 'react';
import {useAppDispatch, useAppSelector} from "../../hooks/redux";
import {Modal} from "../Modal/Modal";
import styles from './Auth.module.css'
import {AuthState, getAuthTitle} from "../../models/AuthState";
import {AuthForm} from "./Form/AuthForm";
import {logout, resetError, resetState} from "../../store/reducers/UserSlice";
// import {resetError, resetState} from "../../store/reducers/UserSlice";


export function Auth() {
    const dispatch = useAppDispatch()
    const {isAuth} = useAppSelector(state => state.userReducer)

    const [modalVisibility, setModalVisibility] = useState(false)
    const [authState, setAuthState] = useState<AuthState>(AuthState.SIGN_IN)

    const modalTitle = getAuthTitle(authState)
    const btnTitle: string = !isAuth ? 'Войти' : 'Выйти'
    const iconName: string = !isAuth ? "fa-right-to-bracket" : "fa-right-from-bracket"
    const iconClasses = ['fa pr-1.5 fa-lg', iconName]

    const closeModalHandler = () => {
        setModalVisibility(false)
        setAuthState(AuthState.SIGN_IN)
        dispatch(resetError())
    }

    const loginHandler = () => {
        if (!isAuth) {
            setModalVisibility(true)
        } else {
            dispatch(logout())
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
                    <AuthForm
                     onClose={closeModalHandler}
                     authState={authState}
                     onRegistration={() => {setAuthState(AuthState.SIGN_IN)}}
                     onSignIn={() => {setAuthState(AuthState.REGISTRATION)}}/>
                </Modal>
            }
        </div>
    )
}