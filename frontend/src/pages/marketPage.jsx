import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import fetchingServices from "../services/fetchingServices";

const StocksPage = () => {
  const navigate = useNavigate();
  const [cryptos, setCryptos] = useState([]);
  const [stocks, setStocks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [searchTerm, setSearchTerm] = useState("");
  const [sortOption, setSortOption] = useState("name");

  useEffect(() => {
    const loadData = async () => {
      try {
        const cryptoSymbols = ["BTC", "ETH", "XRP", "LTC", "SOL", "BNB", "ADA", "DOGE"];
        const cryptoNames = ["Bitcoin", "Ethereum", "XRP", "Litecoin", "Solana", "BNB", "Cardano", "Dogecoin"];
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
        
        const finnhubSymbols = cryptoSymbols.map(convertToFinnhubCode);
        
        const cryptoPrices = await Promise.all(
          finnhubSymbols.map((finnhub_code) => fetchingServices.fetchCryptoPrice(finnhub_code))
        );

        const stockSymbols = ["AAPL", "GOOG", "TEAM", "AMZN", "MSFT", "TSLA", "NVDA", "META"];
        const stockNames = ["Apple Inc", "Alphabet Inc", "Atlassian Corp", "Amazon.com Inc", "Microsoft Corp", "Tesla Inc", "NVIDIA Corp", "Meta Platforms Inc"];
        const stockPrices = await Promise.all(
          stockSymbols.map((symbol) => fetchingServices.fetchStockPrice(symbol))
        );

        setCryptos(cryptoSymbols.map((symbol, index) => ({
          code: symbol,
          name: cryptoNames[index],
          price: cryptoPrices[index]
        })));

        setStocks(stockSymbols.map((symbol, index) => ({
          code: symbol,
          name: stockNames[index],
          price: stockPrices[index]
        })));
      } catch (err) {
        setError("Unable to load market data." + err);
      } finally {
        setLoading(false);
      }
    };

    loadData();
  }, []);

  const handleSort = (option) => {
    setSortOption(option);
    setCryptos([...cryptos].sort((a, b) => option === "name" ? a.name.localeCompare(b.name) : a.price - b.price));
    setStocks([...stocks].sort((a, b) => option === "name" ? a.name.localeCompare(b.name) : a.price - b.price));
  };

  const filteredCryptos = cryptos.filter((crypto) =>
    crypto.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    crypto.code.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const filteredStocks = stocks.filter((stock) =>
    stock.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    stock.code.toLowerCase().includes(searchTerm.toLowerCase())
  );

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div style={{ padding: "24px", margin: "0 0 0 100px" }}>
      <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center", marginBottom: "16px", marginInline: "auto", maxWidth: "906px" }}>
        <input
          type="text"
          placeholder="Search stocks or cryptocurrencies..."
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          style={{ padding: "10px", border: "1px solid #ccc", borderRadius: "8px", flex: 1, fontSize: "16px" }}
        />
        <button onClick={() => handleSort("name")} style={{ marginLeft: "10px", color: 'white' }}>Sort by Name</button>
        <button onClick={() => handleSort("price")} style={{ marginLeft: "10px", color: 'white' }}>Sort by Price</button>
      </div>
      <div style={{ display: "flex", flexDirection: "row", gap: "70px", justifyContent: "center" }}>
        <div>
          <h2>Cryptocurrencies</h2>
            <div style={{ display: "grid", gridTemplateColumns: "repeat(2, minmax(200px, 1fr))", gap: "20px", height: "fit-content" }}>
            {filteredCryptos.map((crypto) => (
              <div
                key={crypto.code}
                onClick={() => navigate(`/view/crypto/${crypto.code}`)}
                style={{
                  backgroundColor: "#fff",
                  padding: "20px",
                  borderRadius: "12px",
                  boxShadow: "0 4px 8px rgba(0, 0, 0, 0.1)",
                  textAlign: "center",
                  color: "#242424",
                  marginBottom: "10px",
                  height: "fit-content",
                  cursor: "pointer"
                }}
              >
                <h3>{crypto.code}</h3>
                <p>{crypto.name}</p>
                <p style={{ fontSize: "18px", fontWeight: "bold", color: "#007BFF" }}>
                  ${crypto.price.toFixed(2)}
                </p>
              </div>
            ))}
          </div>
        </div>
        
        <div>
          <h2>Stocks</h2>
            <div style={{ display: "grid", gridTemplateColumns: "repeat(2, minmax(200px, 1fr))", gap: "20px", height: "fit-content" }}>
            {filteredStocks.map((stock) => (
              <div
                key={stock.code}
                onClick={() => navigate(`/view/stock/${stock.code}`)}
                style={{
                  backgroundColor: "#fff",
                  padding: "20px",
                  borderRadius: "12px",
                  boxShadow: "0 4px 8px rgba(0, 0, 0, 0.1)",
                  textAlign: "center",
                  color: "#242424",
                  marginBottom: "10px",
                  height: "fit-content",
                  cursor: "pointer"
                }}
              >
                <h3>{stock.code}</h3>
                <p>{stock.name}</p>
                <p style={{ fontSize: "18px", fontWeight: "bold", color: "#007BFF" }}>
                  ${stock.price.toFixed(2)}
                </p>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};


export default StocksPage;
