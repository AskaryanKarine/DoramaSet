import {DoramaPreview} from "../components/Dorama/DoramaPreview/DoramaPreview";
import React, {useState} from "react";
import {useCollection} from "../hooks/collection";
import {ListPreview} from "../components/Collection/Preview/ListPreview";
import {AddButton} from "../components/Admin/AddButton/AddButton";
import {Modal} from "../components/Modal/Modal";
import {DoramaCreate} from "../components/Dorama/Form/DoramaCreate";
import {ListCreate} from "../components/Collection/Form/ListCreate";
import {IDorama} from "../models/IDorama";
import {IList} from "../models/IList";

export function PrivateListPage() {
    const [modalVisible, setModalVisible] = useState(false)
    const {userCollection, addPrivateList} = useCollection()

    const createHandler = (list: IList) => {
        setModalVisible(false)
        addPrivateList(list)
    }

    return (
        <>
            <h1>Мои коллекции</h1>
            <div className="grid grid-cols-3">
                {[...userCollection].map(
                    lst => <ListPreview list={lst} key={lst.id} isPublic={false}/>
                )}
            </div>
            <AddButton onOpen={()=>{setModalVisible(true)}}/>
            {modalVisible &&
                <Modal title="Добавить новый список"
                       onClose={()=>{setModalVisible(false)}}>
                    <ListCreate onCreate={createHandler}/>
                </Modal>}
        </>
    )
}