import {Modal} from "../Modal/Modal";
import React, {useState} from "react";
import {instance} from "../../http-common";
import {IPhoto} from "../../models/IPhoto";
import {createPhoto} from "../../hooks/presenters";

interface CreatePhotoProps {
    onClose : ()=>void
    onCreate: (photo:IPhoto)=>void
}

export function CreatePhoto({onCreate, onClose}:CreatePhotoProps) {
    const [url, setURL] = useState("")
    const [error, setError] = useState("")

    const submitHandler = async (event: React.FormEvent) => {
        setError('')
        event.preventDefault()

        if (url.trim().length === 0) {
            setError('Пожалуйста, введите корректные данные')
            return
        }
        createPhoto(url).then(
            r => {
                if (typeof r === 'string') {setError(r)}
                else {onCreate(r)}
            })
    }

    return (<>
        <Modal
            title={"Добавить фото"}
            onClose={onClose}
        >
            <form onSubmit={submitHandler}>
                <input
                    type="text"
                    placeholder='URL'
                    value={url}
                    onChange={(event)=>{setURL(event.target.value)}}/>
                <button type="submit">Добавить</button>
            </form>
        </Modal>
    </>)
}