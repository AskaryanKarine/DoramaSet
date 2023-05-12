import React from 'react';
import {Header} from "./components/Header/Header";
import {Route, Routes} from "react-router-dom";
import {SideMenu} from "./components/SideMenu/SideMenu";
import {AccountPages} from "./pages/AccountPages";
import {HomePage} from "./pages/HomePage";
import {TopMenu} from "./components/TopMenu/TopMenu";
import './index.css'


function App() {
    return (
        <div className="app">
            <Header>
                <TopMenu/>
            </Header>
            <SideMenu/>
            <main className="content">
                <Routes>
                    <Route path="/" element={<HomePage/>}/>
                    <Route path="/account" element={<AccountPages />} />
                </Routes>
            </main>
        </div>
    );
}

export default App;
