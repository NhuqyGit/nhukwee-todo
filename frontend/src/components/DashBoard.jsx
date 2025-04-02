import React from "react";
import Header from "./Header";
import Todo from "./Todo";
import { useAuth } from "../hooks/AuthProvider";
import "./DashBoard.scss";

const DashBoard = () => {
    const common = useAuth();
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
