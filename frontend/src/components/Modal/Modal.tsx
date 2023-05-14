import React, {useEffect} from "react";
import styles from "./Modal.module.css"
import {useAppSelector} from "../../hooks/redux";



interface ModalProps {
    children: React.ReactNode
    title: string
    onClose: () => void
}

export function Modal({children, title, onClose}:ModalProps) {
    let {loading, user, isAuth} = useAppSelector(state => state.userReducer)
    const escFunc = (event: KeyboardEvent) => {
        if (event.key === 'Escape') {
            onClose()
        }
    }

    useEffect(() => {
        document.addEventListener("keydown", escFunc, false)
    })

    return (
        <>
            <div className={styles.modal_overlay}>
                <div className={styles.modal}>
                    <div className={styles.modal_header}>
                        <h1 className='text-3xl mb-1.5'>{title}</h1>
                        <button
                            className={styles.close}
                            onClick={onClose}
                        >
                            <i className="fa-solid fa-xmark fa-lg"></i>
                        </button>
                    </div>
                    {children}
                </div>
            </div>
        </>
    )
}