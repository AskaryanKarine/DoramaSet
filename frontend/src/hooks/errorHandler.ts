import {AxiosError} from "axios";
import {IError} from "../models/IError";
import {logout} from "../store/reducers/UserSlice";

export function errorHandler(e: unknown) {
    const error = e as AxiosError<IError>;
    if (!error.response) {
        console.log(e)
        throw e;
    }
    if (error.response && error.response.status === 401) {
        logout()
    }
    if (error.response && error.response.status >= 500) {
        console.log(error)
        throw new Error(error.response.data.error)
    }
    console.log(error)
    return error.response.data.error
}