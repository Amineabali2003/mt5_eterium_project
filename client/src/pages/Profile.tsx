import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import API from "../services/api";

interface WalletResponse {
    wallet: string | null;
}

const Profile = () => {
    const [wallet, setWallet] = useState<string>("");
    const [initialWallet, setInitialWallet] = useState<string>("");
    const [isWalletModified, setIsWalletModified] = useState<boolean>(false);
    const [loading, setLoading] = useState<boolean>(false);
    const [submitError, setSubmitError] = useState<string | null>(null);
    const [success, setSuccess] = useState<string | null>(null);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchWallet = async () => {
            setLoading(true);
            try {
                const { data } = await API.get<WalletResponse>("/profile/get_wallet");
                const walletValue = data.wallet || ""; 
                setWallet(walletValue);
                setInitialWallet(walletValue);
            } catch (error) {
                console.error("Failed to fetch wallet");
            } finally {
                setLoading(false);
            }
        };

        fetchWallet();
    }, []);

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setWallet(e.target.value);
        setIsWalletModified(e.target.value !== initialWallet);
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setLoading(true);
        setSubmitError(null);
        setSuccess(null);
        try {
            await API.put("/profile/update_wallet", { wallet });
            setSuccess("Wallet updated successfully!");
            setInitialWallet(wallet);
            setIsWalletModified(false);
            setTimeout(() => {
                navigate("/dashboard");
            }, 1000);
        } catch (error) {
            console.error("Failed to update wallet");
            setSubmitError("Failed to update wallet");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="max-w-md mx-auto p-4 space-y-4">
            <h1 className="text-2xl font-bold">Profile</h1>
            {success && <p className="text-green-500">{success}</p>}
            <form onSubmit={handleSubmit} className="space-y-4">
                <input
                    type="text"
                    placeholder="Wallet"
                    value={wallet}
                    onChange={handleInputChange}
                    className="w-full p-2 border rounded"
                />
                {submitError && <p className="text-red-500">{submitError}</p>}
                <button
                    type="submit"
                    disabled={loading || !isWalletModified}
                    className="w-full bg-blue-500 text-white p-2 rounded disabled:bg-blue-300"
                >
                    {loading ? "Updating..." : "Update Wallet"}
                </button>
            </form>
        </div>
    );
};

export default Profile;