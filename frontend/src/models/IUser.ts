import {ISubscribe} from "./ISubscribe";

export interface IUser {
    username: string
    email: string
    points: number
    isAdmin: boolean
    sub: ISubscribe
    lastSub: string
    regData: string
}