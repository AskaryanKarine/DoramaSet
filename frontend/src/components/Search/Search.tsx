import React, {useState} from "react";
import {useDorama} from "../../hooks/dorama";
import {useStaff} from "../../hooks/staff";

interface SearchProps {
    findFunc:(name:string)=>void
    resetFunc:()=>void
}

export function Search({findFunc, resetFunc}:SearchProps) {
    const [find, setFind] = useState("")

    const submitHandler = async (event: React.FormEvent) => {
        event.preventDefault()
        if (find.trim().length === 0) {}
        findFunc(find)
    }

    const clickHandler = () => {
        setFind("")
        resetFunc()
    }

    return (<>
        <form className="flex items-center flex-row justify-center"
              onSubmit={submitHandler}>
            <input
                className="w-[90%] m-0"
                type="text"
                placeholder='Поиск'
                value={find}
                onChange={(event)=>{setFind(event.target.value)}}
            />
            <button type="submit" className="w-auto h-auto border-0 ml-2">
                <i className="fa-solid fa-xl fa-magnifying-glass"></i>
            </button>
            <button className="w-auto h-auto border-0 ml-2"
                    onClick={()=>{
                        clickHandler()
                    }}>
                <i className="fa-regular fa-xl fa-circle-xmark"></i>
            </button>

        </form>
    </>)
}