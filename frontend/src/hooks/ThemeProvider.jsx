import { useContext, createContext, useState } from "react";
const ThemeContext = createContext();

const ThemeProvider = ({ children }) => {
    const [theme, setTheme] = useState(false);

    const changeTheme = () => {
        setTheme(!theme);
    };

    const themeVar = theme ? "light-theme" : "dark-theme";

    return (
        <ThemeContext.Provider
            value={{
                theme,
                themeVar,
                changeTheme,
            }}
        >
            {children}
        </ThemeContext.Provider>
    );
};

export default ThemeProvider;
export const useTheme = () => {
    return useContext(ThemeContext);
};
