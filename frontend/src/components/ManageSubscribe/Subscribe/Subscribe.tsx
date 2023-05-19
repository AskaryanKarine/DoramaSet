import {ISubscribe} from "../../../models/ISubscribe";
import styles from "./subscribe.module.css"
import {useAppDispatch, useAppSelector} from "../../../hooks/redux";
import {subscribe, unsubscribe} from "../../../store/reducers/UserSlice";

interface SubscribeProps {
    sub: ISubscribe
}

export function Subscribe({sub}:SubscribeProps) {
    const { user} = useAppSelector(state => state.userReducer)

    const dispatch = useAppDispatch()

    const subscribeHandler = async () => {
        if (user.sub && user.sub.id != sub.id) {
            await dispatch(subscribe({idSub: sub.id}))
        } else {
            await dispatch(unsubscribe())
        }
    }

    return (
        <>
        <div className={styles.modal}>
            <h1>{sub.name.toUpperCase()}</h1>
            <p>{sub.description}</p>
            <p>{sub.cost} баллов</p>
            <button
                disabled={sub.cost === 0 && sub.id === user.sub.id}
                onClick={subscribeHandler}
            >
                {user.sub && user.sub.id === sub.id ? "Отменить" : "Оформить"}
            </button>
        </div>
        </>
    )
}