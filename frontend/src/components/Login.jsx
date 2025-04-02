import { useState } from "react";
import { useAuth } from "../hooks/AuthProvider";
import { AiOutlineLoading3Quarters } from "react-icons/ai";
import "./Login.scss";

const Login = () => {
    const [input, setInput] = useState({
        username: "",
        password: "",
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

    return (
        <div className="login-container">
            <form onSubmit={handleSubmitEvent}>
                {/* <div className="form_control">
                    <label htmlFor="user-email">Email:</label>
                    <input
                        type="email"
                        id="user-email"
                        name="email"
                        placeholder="example@yahoo.com"
                        aria-describedby="user-email"
                        aria-invalid="false"
                        onChange={handleInput}
                    />
                    <div id="user-email" className="sr-only">
                        Please enter a valid username. It must contain at least 6
                        characters.
                    </div>
                </div> */}
                <div className="form_control">
                    <label htmlFor="user-email">Username:</label>
                    <input
                        type="text"
                        id="username"
                        name="username"
                        placeholder="username"
                        onChange={handleInput}
                        disabled={auth.isLoading}
                    />
                    <div id="user-email" className="sr-only">
                        Please enter a valid username. It must contain at least
                        6 characters.
                    </div>
                </div>
                <div className="form_control">
                    <label htmlFor="password">Password:</label>
                    <input
                        type="password"
                        id="password"
                        name="password"
                        aria-describedby="user-password"
                        aria-invalid="false"
                        onChange={handleInput}
                        disabled={auth.isLoading}
                    />
                    <div id="user-password" className="sr-only">
                        your password should be more than 6 character
                    </div>
                </div>
                <button className="btn-submit" disabled={auth.isLoading}>
                    Submit
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
