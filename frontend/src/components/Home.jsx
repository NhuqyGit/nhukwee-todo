import React from "react";
import "./Home.scss";
import { useTheme } from "../hooks/ThemeProvider";

const Home = () => {
    const common = useTheme();

    return (
        <div className={`home-container ${common.themeVar}`}>
            <button onClick={common.changeTheme}>Switch</button>
            <h1>Welcome to the Home Page</h1>
            <p>This is the main landing page.</p>
        </div>
    );
};

export default Home;
