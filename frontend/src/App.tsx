import React from 'react';
import {Header} from "./components/Header";
import {Sign} from "./components/Sign";
import {Route, Routes} from "react-router-dom";
import {Menu} from "./components/Menu";
import {AccountPages} from "./pages/AccountPages";
import {HomePage} from "./pages/HomePage";

function App() {
    return (
        <div className="app">
            <Header>
                <Sign/>
            </Header>
            <Menu/>
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
