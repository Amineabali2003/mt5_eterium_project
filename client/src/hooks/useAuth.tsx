import React, { createContext, useContext, useState, ReactNode, useEffect } from "react";
import API from "../services/api";

interface AuthContextType {
    user: boolean;
    login: (token: string) => void;
    logout: () => void;
    fetchCurrentUser: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | null>(null);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
    const [user, setUser] = useState<boolean>(false);

    useEffect(() => {
        const token = localStorage.getItem("token");
        if (token) {
            fetchCurrentUser();
        }
    }, []);

    const fetchCurrentUser = async () => {
        try {
            await API.get("/me", {
                headers: { Authorization: `Bearer ${localStorage.getItem("token")}` },
            });
            setUser(true);
        } catch {
            setUser(false);
        }
    };

    const login = (token: string) => {
        localStorage.setItem("token", token);
        fetchCurrentUser();
    };

    const logout = () => {
        localStorage.removeItem("token");
        setUser(false);
    };

    return (
        <AuthContext.Provider value={{ user, login, logout, fetchCurrentUser }}>
            {children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (!context) throw new Error("useAuth must be used within AuthProvider");
    return context;
};