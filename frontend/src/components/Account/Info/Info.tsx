import styles from "./Info.module.css";
import {useAppSelector} from "../../../hooks/redux";

export function Info() {
    const { user} = useAppSelector(state => state.userReducer)
    const term = user.sub && user.sub.cost > 0 ? ["до", user.lastSub].join(" ") : ""

    return (
            <div className={styles.rightContainer}>
                <h2>Информация</h2>
                <p>Подписка {user.sub && user.sub.name} {term} </p>
                <p>Почта: {user.email}</p>
                <p>Дата регистрации: {user.regData}</p>
            </div>
    )
}