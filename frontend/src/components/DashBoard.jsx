import React from "react";
import Header from "./Header";
import Todo from "./Todo";

import { useTheme } from "../hooks/ThemeProvider";
import "./DashBoard.scss";

const DashBoard = () => {
    const common = useTheme();
    return (
        <div
            className="dashboard"
            style={{
                background: `${
                    common.theme === true ? "white" : "var(--body-color)"
                }`,
            }}
        >
            <Header />
            <Todo />;
        </div>
    );
};

export default DashBoard;
