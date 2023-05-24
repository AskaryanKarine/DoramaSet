import {IPhoto} from "./IPhoto";
import {IEpisode} from "./IEpisode";
import {IReview} from "./IReview";

export interface IDorama {
    id?: number
    name: string
    description: string
    genre: string
    status: string
    release_year: number
    posters?: IPhoto[]
    episodes?: IEpisode[]
    rate?: number
    cnt_rate?: number
    reviews?: IReview[]
}
