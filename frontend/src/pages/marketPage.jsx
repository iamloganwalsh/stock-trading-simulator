import React from "react";
import { useNavigate } from "react-router-dom";

const MarketPage = () => {
  const navigate = useNavigate();

  const handleViewStocks = () => {
    navigate("/stocks");
  };

  return (
    <div>
      <h1>Welcome to the Market</h1>
      <p>
        Explore popular stocks and track their performance.
      </p>
      <button
        onClick={handleViewStocks}
        style={{ padding: "10px 20px", fontSize: "16px", color:'#242424', backgroundColor:'lightgrey' }}
      >
        View Popular Stocks
      </button>
    </div>
  );
};

export default MarketPage;
