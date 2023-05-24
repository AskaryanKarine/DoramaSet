import {ListPreview} from "../components/Collection/Preview/ListPreview";
import React from "react";
import {useCollection} from "../hooks/collection";
import {IList} from "../models/IList";

export function PublicListPage() {
    const {publicCollection, favCollection} = useCollection()
    function inFav(list: IList) {
        if (favCollection) {
            const index = favCollection.indexOf(list)
            return index === -1
        }
        return false
    }
    return (
        <>
            <h1>Публичные коллекции</h1>
            <div className="grid grid-cols-3">
                {publicCollection && [...publicCollection].map(
                    lst => <ListPreview list={lst} key={lst.id} inFav={inFav(lst)}/>
                )}
            </div>
        </>
    )
}