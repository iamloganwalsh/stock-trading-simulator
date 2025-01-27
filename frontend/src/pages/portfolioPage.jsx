import React, { useState, useEffect } from "react";
import userServices from "../services/userServices";
import fetchingServices from "../services/fetchingServices";

const Portfolio = ({ balance, profitloss, loading, error }) => {
  const [portfolio, setPortfolio] = useState({
    totalBalance: balance || 0,
    profitLoss: profitloss || 0,
    stocks: [],
  });

  useEffect(() => {
    const fetchStocks = async () => {
      try {
        const stocksData = await userServices.getStockPortfolio();
        const updatedStocks = await Promise.all(
          stocksData.map(async (stock) => {
            const newPrice = await fetchingServices.fetchStockPrice(stock.code);
            return { ...stock, value: newPrice};
          })
        );

        setPortfolio((prevPortfolio) => ({
          ...prevPortfolio,
          stocks: updatedStocks,
        }));
      } catch (err) {
        console.error("Error fetching stocks:", err);
      }
    };
    fetchStocks();
  }, []);

  const fetchUpdatedStockPrices = async () => {
    try {
      const updatedStocks = await Promise.all(
        portfolio.stocks.map(async (stock) => {
          const newPrice = await fetchingServices.fetchStockPrice(stock.code);
          return { ...stock, value: newPrice };
        })
      );
      setPortfolio((prevPortfolio) => ({
        ...prevPortfolio,
        stocks: updatedStocks,
      }));
    } catch (err) {
      console.error("Error fetching updated stock prices:", err);
    }
  };

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
      <button
        onClick={fetchUpdatedStockPrices}
        style={{
          marginTop: "24px",
          padding: "10px 16px",
          backgroundColor: "#16a34a",
          color: "white",
          border: "none",
          borderRadius: "8px",
          cursor: "pointer",
        }}
      >
        Refresh Stock Prices
      </button>
    </div>
  );
};

export default Portfolio;
