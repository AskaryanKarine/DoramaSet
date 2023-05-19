import {IDorama} from "../../../models/IDorama";
import React, {useState} from "react";
import {instance} from "../../../http-common";
import {useDorama} from "../../../hooks/dorama";
import {errorHandler} from "../../../hooks/errorHandler";

interface DoramaFormProps {
    isEdit:boolean
    dorama?:IDorama
    onCreate :(dorama:IDorama)=>void
}

export function DoramaCreate({isEdit, dorama, onCreate}:DoramaFormProps) {
    const [name, setName] = useState(dorama ? dorama.name : "")
    const [description, setDescription] = useState(dorama ? dorama.description : "")
    const [genre, setGenre] = useState(dorama ? dorama.genre : "")
    const [year, setYear] = useState(dorama ? dorama.release_year.toString() : "")
    const [status, setStatus] = useState(dorama ? dorama.status : "in progress")
    const [error, setError] = useState("")

    const submitHandler = async (event: React.FormEvent) => {
        setError('')
        event.preventDefault()

        if (name.trim().length === 0 || description.trim().length === 0 ||
            genre.trim().length === 0 || year.trim().length === 0 ||
            isNaN(parseInt(year.trim()))) {
            setError('Пожалуйста, введите корректные данные')
        }
        const request: IDorama = {
            id: dorama?.id,
            name: name,
            description: description,
            genre: genre,
            status: status,
            release_year: parseInt(year),
        }

        try {
            if (isEdit) {
                await instance.put('/dorama/', request)
                    .then(_ => {onCreate(request)})
            } else {
                const response = await instance.post<{data:IDorama}>('/dorama/', request)
                    .then(r => {onCreate(r.data.data)})
            }
        } catch (e:unknown) {
            errorHandler(e)
        }
    }


    return (<>
        <form onSubmit={submitHandler} onChange={()=>{setError("")}}>
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
                    defaultValue={dorama ? dorama.status : "in progress"}
                >
                    <option value="in progress">в процессе</option>
                    <option value="finish">завершена</option>
                    <option value="will released">анонсирована</option>
                </select>
            </div>
            <button type="submit" className="mt-5">{isEdit ? "Обновить информацию" : "Создать"}</button>
        </form>
    </>)
}