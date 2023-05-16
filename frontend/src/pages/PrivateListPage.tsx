import {DoramaPreview} from "../components/Dorama/DoramaPreview/DoramaPreview";
import React, {useState} from "react";
import {useCollection} from "../hooks/collection";
import {ListPreview} from "../components/Collection/Preview/ListPreview";
import {AddButton} from "../components/Admin/AddButton/AddButton";
import {Modal} from "../components/Modal/Modal";
import {DoramaForm} from "../components/Dorama/Form/DoramaForm";
import {ListForm} from "../components/Collection/Form/ListForm";
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
                    <ListForm onCreate={createHandler}/>
                </Modal>}
        </>
    )
}