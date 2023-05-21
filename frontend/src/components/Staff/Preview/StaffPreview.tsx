import {IStaff} from "../../../models/IStaff";
import styles from "./StaffPreview.module.css";
import {AdminPanel} from "../../Admin/Panel/AdminPanel";
import {useState} from "react";
import {Modal} from "../../Modal/Modal";
import {Dorama} from "../../Dorama/Dorama";
import {DoramaUpdate} from "../../Dorama/Form/DoramaUpdate";
import {UpdateStaff} from "../Form/UpdateStaff";
import {Staff} from "../Staff";
import {IDorama} from "../../../models/IDorama";

interface StaffPreviewProps {
    staff:IStaff
}

export function StaffPreview({staff}:StaffPreviewProps) {
    const [modalVisible, setModalVisible] = useState(false)
    const [editVisible, setEditVisible] = useState(false)
    const [current, setCurrent] = useState(staff)

    const onUpdate = (staff:IStaff) => {
        setCurrent(staff)
        setEditVisible(false)
    }

    return (<>{staff && <>
        <div className={styles.container}>
            <div className="flex justify-center items-center h-full max-w-[150px] w-full">
                {staff.photo ?
                    <img src={staff.photo[0].url} alt={staff.name}></img> :
                    <i className="fa-regular fa-image fa-5x" style={{color: "#787d87"}}></i>
                }
            </div>
            <div className={styles.info}>
                <div>
                    <p className="text-3xl">{current.name}</p>
                    <p>{current.gender}</p>
                    <p>Дата рождения {current.birthday}</p>
                </div>
                <button
                    className={styles.more}
                    onClick={()=>{setModalVisible(true)}}
                >Подробнее</button>
            </div>
            <AdminPanel
                onDelete={()=>{}}
                onEdit={()=>{setEditVisible(true)}}
            />
        </div>
        {modalVisible &&
            <Modal
                title={staff.name}
                onClose={()=>{setModalVisible(false)}}
            >
                <Staff staff={staff}/>
            </Modal>}
        {editVisible &&
            <Modal
                title="Изменить данные о стаффе"
                onClose={() => {setEditVisible(false)
                }}
            >
                <UpdateStaff staff={current} close={onUpdate}/>
            </Modal>}
    </>}</>)
}