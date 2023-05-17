import {IPhoto} from "../models/IPhoto";
import {instance} from "../http-common";
import {errorHandler} from "./errorHandler";

export const createPhoto = async (url:string) => {
    try {
        const request: IPhoto = {
            url: url,
        }
        const response = await instance.post<{data:IPhoto}>('/picture/', request)
        return response.data.data
    } catch (e:unknown) {
        return errorHandler(e)
    }
}

