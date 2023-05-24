import {IReview} from "../../../models/IReview";
import React, {useState} from "react";
import {Emoji} from "emoji-picker-react";
import styles from "./Review.module.css"
import {useAppSelector} from "../../../hooks/redux";
import {instance} from "../../../http-common";

interface ReviewPreviewProps {
    review:IReview
    id?:number
    onDelete:(review:IReview)=>void
}

export function ReviewPreview({review, id, onDelete}:ReviewPreviewProps) {
    const colorStatus = review.access_level >= 3 ? review.username_color : "#000000"
    const [color, setCurColor] = useState(colorStatus);
    const {user} = useAppSelector(state => state.userReducer)

    const iconStyle = "fa-solid fa-user fa-3x fa-border border-0 " +
        "rounded-full bg-white border-black"

    const deleteHandler = async () => {
        await instance.delete<void>(`/dorama/${id}/review`, {
            params: {
                username: user.username,
            },
        })
        onDelete(review)
    }


    return (<>{(review.content && review.content.length > 0 || review.username===user.username) &&
        <div className={styles.container}>
            <div className="flex items-center w-[100%]">
                <span style={{color: color, border:3, borderColor: color, borderStyle:"solid", borderRadius:999999999, marginRight: 10}}>
                        <i className={iconStyle}></i>
                    </span>
                <div className="flex w-[100%] items-center">
                    <div>
                        <div className="flex items-center">
                            <p className="text-lg">{review.username}</p>
                            {review.access_level >= 2 &&
                                <span className="ml-1">
                                <Emoji unified={review.username_emoji} size={25}/>
                            </span>}
                        </div>
                        <p>{[...Array(review.mark)].map((item, index) =>
                            <i className="fa-solid fa-star" style={{color: "#ffc800"}}></i>)}
                        </p>
                    </div>
                    {review.username === user.username &&
                        <button className="ml-auto h-auto w-auto border-0 mr-5"
                        onClick={deleteHandler}>
                            <i className="fa-solid fa-trash"></i>
                        </button>
                    }
                </div>
            </div>
            <div className="mb-2 text-xl">
                <p className="font-bold mt-2">Отзыв: </p>
                <p>{review.content}</p>
            </div>
        </div>}
    </>)
}