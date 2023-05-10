import React from "react";

interface LoginQuestionProps {
    onLogin: () => void
}

export function LoginQuestion({onLogin}:LoginQuestionProps) {
    return (
        <div className="flex center items-center justify-center mt-2">
            <p className='pr-1'>Есть аккаунт?</p>
            <p
                onClick={onLogin}
                className="regP"
            >
                Войдите в систему!
            </p>
        </div>
    )
}