import React, {useState} from "react";
import {instance} from "../../../http-common";
import {IStaff} from "../../../models/IStaff";
import {useStaff} from "../../../hooks/staff";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import moment from 'moment';
import {errorHandler} from "../../../hooks/errorHandler";

interface StaffFormProps {
    isEdit:boolean
    staff?:IStaff
    onCreate : (staff:IStaff)=>void
}

export function StaffForm({isEdit, staff, onCreate}:StaffFormProps) {
    const [name, setName] = useState(staff ? staff.name : "")
    const [gender, setGender] = useState(staff ? staff.gender : "м")
    const [status, setStatus] = useState(staff ? staff.type : "actor")
    const [error, setError] = useState("")
    const {updateStaff} = useStaff()
    const [bDay, setBDay] = useState(staff ? new Date(staff.birthday) : new Date());

    const submitHandler = async (event: React.FormEvent) => {
        setError('')
        event.preventDefault()

        if (name.trim().length === 0 || gender.trim().length === 0
            || status.trim().length === 0) {
            setError('Пожалуйста, введите корректные данные')
        }
        const request: IStaff = {
            id: staff?.id,
            name: name,
            gender: gender,
            type: status,
            birthday: (moment(bDay)).format('DD.MM.yyyy').toString(),
        }

        try {
            if (isEdit) {
                await instance.put('/staff/', request)
                    .then(_ => {updateStaff(request)})
            } else {
                await instance.post<IStaff>('/staff/', request)
                    .then(_ => {onCreate(request)})
            }
        } catch (e:unknown) {
            errorHandler(e)
        }
    }


    return (<>
        <form onSubmit={submitHandler} onChange={()=>{setError("")}

        }>
            <input
                type="text"
                placeholder='Имя'
                value={name}
                onChange={(event)=>{setName(event.target.value)}}
            />

            <DatePicker selected={bDay} onChange={(date) => {date && setBDay(date)}}
                        dateFormat="dd/MM/yyyy"/>

            <div className="flex w-[100%]">
                <p>Пол: </p>
                <select
                    form="dorama-form"
                    required={true}
                    className="w-auto"
                    onChange={(e)=>{setGender(e.target.value)}}
                    defaultValue={staff ? staff.gender : "м"}
                >
                    <option value="м">мужской</option>
                    <option value="ж">женский</option>
                </select>
            </div>

            <div className="flex w-[100%]">
                <p>Роль: </p>
                <select
                    form="dorama-form"
                    required={true}
                    className="w-auto"
                    onChange={(e)=>{setStatus(e.target.value)}}
                    defaultValue={staff ? staff.type : "actor"}
                >
                    <option value="actor">актер</option>
                    <option value="director">режиссёр</option>
                    <option value="screenwriter">сценарист</option>
                </select>
            </div>
            <button type="submit" className="mt-5">{isEdit ? "Обновить информацию" : "Создать"}</button>
        </form>
    </>)
}