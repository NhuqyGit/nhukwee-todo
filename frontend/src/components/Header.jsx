import React from "react";
import DarkBack from "../image/background.jpg";
import LightBack from "../image/background1.png";
import "./Header.css";

const Header = (props) => {
    const theme = props.theme === false ? DarkBack : LightBack;
    return (
        <header>
            <div style={{backgroundImage: `url(${theme})`, transition: "all ease .5s"}} className="background"></div>
        </header>
    );
}

 

export default Header;