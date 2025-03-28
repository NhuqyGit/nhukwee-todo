import React from "react";
import Header from "./Header";
import Todo from "./Todo";

const DashBoard = ({ theme, changeTheme }) => {
    return (
        <>
            <Header theme={theme} />
            <Todo changeTheme={changeTheme} />;
        </>
    );
};

export default DashBoard;
