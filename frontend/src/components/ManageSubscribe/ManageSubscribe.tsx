import {Subscribe} from "./Subscribe/Subscribe";
import {useAppSelector} from "../../hooks/redux";
import styles from "./ManageSubscribe.module.css"
import {useSubscribe} from "../../hooks/subscribe";
import {Loading} from "../Loading/Loading";
import {ErrorMessage} from "../ErrorMessage/ErrorMessage";

export function ManageSubscribe() {
    const {error} = useAppSelector(state => state.userReducer)
    const {subscribes, subErr, loading} = useSubscribe()

    const errMsg = subErr.length > 0 ? subErr : error as string

    return (
        <>
            {errMsg && <ErrorMessage error={errMsg}/>}
            {loading && <Loading/>}
            <div className={styles.content}>
                {[...subscribes].map(
                    sub => <Subscribe sub={sub} key={sub.id}/>
                )}
            </div>
        </>

    )
}