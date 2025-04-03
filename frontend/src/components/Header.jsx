import React from "react";
// import DarkBack from "../image/background.jpg";
// import LightBack from "../image/background1.png";
import "./Header.scss";
import User from "./User";
import Logo from "../image/logo.svg";
// import { useAuth } from "../hooks/AuthProvider";
// import { useTheme } from "../hooks/ThemeProvider";

const Header = () => {
    // const common = useTheme();
    // const theme = common.theme === false ? DarkBack : LightBack;

    return (
        <header>
            <a id="logo" href="/">
                <img src={Logo} alt="doodoo" />
                <h1 className="text">doodoo</h1>
            </a>
            {/* <div
                style={{
                    backgroundImage: `url(${theme})`,
                    transition: "all ease .5s",
                }}
                className="background"
            ></div> */}
            <User />
        </header>
    );
};

export default Header;
