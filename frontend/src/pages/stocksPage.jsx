import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Search, ArrowUpDown } from "lucide-react";
import fetchingServices from "../services/fetchingServices";

const StocksPage = () => {
  const navigate = useNavigate();
  const [popularStocks, setPopularStocks] = useState([]);
  const [filteredStocks, setFilteredStocks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [searchTerm, setSearchTerm] = useState("");
  const [sortOption, setSortOption] = useState("name");

  useEffect(() => {
    const loadPopularStocks = async () => {
      try {
        const symbols = ["AAPL", "GOOGL", "TSLA", "MSFT", "AMZN", "NVDA", "GOOG", "META"];
        const names = ["Apple Inc", "Alphabet Inc", "Tesla Inc", "Microsoft Corp", "Amazon.com Inc", "NVIDIA Corp", "Alphabet Inc", "Meta Platforms Inc"];
        const quotes = await Promise.all(
          symbols.map((symbol) => fetchingServices.fetchStockPrice(symbol))
        );
        
        const stocks = symbols.map((symbol, index) => ({
          code: symbol,
          name: names[index],
          price: quotes[index]
        }));
        
        setPopularStocks(stocks);
        setFilteredStocks(stocks);
      } catch (err) {
        setError("Unable to load popular stocks.");
      } finally {
        setLoading(false);
      }
    };

    loadPopularStocks();
  }, []);

  useEffect(() => {
    const filtered = popularStocks.filter(
      (stock) =>
        (stock.name && stock.name.toLowerCase().includes(searchTerm.toLowerCase())) ||
        (stock.code && stock.code.toLowerCase().includes(searchTerm.toLowerCase()))
    );

    const sorted = [...filtered].sort((a, b) => {
      if (sortOption === "price") {
        return (b.price ?? 0) - (a.price ?? 0);
      }
      return (a.name ?? "").localeCompare(b.name ?? "");
    });

    setFilteredStocks(sorted);
  }, [searchTerm, sortOption, popularStocks]);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div style={{ padding: "24px", maxWidth: "1200px", margin: "0 auto" }}>
      <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center", marginBottom: "16px" }}>
        <input
          type="text"
          placeholder="Search stocks..."
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          style={{ padding: "10px", border: "1px solid #ccc", borderRadius: "8px", flex: 1, fontSize: "16px" }}
        />
        <button 
          onClick={() => setSortOption(sortOption === "name" ? "price" : "name")}
          style={{ marginLeft: "10px", padding: "10px 16px", backgroundColor: "#007BFF", color: "#fff", border: "none", borderRadius: "8px", cursor: "pointer", fontSize: "16px" }}
        >
          <ArrowUpDown size={18} style={{ marginRight: "5px" }} /> Sort by {sortOption === "name" ? "Price" : "Name"}
        </button>
      </div>

      <div style={{
        display: "grid",
        gridTemplateColumns: "repeat(auto-fit, minmax(250px, 1fr))",
        gap: "20px",
        alignItems: "stretch"
      }}>
        {filteredStocks.map((stock) => (
          <div
            key={stock.code}
            onClick={() => navigate(`/stocksportfolio/${stock.code}`)}
            style={{
              backgroundColor: "#fff",
              padding: "20px",
              borderRadius: "12px",
              boxShadow: "0 4px 8px rgba(0, 0, 0, 0.1)",
              cursor: "pointer",
              textAlign: "center",
              transition: "transform 0.2s, box-shadow 0.2s",
              display: "flex",
              flexDirection: "column",
              alignItems: "center",
              justifyContent: "center"
            }}
            onMouseEnter={(e) => e.currentTarget.style.transform = "scale(1.05)"}
            onMouseLeave={(e) => e.currentTarget.style.transform = "scale(1)"}
          >
            <h3 style={{ fontSize: "20px", fontWeight: "bold", marginBottom: "5px" }}>{stock.code}</h3>
            <p style={{ fontSize: "14px", color: "#555", marginBottom: "10px" }}>{stock.name}</p>
            <p style={{ fontSize: "18px", fontWeight: "bold", color: "#007BFF" }}>
              ${stock.price.toFixed(2)}
            </p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default StocksPage;