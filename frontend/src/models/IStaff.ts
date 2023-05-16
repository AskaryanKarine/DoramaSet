import {IPhoto} from "./IPhoto";

export interface IStaff {
    id?: number
    name: string
    birthday: string
    type: string
    gender: string
    photo?: IPhoto[]
}