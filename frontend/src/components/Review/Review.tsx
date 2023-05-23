import {IReview} from "../../models/IReview";
import styles from "../Dorama/Dorama.module.css";
import React, {useState} from "react";
import {ReviewPreview} from "./ReviewPreview/ReviewPreview";
import {Modal} from "../Modal/Modal";
import {useAppSelector} from "../../hooks/redux";
import {instance} from "../../http-common";
import {IDorama} from "../../models/IDorama";

interface ReviewProps {
    id?:number
    reviews?:IReview[]
    add: (mark:number)=>void
    del: (mark:number)=>void
}

export function Review({reviews, id, add, del}:ReviewProps) {
    const [current, setCurrent] = useState<IReview[]>(reviews ? reviews : [])
    const {user} = useAppSelector(state => state.userReducer)
    const [modal, setModal] = useState(false)
    const [content, setContent] = useState("")
    const [mark, setMark] = useState("5")
    let visButton = true
    for (let i = 0; i < current.length; i++) {
        if (current[i].username === user.username) {
            visButton = false
            break
        }
    }

    const submitHandler = async (event: React.FormEvent) => {
        event.preventDefault()
        const response = await instance.post<{data:IReview}>(`/dorama/${id}/review`, {
            username: user.username,
            mark: parseInt(mark),
            content: content,
        })
        if (current.length > 0) {
            setCurrent(prev => [...prev, response.data.data])
        } else {
            setCurrent([response.data.data])
        }
        add(parseInt(mark))
        setModal(false)
    }

    function deleteReview(review:IReview) {
        if (current) {
            const index = current.indexOf(review)
            if (index !== -1) {
                setCurrent(current.filter((_, i) => i !== index))
            }
            del(review.mark)
        }
    }
    return (<>
        <div>
            <div className={styles.addHeader}>
                <h2>Отзывы</h2>
                {visButton && <button onClick={()=>{setModal(true)}}>
                    <i className="fa-solid fa-plus fa-border border-2 rounded-full bg-white border-black"></i>
                </button>}
            </div>
            {current && current.length > 0 ? [...current].map((rev, index) =>
                <ReviewPreview
                    key={index}
                    id={id}
                    review={rev}
                    onDelete={deleteReview}/>
            ) : <p>
                Еще нет отзывов. Оставьте первым!
            </p>}
        </div>
        {modal && <Modal title="Добавить новый отзыв" onClose={()=>{setModal(false)}}>
            <form onSubmit={submitHandler} className="mt-0">
                <div className="flex w-[100%]">
                    <p>Оценка: </p>
                    <select
                        required={true}
                        className="w-auto"
                        onChange={(e)=>{setMark(e.target.value)}}
                        defaultValue={mark}
                    >
                        <option value={1}>1</option>
                        <option value={2}>2</option>
                        <option value={3}>3</option>
                        <option value={4}>4</option>
                        <option value={5}>5</option>
                    </select>
                </div>
                <textarea
                    rows={3}
                    placeholder='Отзыв (не обязательно)'
                    value={content}
                    onChange={(event)=>{setContent(event.target.value)}}
                />
                <button type="submit">Добавить</button>
            </form>
        </Modal>}
    </>)
}