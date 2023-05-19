import {ListPreview} from "../components/Collection/Preview/ListPreview";
import React from "react";
import {useCollection} from "../hooks/collection";

export function FavoritePage() {
    const {favCollection} = useCollection()
    return (
        <>
            <h1>Избранное</h1>
            <div className="grid grid-cols-3">
                {favCollection ? [...favCollection].map(
                    lst => <ListPreview list={lst} key={lst.id} isPublic={true}/>
                ) : "В избранном пусто"}
            </div>
        </>
    )
}