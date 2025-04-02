import { useContext, createContext, useState } from "react";
import { useNavigate } from "react-router-dom";
const AuthContext = createContext();

const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(null);
    const [token, setToken] = useState(localStorage.getItem("site") || "");
    const [isLoading, setIsLoading] = useState(false);
    const [theme, setTheme] = useState(false);
    const navigate = useNavigate();

    const changeTheme = () => {
        setTheme(!theme);
    };

    const loginAction = async (data) => {
        try {
            setIsLoading(true);
            const response = await fetch("http://localhost:8080/signin", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(data),
            });

            const res = await response.json();
            console.log(res);

            if (res.data && res?.success) {
                setUser(res.data.user_id);
                setToken(res.data.access_token);
                localStorage.setItem("site", res.data.access_token);
                navigate("/dashboard");
                return;
            }
            throw new Error(res.message);
        } catch (err) {
            console.error(err);
        } finally {
            setIsLoading(false);
        }
    };
    const logOut = () => {
        setUser(null);
        setToken("");
        localStorage.removeItem("site");
        navigate("/login");
    };

    return (
        <AuthContext.Provider
            value={{
                user,
                token,
                loginAction,
                logOut,
                theme,
                changeTheme,
                isLoading,
            }}
        >
            {children}
        </AuthContext.Provider>
    );
};

export default AuthProvider;
export const useAuth = () => {
    return useContext(AuthContext);
};
