import React from "react";
import { FaUser } from "react-icons/fa6";
import { FaSignOutAlt, FaCog, FaBell } from "react-icons/fa";
import "./User.scss";
import { useAuth } from "../hooks/AuthProvider";
import { useNavigate } from "react-router-dom";

const User = () => {
    const user = useAuth();
    const navigate = useNavigate();

    return (
        <div className="user-container">
            <div className="btn-user" onClick={() => navigate("/login")}>
                <FaUser color="white" />
            </div>

            {/* Dropdown menu */}
            {user.token ? (
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
            ) : null}
        </div>
    );
};

export default User;
