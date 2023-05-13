import React from "react";
import styles from "./Modal.module.css"



interface ModalProps {
    children: React.ReactNode
    title: string
    onClose: () => void
}

export function Modal({children, title, onClose}:ModalProps) {
    return (
        <>
            <div
                className='absolute bg-black/50 top-0 right-0 left-0 bottom-0 h-screen flex justify-center items-center'
            >
                <div
                    className='w-[500px] p-5 rounded-2xl bg-white'
                >
                    <div className="relative flex items-center justify-center">
                        <h1 className='text-2xl text-center mb-2'>{title}</h1>
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