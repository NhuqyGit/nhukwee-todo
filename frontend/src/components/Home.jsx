import React from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../hooks/AuthProvider";
import Header from "./Header";
import "./Home.scss";

const Home = () => {
    const common = useAuth();
    const navigate = useNavigate();
    return (
        <div className="home-container">
            <Header />
            <button onClick={() => navigate("/login")}>Login</button>
            <h1>Welcome to the Home Page</h1>
            <p>This is the main landing page.</p>
        </div>
    );
};

export default Home;
