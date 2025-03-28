import React from "react";
import { useNavigate } from "react-router-dom";

const Home = ({ theme }) => {
    const navigate = useNavigate();
    return (
        <div
            style={{
                textAlign: "center",
                padding: "50px",
                background: theme ? "#f4f4f4" : "#333",
                color: theme ? "#000" : "#fff",
            }}
        >
            <button onClick={() => navigate("/login")}>Login</button>
            <h1>Welcome to the Home Page</h1>
            <p>This is the main landing page.</p>
        </div>
    );
};

export default Home;
