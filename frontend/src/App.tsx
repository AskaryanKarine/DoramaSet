import React from 'react';
import {Header} from "./components/Header";
import {Login} from "./components/Login";
import {Route, Routes} from "react-router-dom";
import {Menu} from "./components/Menu";

function App() {
    return (
        <div>
            <Header>
                <Login/>
            </Header>
            {/*<p> h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>*/}
            {/*    h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>*/}
            {/*    h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>*/}
            {/*    h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>h <br/>h <br/></p>*/}
            {/*<Routes>*/}
            {/*    /!*<Route path="/"/>*!/*/}
            {/*</Routes>*/}
            <Menu/>
        </div>
    );
}

export default App;
