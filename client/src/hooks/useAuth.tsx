import React, { createContext, useContext, useState, ReactNode, useEffect } from "react";

interface AuthContextType {
    user: boolean;
    login: (token: string) => void;
    logout: () => void;
}

const AuthContext = createContext<AuthContextType | null>(null);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
    const [user, setUser] = useState<boolean>(false);

    useEffect(() => {
        const token = localStorage.getItem("token");
        if (token) {
            setUser(true);
        }
    }, []);

    const login = (token: string) => {
        localStorage.setItem("token", token);
        setUser(true);
    };

    const logout = () => {
        localStorage.removeItem("token");
        setUser(false);
    };

    return (
        <AuthContext.Provider value={{ user, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (!context) throw new Error("useAuth must be used within AuthProvider");
    return context;
};