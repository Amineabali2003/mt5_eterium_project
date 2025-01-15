import React from "react";
import { Link, useNavigate } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

const Navbar = () => {
    const { user, logout } = useAuth();
    const navigate = useNavigate();

    const handleLogout = () => {
        logout();
        localStorage.removeItem("token");
        navigate("/login");
    };

    return (
        <nav className="bg-gray-800 text-white p-4 flex justify-between">
            <Link to="/" className="font-bold text-lg">My App</Link>
            <div>
                {user ? (
                    <>
                        <Link to="/dashboard" className="mr-4">Dashboard</Link>
                        <Link to="/profile" className="mr-4">Profile</Link>
                        <button onClick={handleLogout} className="bg-red-500 px-4 py-2 rounded">Logout</button>
                    </>
                ) : (
                    <>
                        <Link to="/login" className="mr-4">Login</Link>
                        <Link to="/register" className="bg-blue-500 px-4 py-2 rounded">Register</Link>
                    </>
                )}
            </div>
        </nav>
    );
};

export default Navbar;