import {IDorama} from "../models/IDorama";

export function upToFirst(str:string) {
    if (str) {
        return str.charAt(0).toUpperCase() + str.slice(1)
    }
    return str
}