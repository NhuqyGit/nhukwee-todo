import { useState } from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import "./App.css";

import Login from "./Login";
import AuthProvider from "../hooks/AuthProvider";
import PrivateRoute from "./PrivateRoute";
import DashBoard from "./DashBoard";
import Home from "./Home";

function App() {
    const [theme, setTheme] = useState(false);
    const changeTheme = (t) => {
        setTheme(t);
    };
    return (
        <div
            className="App"
            style={{
                background: `${theme === true ? "white" : "var(--body-color)"}`,
            }}
        >
            <Router>
                <AuthProvider>
                    <Routes>
                        <Route path="/" element={<Home theme={theme} />} />
                        <Route path="/login" element={<Login />} />
                        <Route element={<PrivateRoute />}>
                            <Route
                                path="/dashboard"
                                element={
                                    <DashBoard
                                        theme={theme}
                                        changeTheme={changeTheme}
                                    />
                                }
                            />
                        </Route>
                        {/* Other routes */}
                    </Routes>
                </AuthProvider>
            </Router>
            {/* <Header theme={theme}/>
            <Todo changeTheme={changeTheme}/>; */}
        </div>
    );
}

export default App;
