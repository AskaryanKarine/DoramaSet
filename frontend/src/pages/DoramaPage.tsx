import {useDorama} from "../hooks/dorama";
import {DoramaPreview} from "../components/Dorama/DoramaPreview/DoramaPreview";
import React, {useState} from "react";
import {Modal} from "../components/Modal/Modal";
import {useAppSelector} from "../hooks/redux";
import {IDorama} from "../models/IDorama";
import {DoramaCreate} from "../components/Dorama/Form/DoramaCreate";
import {Loading} from "../components/Loading/Loading";
import {Search} from "../components/Search/Search";
import {AddButton} from "../components/Admin/AddButton/AddButton";

export function DoramaPage() {
    const {dorama, doramaErr, loading, addDorama, findDorama, resetDorama} = useDorama()
    const [modalVisible, setModalVisible] = useState(false)
    const {user} = useAppSelector(state => state.userReducer)

    const createHandler = (dorama: IDorama) => {
        setModalVisible(false)
        addDorama(dorama)
    }

    return (
        <>
            {loading && <Loading/>}
            <h1>Дорамы</h1>
            <div>
                <Search findFunc={findDorama} resetFunc={resetDorama}/>
            </div>
            <div className="grid grid-cols-2">
                {dorama ? [...dorama].map(
                    drm => <DoramaPreview dorama={drm} key={drm.id}/>
                ) : "Ничего не найдено"}
            </div>
            {user.isAdmin &&
                <AddButton onOpen={()=>{setModalVisible(true)}}/>
            }
            {modalVisible &&
                <Modal title="Добавить новую дораму"
                onClose={()=>{setModalVisible(false)}}>
                    <DoramaCreate onCreate={createHandler} isEdit={false}/>
                </Modal>}
        </>
    )
}