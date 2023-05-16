import styles from "../Dorama/DoramaPreview/DoramaPreview.module.css";
import {useAppSelector} from "../../hooks/redux";

interface AdminPanelProps {
    onEdit: ()=>void
    onDelete: ()=>void
}

export function AdminPanel({onEdit, onDelete}:AdminPanelProps) {
    const {user} = useAppSelector(state => state.userReducer)
    return (
        <>
            {user.isAdmin &&
                <div className={styles.adminPanel}>
                    <button className="w-auto h-auto" onClick={onEdit}> {/*редактировать TODO*/}
                        <i className="fa-regular fa-pen-to-square"></i>
                    </button>
                    <button className="w-auto h-auto" onClick={onDelete}> {/*удалить TODO*/}
                        <i className="fa-regular fa-trash-can"></i>
                    </button>
                </div>
            }
        </>
    )
}