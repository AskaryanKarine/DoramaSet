import {IPhoto} from "../models/IPhoto";
import {instance} from "../http-common";
import {errorHandler} from "./errorHandler";

export const createPhoto = async (url:string) => {
    try {
        const request: IPhoto = {
            url: url,
        }
        const response = await instance.post<{Data:IPhoto}>('/picture/', request)
        return response.data.Data
    } catch (e:unknown) {
        return errorHandler(e)
    }
}

