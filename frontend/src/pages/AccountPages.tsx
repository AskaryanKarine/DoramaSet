import {useAppSelector} from "../hooks/redux";

export function AccountPages() {
    const {user} = useAppSelector(state => state.userReducer)
    return (
        <>
            <div className="justify-center flex">
                <div className="flex items-center ">
                    <i className="fa-solid fa-user fa-5x fa-border border-black border-solid border-2 rounded-full"></i>
                    <h1 className='text-5xl font-bold pl-4'>{user.username}</h1>
                </div>
            </div>
            <div>
                <h2 className='text-3xl mt-2.5'>Информация о пользователе</h2>
                <p>Почта: {user.email}</p>
                <p>Роль: {user.isAdmin ? "Администратор" : "Пользователь"}</p>
                <p>Баллы: {user.points}</p>
                <p>Дата регистрации: {user.regData}</p>
                <p>Информация о подписке:</p>
                {/*<p>Описание: {user.sub.duration }</p>*/}
                {/*<p>Длительность: {user.sub.duration}</p>*/}
            </div>
        </>

    )
}