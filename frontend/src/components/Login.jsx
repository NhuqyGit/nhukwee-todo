import { useState } from "react";
import { useAuth } from "../hooks/AuthProvider";
import { AiOutlineLoading3Quarters } from "react-icons/ai";
import "./Login.scss";
import { useTheme } from "../hooks/ThemeProvider";

import {
    FaCheck,
    FaUser,
    FaUserPlus,
    FaSignInAlt,
    FaTimes,
    FaGoogle,
} from "react-icons/fa";

const Login = () => {
    const common = useTheme();
    const [input, setInput] = useState({
        username: "",
        password: "",
    });

    const [registerInput, setRegInput] = useState({
        username: "",
        email: "",
        password: "",
        verifyPassword: "",
    });

    const auth = useAuth();
    const handleSubmitEvent = (e) => {
        e.preventDefault();
        if (input.username !== "" && input.password !== "") {
            //dispatch action from hooks
            auth.loginAction(input);
            return;
        }
        alert("please provide a valid input");
    };

    const handleInput = (e) => {
        const { name, value } = e.target;
        setInput((prev) => ({
            ...prev,
            [name]: value,
        }));
    };

    const handleRegInput = (e) => {
        const { name, value } = e.target;
        setRegInput((prev) => ({
            ...prev,
            [name]: value,
        }));
    };

    return (
        <div className={`login-container ${common.themeVar}`}>
            <form className="form-register" onSubmit={handleSubmitEvent}>
                <div className="title">
                    <FaUserPlus size={20} />
                    register
                </div>
                <div className="form_control">
                    <input
                        type="text"
                        id="username"
                        name="username"
                        placeholder="username"
                        onChange={handleRegInput}
                        disabled={auth.isLoading}
                    />
                    <div className="input-status">
                        <div>
                            <FaCheck color="var(--dark-primary)" />
                        </div>
                        <div className="">
                            <FaUser color="var()" />
                        </div>
                        <div className="">
                            <FaTimes color="var(--error-color)" />
                        </div>
                    </div>
                </div>
                <div className="form_control">
                    <input
                        type="text"
                        id="email"
                        name="email"
                        placeholder="email"
                        onChange={handleRegInput}
                        disabled={auth.isLoading}
                    />
                    <div className="input-status">
                        <div>
                            <FaCheck color="var(--dark-primary)" />
                        </div>
                        <div className="">
                            <FaTimes color="var(--error-color)" />
                        </div>
                    </div>
                </div>
                <div className="form_control">
                    <input
                        type="password"
                        id="password"
                        name="password"
                        placeholder="password"
                        aria-describedby="user-password"
                        aria-invalid="false"
                        onChange={handleRegInput}
                        disabled={auth.isLoading}
                    />
                    <div className="input-status">
                        <div>
                            <FaCheck color="var(--dark-primary)" />
                        </div>
                        <div className="">
                            <FaTimes color="var(--error-color)" />
                        </div>
                    </div>
                </div>
                <div className="form_control">
                    <input
                        type="password"
                        id="verify-password"
                        name="verify-password"
                        placeholder="verify password"
                        aria-describedby="user-password"
                        aria-invalid="false"
                        onChange={handleRegInput}
                        disabled={auth.isLoading}
                    />
                    <div className="input-status">
                        <div>
                            <FaCheck color="var(--dark-primary)" />
                        </div>
                        <div className="">
                            <FaTimes color="var(--error-color)" />
                        </div>
                    </div>
                </div>
                <button className="btn-submit" disabled={auth.isLoading}>
                    <FaUserPlus size={20} />
                    sign up
                </button>
            </form>
            <form className="form-login" onSubmit={handleSubmitEvent}>
                <div className="title">
                    <FaSignInAlt size={20} />
                    login
                </div>
                <div className="icon-gg">
                    <FaGoogle size={18} />
                </div>
                <div className="line-or">
                    <div></div>
                    or
                    <div></div>
                </div>
                <div className="form_control">
                    <input
                        type="text"
                        id="email"
                        name="email"
                        placeholder="email"
                        onChange={handleInput}
                        disabled={auth.isLoading}
                    />
                </div>
                <div className="form_control">
                    <input
                        type="password"
                        id="password"
                        name="password"
                        placeholder="password"
                        aria-describedby="user-password"
                        aria-invalid="false"
                        onChange={handleInput}
                        disabled={auth.isLoading}
                    />
                </div>
                <div className="input-checkbox">
                    <input type="checkbox" />
                    remember me
                </div>
                <button className="btn-submit" disabled={true}>
                    <FaSignInAlt size={20} />
                    sign in
                </button>
            </form>
            {auth.isLoading ? (
                <div className="loading">
                    <AiOutlineLoading3Quarters
                        size={50}
                        color="white"
                        className="spinner-icon"
                    />
                </div>
            ) : (
                <></>
            )}
        </div>
    );
};

export default Login;
