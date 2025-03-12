import React, { useState, useEffect } from "react";
import userServices from "../services/userServices";
import fetchingServices from "../services/fetchingServices";
import { useNavigate } from 'react-router-dom';

const Portfolio = ({ balance, profitloss, loading, error }) => {
  const [portfolio, setPortfolio] = useState({
    totalBalance: balance || 0,
    profitLoss: profitloss || 0,
    stocks: [],
    crypto: [],
  });

  const navigate = useNavigate();

  const convertToFinnhubCode = (crypto_code) => {
    const cryptoCode = (crypto_code || "BTC").toUpperCase();
  
    if (cryptoCode.includes("-USD")) {
      return `BINANCE:${cryptoCode.replace("-USD", "")}USDT`;
    } 
    if (cryptoCode.startsWith("BINANCE:") && cryptoCode.endsWith("USDT")) {
      return cryptoCode;
    }
    
    return `BINANCE:${cryptoCode}USDT`;
  };
  

  useEffect(() => {
    const fetchAssets = async () => {
      try {
        const stocksData = await userServices.getStockPortfolio();
        const cryptoData = await userServices.getCryptoPortfolio(); 

        const validStocks = Array.isArray(stocksData) ? stocksData : [];
        const updatedStocks = await Promise.all(
          validStocks.map(async (stock) => {
            const newPrice = await fetchingServices.fetchStockPrice(stock.code);
            return { ...stock, value: newPrice };
          })
        );

        const validCrypto = Array.isArray(cryptoData) ? cryptoData : [];
        const updatedCrypto = await Promise.all(
          validCrypto.map(async (crypto) => {
            const finnhub_code = convertToFinnhubCode(crypto.code);
            const fetchedPrice = await fetchingServices.fetchCryptoPrice(finnhub_code);
            const newPrice = fetchedPrice * crypto.crypto_count;
            return { ...crypto, value: newPrice };
          })
        );

        setPortfolio((prevPortfolio) => ({
          ...prevPortfolio,
          stocks: updatedStocks,
          crypto: updatedCrypto, 
        }));
      } catch (err) {
        console.error("Error fetching assets:", err);
        setPortfolio((prev) => ({
          ...prev,
          stocks: [],
          crypto: [], 
        }));
      }
    };
    

    fetchAssets();
  }, []);

  const fetchUpdatedPrices = async () => {
    try {
      const updatedStocks = await Promise.all(
        portfolio.stocks.map(async (stock) => {
          const newPrice = await fetchingServices.fetchStockPrice(stock.code);
          return { ...stock, value: newPrice };
        })
      );

      const updatedCrypto = await Promise.all(
        portfolio.crypto.map(async (crypto) => {
          const finnhub_code = convertToFinnhubCode(crypto.code);
          const fetchedPrice = await fetchingServices.fetchCryptoPrice(finnhub_code);
          const newPrice = fetchedPrice * crypto.crypto_count;
          return { ...crypto, value: newPrice };
        })
        
      );

      setPortfolio((prevPortfolio) => ({
        ...prevPortfolio,
        stocks: updatedStocks,
        crypto: updatedCrypto,
      }));
    } catch (err) {
      console.error("Error fetching updated prices:", err);
    }
  };

  const handleRedirectStock = (stockCode) => {
    navigate(`/view/stock/${stockCode}`);
  };

  const handleRedirectCrypto = (cryptoCode) => {
    navigate(`/view/crypto/${cryptoCode}`);
  };

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div style={{ padding: "24px", marginLeft: "100px" }}>
      <div
        style={{
          display: "grid",
          gridTemplateColumns: "1fr 1fr",
          gap: "24px",
          maxWidth: "1200px",
          marginInline: "auto",
        }}
      >
        <div
          style={{
            backgroundColor: "#f5f5f5",
            boxShadow: "0 2px 4px rgba(70, 167, 25, 0.1)",
            padding: "16px",
            textAlign: "center",
            borderRadius: "10px",
          }}
        >
          <h2 style={{ fontSize: "20px", fontWeight: "bold", color: "#242424" }}>Total Balance</h2>
          <p style={{ fontSize: "24px", color: "#16a34a" }}>${portfolio.totalBalance}</p>
        </div>
        <div
          style={{
            backgroundColor: "#f5f5f5",
            boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
            padding: "16px",
            textAlign: "center",
            borderRadius: "10px",
          }}
        >
          <h2 style={{ fontSize: "20px", fontWeight: "bold", color: "#242424" }}>Profit/Loss</h2>
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
                color: "#242424",
                borderRadius: "10px",
                cursor: "pointer"
              }}
              onMouseEnter={(e) => (e.currentTarget.style.transform = "scale(1.05)")}
              onMouseLeave={(e) => (e.currentTarget.style.transform = "scale(1)")}
            >
              <h3 style={{ fontSize: "18px", fontWeight: "bold" }}>{stock.code}</h3>
              <p style={{ color: "#4b5563" }}>Shares: {stock.stock_count}</p>
              <p style={{ color: "#4b5563" }}>Value: ${stock.value}</p>
              <button onClick={() => handleRedirectStock(stock.code)}  style={{
                  marginTop: "8px",
                  padding: "6px 12px",
                  backgroundColor: "#3b82f6",
                  color: "white",
                  border: "none",
                  borderRadius: "4px",
                  cursor: "pointer",
                  marginRight: "8px",
                }}>Buy</button>
              <button onClick={() => handleRedirectStock(stock.code)} style={{
                  marginTop: "8px",
                  padding: "6px 12px",
                  backgroundColor: "#ef4444",
                  color: "white",
                  border: "none",
                  borderRadius: "4px",
                  cursor: "pointer",
                }}>Sell</button>
            </div>
          ))}
        </div>
      </div>

      <div>
        <h2 style={{ fontSize: "24px", fontWeight: "bold", marginBottom: "16px" }}>Your Cryptos</h2>
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
          {portfolio.crypto.map((crypto, index) => (
            <div
              key={index}
              style={{
                backgroundColor: "#ffffff",
                boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
                padding: "16px",
                textAlign: "center",
                transition: "transform 0.2s",
                color: "#242424",
                borderRadius: "10px",
                cursor: "pointer"
              }}
              onMouseEnter={(e) => (e.currentTarget.style.transform = "scale(1.05)")}
              onMouseLeave={(e) => (e.currentTarget.style.transform = "scale(1)")}
            >
              <h3 style={{ fontSize: "18px", fontWeight: "bold" }}>{crypto.code}</h3>
              <p style={{ color: "#4b5563" }}>Coins: {crypto.crypto_count}</p>
              <p style={{ color: "#4b5563" }}>Value: ${crypto.value}</p>
              <button onClick={() => handleRedirectCrypto(crypto.code)}  style={{
                  marginTop: "8px",
                  padding: "6px 12px",
                  backgroundColor: "#3b82f6",
                  color: "white",
                  border: "none",
                  borderRadius: "4px",
                  cursor: "pointer",
                  marginRight: "8px",
                }}>Buy</button>
              <button onClick={() => handleRedirectCrypto(crypto.code)} style={{
                  marginTop: "8px",
                  padding: "6px 12px",
                  backgroundColor: "#ef4444",
                  color: "white",
                  border: "none",
                  borderRadius: "4px",
                  cursor: "pointer",
                }}>Sell</button>
            </div>
          ))}
        </div>
      </div>

      <button onClick={fetchUpdatedPrices} style={{
          marginTop: "24px",
          padding: "10px 16px",
          backgroundColor: "#16a34a",
          color: "white",
          border: "none",
          borderRadius: "8px",
          cursor: "pointer",
        }}>Refresh Prices</button>
    </div>
  );
};

export default Portfolio;
