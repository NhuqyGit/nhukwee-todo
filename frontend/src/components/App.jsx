import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import "./App.scss";

import Login from "./Login";
import AuthProvider from "../hooks/AuthProvider";
import PrivateRoute from "./PrivateRoute";
import DashBoard from "./DashBoard";
import Home from "./Home";

function App() {
    return (
        <div className="App">
            <Router>
                <AuthProvider>
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/login" element={<Login />} />
                        <Route element={<PrivateRoute />}>
                            <Route path="/dashboard" element={<DashBoard />} />
                        </Route>
                        {/* Other routes */}
                    </Routes>
                </AuthProvider>
            </Router>
        </div>
    );
}

export default App;
