import React from "react";

interface RegistrationQuestionProps {
    onRegistration: () => void
}

export function RegistrationQuestion({onRegistration}:RegistrationQuestionProps) {
    return (
        <div className="flex center items-center justify-center mt-2">
            <p className='pr-1'>Нет аккаунта?</p>
            <p
                onClick={onRegistration}
                className="regP"
            >
                Зарегистрируйте!
            </p>
        </div>
    )
}