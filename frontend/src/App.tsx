import React, {useEffect} from 'react';
import {Header} from "./components/Header/Header";
import {Route, Routes} from "react-router-dom";
import {SideMenu} from "./components/SideMenu/SideMenu";
import {AccountPages} from "./pages/AccountPages";
import {HomePage} from "./pages/HomePage";
import {TopMenu} from "./components/TopMenu/TopMenu";
import './index.css'
import {useAppDispatch, useAppSelector} from "./hooks/redux";
import {getUser} from "./store/reducers/UserSlice";
import {DoramaPage} from "./pages/DoramaPage";
import {StaffPage} from "./pages/StaffPage";
import {PublicListPage} from "./pages/PublicListPage";
import {PrivateListPage} from "./pages/PrivateListPage";
import {FavoritePage} from "./pages/FavoritePage";

function App() {
    const {isAuth, user} = useAppSelector(state => state.userReducer)
    const dispatch = useAppDispatch()
    useEffect(() => {
            dispatch(getUser())
        }, [])
    return (
        <div className="app">
            <Header>
                <TopMenu/>
            </Header>
            <SideMenu/>
            <main className="content">
                <Routes>
                    <Route path="/" element={<PublicListPage/>}/>
                    <Route path="/account" element={<AccountPages />} />
                    <Route path="/list/public" element={<PublicListPage/>}/>
                    <Route path="/dorama" element={<DoramaPage/>}/>
                    <Route path="/staff" element={<StaffPage/>}/>
                    <Route path="/list" element={<PrivateListPage/>}/>
                    <Route path="/list/favorite" element={<FavoritePage/>}/>
                </Routes>
            </main>
        </div>
    );
}

export default App;
