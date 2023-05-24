import {IStaff} from "../../../models/IStaff";
import styles from "./UpdateStaff.module.css";
import SimpleImageSlider from "react-simple-image-slider";
import React, {useState} from "react";
import {StaffForm} from "./StaffForm";
import {IPhoto} from "../../../models/IPhoto";
import {instance} from "../../../http-common";
import {CreatePhoto} from "../../CreatePhoto/CreatePhoto";


interface UpdateStaffProps {
    staff:IStaff
    close: (staff:IStaff)=>void
}

export function UpdateStaff({staff, close}:UpdateStaffProps) {
    const [error, setError] = useState("")
    const [photoVisible, setPhotoVisible] = useState(false)


    const onCreate = async (photo:IPhoto) => {
        setError('')

        const url = ["/staff/", staff.id, "/picture"].join("")
        await instance.post<void>(url, {
            id: photo.id
        })
        if (staff.photo) {
            staff.photo = [...staff.photo, photo]
        }  else {
            staff.photo = [photo]
        }
        setPhotoVisible(false)
    }

    return (<>
        <div className={styles.container}>
            <div className="w-auto mr-5 relative">
                {staff.photo ? <SimpleImageSlider
                    width={300}
                    height={400}
                    autoPlay={true}
                    images={staff.photo}
                    showBullets={false}
                    showNavs={true}
                    navSize={20}
                /> : <i className="fa-regular fa-image fa-5x" style={{color: "#787d87"}}></i>}
                <button className="w-[100%] mt-5" onClick={()=>{setPhotoVisible(true)}}>Добавить постер</button>
            </div>
            <div className={styles.info}>
                <StaffForm staff={staff} isEdit={true} onCreate={close}/>
            </div>
        </div>
        {photoVisible && <
            CreatePhoto
            onClose={()=>{setPhotoVisible(false)}}
            onCreate={onCreate}
        />}
    </>)
}