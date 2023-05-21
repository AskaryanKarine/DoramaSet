import {IDorama} from "../../../models/IDorama";
import styles from "./DoramaUpdate.module.css";
import SimpleImageSlider from "react-simple-image-slider";
import React, {useState} from "react";
import {CreatePhoto} from "../../CreatePhoto/CreatePhoto";
import {IPhoto} from "../../../models/IPhoto";
import {instance} from "../../../http-common";
import {useStaff} from "../../../hooks/staff";
import {EpisodeInfo} from "../Info/EpisodeInfo";
import {StaffInfo} from "../Info/StaffInfo";
import {DoramaCreate} from "./DoramaCreate";
import {errorHandler} from "../../../hooks/errorHandler";
import {IEpisode} from "../../../models/IEpisode";

interface UpdateFormProps {
    dorama:IDorama
    onClose:(dorama:IDorama)=>void
    addEpisode?:(ep:IEpisode)=>void
}

export function DoramaUpdate({dorama, onClose, addEpisode}:UpdateFormProps) {
    const [error, setError] = useState("")
    const [photoVisible, setPhotoVisible] = useState(false)

    const onCreatePhoto = async (photo:IPhoto) => {
        setError('')

        try {
            const url = ["/dorama/", dorama.id, "/picture"].join("")
            await instance.post<void>(url, {
                id: photo.id
            })
            if (dorama.posters) {
                dorama.posters = [...dorama.posters, photo]
            }  else {
                dorama.posters = [photo]
            }
        } catch (e:unknown) {
            errorHandler(e)
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
            <DoramaCreate dorama={dorama} isEdit={true} onCreate={onClose}/>
            </div>
        </div>
        <EpisodeInfo id={dorama.id} isEdit={true} add={addEpisode}/>
        <StaffInfo id={dorama.id} isEdit={true} idDorama={dorama.id}/>
        {photoVisible && <
            CreatePhoto
            onClose={()=>{setPhotoVisible(false)}}
            onCreate={onCreatePhoto}
        />}
    </>)
}