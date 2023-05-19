import {useAppSelector} from "../../../hooks/redux";
import React from "react";

interface AddButtonProps {
    onOpen: ()=>void
}

export function AddButton({onOpen}:AddButtonProps) {
    return (<>
            <button
                className="w-auto h-auto border-0 fixed bottom-5 right-5"
                onClick={onOpen}
            >
                <i className="fa-solid fa-plus fa-2x fa-border border-2 rounded-full bg-white border-black"></i>
            </button>
    </>)
}