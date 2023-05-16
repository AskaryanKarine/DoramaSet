import {IDorama} from "../../../models/IDorama";
import styles from "./UpdateForm.module.css";
import SimpleImageSlider from "react-simple-image-slider";
import React, {useState} from "react";
import {CreatePhoto} from "../../CreatePhoto/CreatePhoto";
import {IPhoto} from "../../../models/IPhoto";
import {instance} from "../../../http-common";
import {useDorama} from "../../../hooks/dorama";
import {useStaff} from "../../../hooks/staff";
import {StaffShort} from "../../Staff/Short/StaffShort";
import {EpisodeInfo} from "../Info/EpisodeInfo";
import {StaffInfo} from "../Info/StaffInfo";
import {DoramaForm} from "./DoramaForm";

interface UpdateFormProps {
    dorama:IDorama
}

export function UpdateForm({dorama}:UpdateFormProps) {
    const [error, setError] = useState("")
    const {staff, staffErr, staffLoading} = useStaff(dorama.id)

    const [photoVisible, setPhotoVisible] = useState(false)

    const onCreate = async (photo:IPhoto) => {
        setError('')

        const url = ["/dorama/", dorama.id, "/picture"].join("")
        await instance.post<void>(url, {
            id: photo.id
        })
        if (dorama.posters) {
            dorama.posters = [...dorama.posters, photo]
        }  else {
            dorama.posters = [photo]
        }
        setPhotoVisible(false)
    }

    return (<>
        <div className={styles.container}>
            <div className="w-auto mr-5 relative">
                {dorama.posters ? <SimpleImageSlider
                    width={300}
                    height={400}
                    autoPlay={true}
                    images={dorama.posters}
                    showBullets={false}
                    showNavs={true}
                    navSize={20}
                /> : <i className="fa-regular fa-image fa-5x" style={{color: "#787d87"}}></i>}
                <button className="w-[100%] mt-5" onClick={()=>{setPhotoVisible(true)}}>Добавить постер</button>
            </div>
            <div className={styles.info}>
            <DoramaForm dorama={dorama} isEdit={true} onCreate={()=>{}}/>
            </div>
        </div>
        <EpisodeInfo id={dorama.id} isEdit={true}/>
        <StaffInfo id={dorama.id} isEdit={true} idDorama={dorama.id}/>

        {photoVisible && <
            CreatePhoto
            onClose={()=>{setPhotoVisible(false)}}
            onCreate={onCreate}
        />}
    </>)
}