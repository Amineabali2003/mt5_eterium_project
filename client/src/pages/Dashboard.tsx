import React, { useState, useEffect } from "react";
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
} from "chart.js";
import API from "services/api";
import { Line } from "react-chartjs-2";

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
);

interface PriceEvolution {
    date: string;
    price: number;
}

interface Transaction {
    hash: string;
    from: string;
    to: string;
    value: string;
    timeStamp: string;
    gasPrice: string;
    gasUsed: string;
}

const isValidData = (data: any): data is PriceEvolution[] => {
    return Array.isArray(data) && data.every(
        (entry) => entry.date && typeof entry.price === "number"
    );
};

const Dashboard = () => {
    const [chartData, setChartData] = useState<PriceEvolution[]>([]);
    const [transactions, setTransactions] = useState<Transaction[]>([]);
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);
    const [currentPage, setCurrentPage] = useState<number>(1);
    const rowsPerPage = 10;

    useEffect(() => {
        const fetchChartData = async () => {
            try {
                setIsLoading(true);
                const response = await API.get("/wallet/get_data");
                if (isValidData(response.data)) {
                    setChartData(response.data);
                    setError(null);
                } else {
                    throw new Error("Invalid data format received from the API.");
                }
            } catch (err: any) {
                setError(err.message || "Failed to fetch data. Please try again later.");
            } finally {
                setIsLoading(false);
            }
        };

        const fetchTransactions = async () => {
            try {
                setIsLoading(true);
                const response = await API.get("/wallet/get_transactions");
                if (response.data?.result && Array.isArray(response.data.result)) {
                    setTransactions(response.data.result);
                } else {
                    throw new Error("Invalid transaction data received from the API.");
                }
            } catch (err: any) {
                setError(err.message || "Failed to fetch transactions. Please try again later.");
            } finally {
                setIsLoading(false);
            }
        };

        fetchChartData();
        fetchTransactions();
    }, []);

    const paginatedTransactions = transactions.slice(
        (currentPage - 1) * rowsPerPage,
        currentPage * rowsPerPage
    );

    const totalPages = Math.ceil(transactions.length / rowsPerPage);

    const chartDisplayData = {
        labels: chartData.map((entry) => entry.date),
        datasets: [
            {
                label: "Crypto Wallet Price Evolution",
                data: chartData.map((entry) => entry.price),
                borderColor: "rgba(75, 192, 192, 1)",
                backgroundColor: "rgba(75, 192, 192, 0.2)",
                tension: 0.4,
            },
        ],
    };

    const chartOptions = {
        responsive: true,
        plugins: {
            legend: { display: true },
            title: { display: true, text: "Crypto Wallet Price Evolution" },
        },
        scales: {
            x: { title: { display: true, text: "Date" } },
            y: { title: { display: true, text: "Price" }, beginAtZero: false },
        },
    };

    return (
        <div style={{ padding: "2rem" }} className="bg-gray-50 min-h-screen">
            <h1 className="text-3xl font-bold text-center mb-8">Crypto Wallet</h1>
            {isLoading && (
                <div className="flex justify-center items-center">
                    <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900"></div>
                </div>
            )}
            {error && <p className="text-red-500 text-center">{error}</p>}
            {!isLoading && !error && chartData.length > 0 && (
                <div
                    role="graphics-document"
                    aria-label="Crypto Wallet Price Evolution Chart"
                    className="max-w-4xl mx-auto bg-white p-6 rounded-lg shadow-lg"
                >
                    <Line data={chartDisplayData} options={chartOptions} />
                </div>
            )}
            {!isLoading && !error && transactions.length > 0 && (
                <div className="max-w-6xl mx-auto mt-8">
                    <div className="overflow-x-auto">
                        <table className="w-full text-left border-collapse bg-white shadow-lg rounded-lg">
                            <thead>
                                <tr className="bg-gray-200">
                                    <th className="p-3 border">Hash</th>
                                    <th className="p-3 border">From</th>
                                    <th className="p-3 border">To</th>
                                    <th className="p-3 border">Value</th>
                                    <th className="p-3 border">Timestamp</th>
                                    <th className="p-3 border">Gas Price</th>
                                    <th className="p-3 border">Gas Used</th>
                                </tr>
                            </thead>
                            <tbody>
                                {paginatedTransactions.map((transaction, index) => (
                                    <tr
                                        key={transaction.hash}
                                        className={index % 2 === 0 ? "bg-gray-100" : "bg-white"}
                                    >
                                        <td className="p-3 border">{transaction.hash}</td>
                                        <td className="p-3 border">{transaction.from}</td>
                                        <td className="p-3 border">{transaction.to}</td>
                                        <td className="p-3 border">{transaction.value}</td>
                                        <td className="p-3 border">{transaction.timeStamp}</td>
                                        <td className="p-3 border">{transaction.gasPrice}</td>
                                        <td className="p-3 border">{transaction.gasUsed}</td>
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                    </div>
                    <div className="flex justify-between items-center mt-4">
                        <button
                            onClick={() => setCurrentPage((prev) => Math.max(prev - 1, 1))}
                            disabled={currentPage === 1}
                            className={`px-4 py-2 bg-gray-800 text-white rounded ${
                                currentPage === 1 ? "opacity-50 cursor-not-allowed" : ""
                            }`}
                        >
                            Previous
                        </button>
                        <span className="text-gray-700">
                            Page {currentPage} of {totalPages}
                        </span>
                        <button
                            onClick={() => setCurrentPage((prev) => Math.min(prev + 1, totalPages))}
                            disabled={currentPage === totalPages}
                            className={`px-4 py-2 bg-gray-800 text-white rounded ${
                                currentPage === totalPages ? "opacity-50 cursor-not-allowed" : ""
                            }`}
                        >
                            Next
                        </button>
                    </div>
                </div>
            )}
            {!isLoading && !error && transactions.length === 0 && (
                <p className="text-center text-gray-500">No transactions available.</p>
            )}
        </div>
    );
};

export default Dashboard;