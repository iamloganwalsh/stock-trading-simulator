import React, { useEffect, useState } from "react";
import userServices from "../services/userServices";

const TradeHistory = () => {
  const [trades, setTrades] = useState([]);
  const [error, setError] = useState(null);  
  const [totalValue, setTotalValue] = useState(0);
  const [profitLoss, setProfitLoss] = useState(0);

  useEffect(() => {
    const fetchTradeHistory = async () => {
      try {
        const data = await userServices.getTradeHistory();
        setTrades(data);
        
        const total = data.reduce((sum, trade) => sum + trade.cost, 0);
        setTotalValue(total);
        
        const pl = data.reduce((sum, trade) => sum + (trade.profit || 0), 0);
        setProfitLoss(pl);
      } catch (err) {
        console.error("Failed to fetch trade history:", err);
        if (err && err.message && err.message == 'Network Error') {
          setError("Unable to load trade history. Please try again later.");
        } else {
          setError("Trade history is empty.")
        }
      }
    };

    fetchTradeHistory();
  }, []);

  if (error) {
    return <div style={{marginLeft: '100px'}}>{error}</div>;
  }

  return (
    <div style={{ padding: "24px", marginLeft: "100px" }}> 
      <div>
        <h2 style={{ fontSize: "24px", fontWeight: "bold", marginBottom: "16px" }}>Your Trades History</h2>
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
          {trades.map((trade, index) => (
            <div
              key={index}
              style={{
                backgroundColor: "#ffffff",
                boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
                padding: "16px",
                textAlign: "center",
                transition: "transform 0.2s",
                borderRadius: "8px",
                position: "relative",
                zIndex: 1,
              }}
              onMouseEnter={(e) => (e.currentTarget.style.transform = "scale(1.05)")}
              onMouseLeave={(e) => (e.currentTarget.style.transform = "scale(1)")}
            >
              <h3 style={{ fontSize: "18px", fontWeight: "bold", marginBottom: "8px", color:'#242424' }}>{trade.code}</h3>
              <div style={{ color: "#4b5563", marginBottom: "8px" }}>
                <p>Type: {trade.type}</p>
                <p>Method: {trade.method}</p>
                <p style={{ color: "#16a34a", fontWeight: "bold" }}>
                  Cost: ${trade.cost.toFixed(2)}
                </p>
                <p>Date: {trade.date}</p>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default TradeHistory;
