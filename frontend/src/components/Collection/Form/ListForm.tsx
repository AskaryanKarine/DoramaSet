import {IList} from "../../../models/IList";
import React, {useState} from "react";
import {useDorama} from "../../../hooks/dorama";
import {IDorama} from "../../../models/IDorama";
import {instance} from "../../../http-common";
import {useCollection} from "../../../hooks/collection";

interface ListFormInterface {
    list?:IList
    onCreate: (list:IList)=>void
}

interface listResponse {
    data:IList
}

export function ListForm({list, onCreate}:ListFormInterface) {
    
    const [name, setName] = useState("")
    const [description, setDescription] = useState("")
    const [type, setType] = useState("private")
    const [error, setError] = useState("")
    const {} = useCollection()


    const submitHandler = async (event: React.FormEvent) => {
        setError('')
        event.preventDefault()

        if (name.trim().length === 0 || description.trim().length === 0) {
            setError('Пожалуйста, введите корректные данные')
        }
        const request: IList = {
            name: name,
            description: description,
            type: type,
        }

        await instance.post<listResponse>('/list/', request)
            .then(r => {
                r.data.data.type = type
                onCreate(r.data.data)
                console.log(r.data.data)})
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
            <div className="flex w-[100%]">
                <p>Тип: </p>
                <select
                    form="dorama-form"
                    required={true}
                    className="w-auto"
                    onChange={(e)=>{setType(e.target.value)}}
                    defaultValue={"private"}
                >
                    <option value="private">приватный</option>
                    <option value="public">публичный</option>
                </select>
            </div>
            <button type="submit" className="mt-5">Создать</button>
        </form>
    </>)
}
