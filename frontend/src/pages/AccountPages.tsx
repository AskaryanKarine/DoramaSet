import {Account} from "../components/Account/Account";
import {Info} from "../components/Account/Info/Info";

export function AccountPages() {
    return (
        <>
            <h1>Профиль</h1>
            <div className="flex items-center m-3">
                <Account/>
                <Info/>
            </div>
        </>

    )
}