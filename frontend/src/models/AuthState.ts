export enum AuthState {
    SIGN_IN,
    REGISTRATION,
}

export function getAuthTitle(state: AuthState) {
    switch (state) {
        case AuthState.SIGN_IN:
            return "Вход";
        case AuthState.REGISTRATION:
            return  "Регистрация"
        default:
            return "Вход"
    }
}