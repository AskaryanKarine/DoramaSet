import {IStaff} from "../../../models/IStaff";
import {upToFirst} from "../../../hooks/upToFirst";
import styles from  "./StaffShort.module.css"
import {IPhoto} from "../../../models/IPhoto";
import {instance} from "../../../http-common";
import {useStaff} from "../../../hooks/staff";


interface StaffShortProps {
    staff: IStaff
    isEdit: boolean
    idDorama?:number
    addStaff?:(staff:IStaff) => void
}

export function StaffShort({staff, isEdit, idDorama, addStaff}:StaffShortProps) {
    const onCreate = async () => {
        const url = ["/dorama", idDorama, "staff"].join("/")
        await instance.post<void>(url, {
            id: staff.id
        }).then(
            _ => {addStaff && addStaff(staff)}
        )
    }

    return (
        <>
            <div className={styles.container}>
                {staff.photo ?
                    <img src={staff.photo[0].url} alt={staff.name} height={100} width={100}/>
                    : <i className="fa-regular fa-image fa-xl" style={{color: "#787d87"}}></i>}
                <p>{upToFirst(staff.type)}</p>
                <p>{staff.name}</p>
                {isEdit &&
                    <button onClick={onCreate}>Добавить в дораму</button>}
            </div>
        </>
    )
}