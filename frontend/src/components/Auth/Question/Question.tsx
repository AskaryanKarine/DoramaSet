import React from "react";
import {AuthState} from "../../../models/AuthState";
import styles from "./Question.module.css"

interface AuthQuestionProps {
    status: AuthState
    onAction: () => void
}

export function AuthQuestion({status, onAction}:AuthQuestionProps) {
    let question, action :string
    switch (status) {
        case AuthState.REGISTRATION:
            question = "Есть аккаунт?"
            action = "Войдите в систему!"
            break
        case AuthState.SIGN_IN:
            question = "Нет аккаунта?"
            action = "Зарегистрируйте!"
            break
    }
    return (
        <div className="flex center items-center justify-center mt-2">
            <p className='pr-1'>{question}</p>
            <p
                onClick={onAction}
                className={styles.question}
            >
                {action}
            </p>
        </div>
    )
}