import styles from "./DoramaShort.module.css"
import {IDorama} from "../../../models/IDorama";
import {instance} from "../../../http-common";
import {upToFirst} from "../../../hooks/upToFirst";


interface DoramaShortProps {
    dorama: IDorama
    isEdit: boolean
    idList?:number
}

export function DoramaShort({dorama, isEdit, idList}:DoramaShortProps) {
    const onCreate = async () => {
        const url = ["/list/", idList].join("")
        console.log(url)
        await instance.delete<void>(url, {
            params: {
                id: dorama.id
            }
        }).then(_ => {})
    }

    return (
        <>
            <div className={styles.container}>
                {dorama.posters ?
                    <img src={dorama.posters[0].url} alt={dorama.name} height={100} width={100}/>
                    : <i className="fa-regular fa-image fa-xl" style={{color: "#787d87"}}></i>}
                <div className="p-2 m-1 w-[50%]">
                    <p>{dorama.name}</p>
                    <p>{upToFirst(dorama.genre)}</p>
                    <p>{dorama.release_year}</p>
                    <p>Серий: {dorama.episodes ?  dorama.episodes.length : 0}, {dorama.status}</p>
                </div>
                {isEdit &&
                    <button onClick={onCreate}>Удалить</button>}
            </div>
        </>
    )
}