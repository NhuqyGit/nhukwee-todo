import React, { useState } from "react";
import { FaUser } from "react-icons/fa6";
import { FaSignOutAlt, FaCog, FaBell } from "react-icons/fa";
import "./User.scss";
import { useAuth } from "../hooks/AuthProvider";

const User = () => {
    const [isOpen, setIsOpen] = useState(false);
    const user = useAuth();

    const toggleMenu = () => {
        setIsOpen(!isOpen);
    };

    return (
        <div className="user-container">
            <div className="btn-user" onClick={toggleMenu}>
                <FaUser color="white" />
            </div>

            {/* Dropdown menu */}
            <div className={`dropdown-menu `}>
                <div className="menu-item">
                    <FaCog color="white" />
                </div>
                <div className="menu-item">
                    <FaBell color="white" />
                </div>
                <div className="menu-item" onClick={user.logOut}>
                    <FaSignOutAlt color="white" />
                </div>
            </div>
        </div>
    );
};

export default User;
