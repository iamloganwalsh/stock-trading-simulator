import React, { useState, useEffect } from "react";
import userServices from "../services/userServices";

const Portfolio = ({ balance, profitloss, loading, error }) => {
  const [portfolio, setPortfolio] = useState({
    totalBalance: balance || 0,
    profitLoss: profitloss || 0,
    stocks: [],
  });

  useEffect(() => {
    // Fetch portfolio data from an API or use dummy data for now
    const fetchStocks = async () => {
      try {
        const stocksData = await userServices.getStockPortfolio();
        setPortfolio((prevPortfolio) => ({
          ...prevPortfolio,
          stocks: stocksData,
        }));
      } catch (err) {
        console.error("Error fetching stocks:", err);
      }
    };
    fetchStocks();
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  return (
    <div style={{ padding: "24px" }}>
      <div
        style={{
          display: "grid",
          gridTemplateColumns: "1fr 1fr",
          gap: "24px",
        }}
      >
        <div
          style={{
            backgroundColor: "#f5f5f5",
            boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
            padding: "16px",
            textAlign: "center",
          }}
        >
          <h2 style={{ fontSize: "20px", fontWeight: "bold" }}>Total Balance</h2>
          <p style={{ fontSize: "24px", color: "#16a34a" }}>${portfolio.totalBalance}</p>
        </div>
        <div
          style={{
            backgroundColor: "#f5f5f5",
            boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
            padding: "16px",
            textAlign: "center",
          }}
        >
          <h2 style={{ fontSize: "20px", fontWeight: "bold" }}>Profit/Loss</h2>
          <p
            style={{
              fontSize: "24px",
              color: portfolio.profitLoss >= 0 ? "#16a34a" : "#dc2626",
            }}
          >
            ${portfolio.profitLoss}
          </p>
        </div>
      </div>

      <div>
        <h2 style={{ fontSize: "24px", fontWeight: "bold", marginBottom: "16px" }}>Your Stocks</h2>
        <div
          style={{
            display: "grid",
            gridTemplateColumns: "repeat(auto-fit, minmax(200px, 1fr))",
            gap: "24px",
            width: "100%", 
            maxWidth: "1200px", 
            margin: "0 auto", 
          }}
        >
          {portfolio.stocks.map((stock, index) => (
            <div
              key={index}
              style={{
                backgroundColor: "#ffffff",
                boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
                padding: "16px",
                textAlign: "center",
                transition: "transform 0.2s",
              }}
              onMouseEnter={(e) => (e.currentTarget.style.transform = "scale(1.05)")}
              onMouseLeave={(e) => (e.currentTarget.style.transform = "scale(1)")}
            >
              <h3 style={{ fontSize: "18px", fontWeight: "bold" }}>{stock.code}</h3>
              <p style={{ color: "#4b5563" }}>Shares: {stock.stock_count}</p>
              <p style={{ color: "#4b5563" }}>Value: ${stock.value}</p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Portfolio;
