import React from "react";
import { useNavigate } from "react-router-dom";

const MarketPage = () => {
  const navigate = useNavigate();

  const handleViewStocks = () => {
    navigate("/stocks");
  };

  return (
    <div className="p-4 flex flex-col items-center justify-center min-h-screen">
      <h1 className="text-4xl font-bold mb-6">Welcome to the Market</h1>
      <p className="text-lg text-gray-600 mb-6">
        Explore popular stocks and track their performance.
      </p>
      <button
        onClick={handleViewStocks}
        className="bg-blue-500 text-white px-6 py-3 rounded-lg shadow hover:bg-blue-600 transition"
      >
        View Popular Stocks
      </button>
    </div>
  );
};

export default MarketPage;
