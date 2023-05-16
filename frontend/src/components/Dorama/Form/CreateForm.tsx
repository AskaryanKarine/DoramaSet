import {AuthState} from "../../../models/AuthState";
import React, {useState} from "react";
import {IDorama} from "../../../models/IDorama";
import {instance} from "../../../http-common";

interface CreateFormProps {
    onCreate: (dorama: IDorama) => void
}

export function CreateForm({onCreate}:CreateFormProps) {
    const [name, setName] = useState('')
    const [description, setDescription] = useState('')
    const [genre, setGenre] = useState('')
    const [year, setYear] = useState('')
    const [error, setError] = useState("")
    const [status, setStatus] = useState("in progress")


    const submitHandler = async (event: React.FormEvent) => {
        setStatus('')
        event.preventDefault()

        if (name.trim().length === 0 || description.trim().length === 0 ||
            genre.trim().length === 0 || year.trim().length === 0) {
            setStatus('Пожалуйста, введите корректные данные')
        }
        const request: IDorama = {
            name: name,
            description: description,
            genre: genre,
            status: status,
            release_year: parseInt(year),
        }

        await instance.post<IDorama>('/dorama/', request)
        onCreate(request)
    }

    return (
        <>
            <form id="dorama-form" onSubmit={submitHandler}>
                <input
                    type="text"
                    placeholder='Название'
                    value={name}
                    onChange={(event)=>{setName(event.target.value)}}
                />
                <textarea
                    rows={5}
                    placeholder="Описание"
                    value={description}
                    onChange={(event) => {setDescription(event.target.value)}}
                />
                <input
                    type="text"
                    placeholder="Жанр"
                    value={genre}
                    onChange={(event)=>{setGenre(event.target.value)}}
                />
                <input
                    type="text"
                    placeholder="Год выхода"
                    value={year}
                    onChange={(event)=>{setYear(event.target.value)}}
                />
                <div className="flex w-[100%]">
                    <p>Статус: </p>
                    <select
                        form="dorama-form"
                        required={true}
                        className="w-auto"
                        onChange={(e)=>{setStatus(e.target.value)}}
                    >
                        <option value="in progress">в процессе</option>
                        <option value="finish">завершена</option>
                        <option value="will released">анонсирована</option>
                    </select>
                </div>
                {/*TODO информационное сообщение*/}
                <p></p>
                <button type="submit" className="mt-2.5">Создать</button>
            </form>
        </>
    )
}