import {IStaff} from "../../../models/IStaff";
import {toUpperFirst} from "../../../hooks/toUpperFirst";
import styles from  "./StaffShort.module.css"


interface StaffShortProps {
    staff: IStaff
}

export function StaffShort({staff}:StaffShortProps) {
    return (
        <>
            <div className={styles.container}>
                {staff.photo ?
                    <img src={staff.photo[0].url} alt={staff.name} height={100} width={100}/>
                    : <i className="fa-regular fa-image fa-xl" style={{color: "#787d87"}}></i>}
                <p>{toUpperFirst(staff.type)}</p>
                <p>{staff.name}</p>
            </div>
        </>
    )
}