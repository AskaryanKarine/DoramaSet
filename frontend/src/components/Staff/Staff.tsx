import {IStaff} from "../../models/IStaff";
import styles from "../Dorama/Dorama.module.css";
import {Slider} from "../Slider/Slider";

interface StaffProps {
    staff:IStaff
}

export function Staff({staff}:StaffProps) {
    return (<>{staff && <>
        <div className={styles.container}>
            <Slider photo={staff.photo}/>
            <div className={styles.info}>
                <p>Дата рождения: {staff.birthday}</p>
                <p>Роль: {staff.type}</p>
                <p>Пол: {staff.gender}</p>
            </div>
        </div>
    </>}</>)
}