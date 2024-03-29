import {useAppSelector} from "../hooks/redux";
import React, {useState} from "react";
import {useStaff} from "../hooks/staff";
import {Loading} from "../components/Loading/Loading";
import {StaffPreview} from "../components/Staff/Preview/StaffPreview";
import {Search} from "../components/Search/Search";
import {AddButton} from "../components/Admin/AddButton/AddButton";
import {Modal} from "../components/Modal/Modal";
import {IStaff} from "../models/IStaff";
import {StaffForm} from "../components/Staff/Form/StaffForm";

export function StaffPage() {
    const {user} = useAppSelector(state => state.userReducer)
    const {staff, staffErr, staffLoading, addAllStaff, findStaff, resetStaff} = useStaff()
    const [modalVisible, setModalVisible] = useState(false)

    const createHandler = (staff: IStaff) => {
        setModalVisible(false)
        addAllStaff(staff)
    }
    console.log(staff ? "q" : "b")
    return (
        <>
            {staffLoading && <Loading/>}
            <h1>Стафф</h1>
            <div>
                <Search findFunc={findStaff} resetFunc={resetStaff}/>
            </div>
            <div className="grid grid-cols-2">
                {staff.length > 0 ?[...staff].map(
                    stf => <StaffPreview staff={stf} key={stf.id}/>
                ) : <p className="relative block text-center mt-5 text-xl">
                    Ничего не найдено
                </p>}
            </div>
            {user.isAdmin &&
                <AddButton onOpen={()=>{setModalVisible(true)}}/>
            }
            {modalVisible &&
                <Modal title="Добавить нового стаффа"
                       onClose={()=>{setModalVisible(false)}}>
                    <StaffForm onCreate={createHandler} isEdit={false}/>
                </Modal>}
        </>
    )
}