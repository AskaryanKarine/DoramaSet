import {useDorama} from "../hooks/dorama";
import {DoramaPreview} from "../components/Dorama/DoramaPreview/DoramaPreview";
import {useState} from "react";
import {Modal} from "../components/Modal/Modal";
import {CreateForm} from "../components/Dorama/Form/CreateForm";
import {useAppSelector} from "../hooks/redux";
import {IDorama} from "../models/IDorama";

export function DoramaPage() {
    const {user} = useAppSelector(state => state.userReducer)
    const {dorama, doramaErr, loading, addDorama} = useDorama()
    const [modalVisible, setModalVisible] = useState(false)

    const createHandler = (dorama: IDorama) => {
        setModalVisible(false)
        addDorama(dorama)
    }

    return (
        <>
            <h1>Дорамы</h1>
            <div className="grid grid-cols-2">
                {[...dorama].map(
                    drm => <DoramaPreview dorama={drm} key={drm.id}/>
                )}
            </div>

            {user.isAdmin &&
                <button
                    className="w-auto h-auto border-0 fixed bottom-5 right-5"
                    onClick={()=>{setModalVisible(true)}}
                >
                    <i className="fa-solid fa-plus fa-2x fa-border border-2 rounded-full bg-white border-black"></i>
                </button>
            }

            {modalVisible &&
                <Modal title="Добавить новую дораму"
                onClose={()=>{setModalVisible(false)}}>
                    <CreateForm onCreate={createHandler}/>
                </Modal>}
        </>
    )
}