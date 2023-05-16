import {ListPreview} from "../components/Collection/Preview/ListPreview";
import {AddButton} from "../components/Admin/AddButton/AddButton";
import {Modal} from "../components/Modal/Modal";
import {ListForm} from "../components/Collection/Form/ListForm";
import React, {useState} from "react";
import {useCollection} from "../hooks/collection";

export function PublicListPage() {
    const {publicCollection} = useCollection()
    return (
        <>
            <h1>Публичные коллекции</h1>
            <div className="grid grid-cols-3">
                {publicCollection && [...publicCollection].map(
                    lst => <ListPreview list={lst} key={lst.id} isPublic={true}/>
                )}
            </div>
        </>
    )
}