import {IDorama} from "./IDorama";

export interface IList {
    id?:number
    name:string
    description:string
    creator_name?:string
    type:string
    doramas?:IDorama[]
}