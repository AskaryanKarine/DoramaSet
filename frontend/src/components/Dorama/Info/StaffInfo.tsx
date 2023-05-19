import {StaffShort} from "../../Staff/Short/StaffShort";
import {useStaff} from "../../../hooks/staff";
import {Loading} from "../../Loading/Loading";
import styles from "../Dorama.module.css";
import React, {useState} from "react";
import {Modal} from "../../Modal/Modal";
import {Search} from "../../Search/Search";

interface StaffInfoProps {
    id?:number
    isEdit:boolean
    idDorama?:number
}

export function StaffInfo({id, isEdit, idDorama}:StaffInfoProps) {
    const [modal, setModal] = useState(false)
    const {staffDorama, staffErr, staffLoading, addStaff} = useStaff(id)
    const {findStaff, resetStaff, staff} = useStaff()

    return (<>
        {staffLoading && <Loading/>}
        <div>
            <div className={styles.addHeader}>
                <h2>Стафф</h2>
                {isEdit && <button onClick={()=>{setModal(true)}}>
                    <i className="fa-solid fa-plus fa-border border-2 rounded-full bg-white border-black"></i>
                </button>}
            </div>
            <div className="grid grid-cols-5">
                {staffDorama ? [...staffDorama].map(
                    staff => <StaffShort staff={staff} key={staff.id} isEdit={false}/>
                ) : "Нет стаффа"}
            </div>
        </div>
        {modal && <Modal title={"Выберите стафф"} onClose={()=>{setModal(false)}}>
            <div>
                <Search findFunc={findStaff} resetFunc={resetStaff}/>
                <div className="grid grid-cols-3">
                    {staff ? [...staff].map(
                        staff => <StaffShort staff={staff}
                                             key={staff.id}
                                             isEdit={true}
                                             idDorama={idDorama}
                                             addStaff={addStaff}
                        />
                    ) : "Нет стаффа"}
                </div>
            </div>
        </Modal>}
    </>)
}