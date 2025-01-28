import React, { useState, useEffect } from "react";
import fetchingServices from "../services/fetchingServices";

const fetchPopularStocks = async () => {
    try {
        const symbols = ['AAPL', 'GOOGL', 'TSLA', 'MSFT', 'AMZN', 'NVDA', 'GOOG', 'META'];
        const names = ['Apple Inc', 'Alphabet Inc', 'Tesla Inc', 'Microsoft Corp', 'Amazon.com Inc', 'NVIDIA Corp', 'Alphabet Inc', 'Meta Platforms Inc']
        const quotes = await Promise.all(
          symbols.map((symbol) => fetchingServices.fetchStockPrice(symbol))
        );
    
        return symbols.map((symbol, index) => ({
          code: symbol,
          name: names[index],
          price: quotes[index],
        }));
      } catch (error) {
        console.error('Failed to fetch popular stocks:', error);
        throw error;
      }
};

const StocksPage = () => {
  const [popularStocks, setPopularStocks] = useState([]);
  const [filteredStocks, setFilteredStocks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [searchTerm, setSearchTerm] = useState("");
  const [sortOption, setSortOption] = useState("name");

  useEffect(() => {
    const loadPopularStocks = async () => {
      try {
        const stocks = await fetchPopularStocks();
        setPopularStocks(stocks);
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

    const sorted = filtered.sort((a, b) => {
      if (sortOption === "price") return b.price - a.price; 
      return a.name.localeCompare(b.name); 
    });

    setFilteredStocks(sorted);
  }, [searchTerm, sortOption, popularStocks]); 

  if (loading) {
    return <div>Loading popular stocks...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">Stocks</h1>

      <input
        type="text"
        placeholder="Search stocks..."
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
        className="border p-2 rounded w-full mb-4"
      />

      <div className="flex items-center mb-4">
        <label className="mr-2">Sort by:</label>
        <select
          value={sortOption}
          onChange={(e) => setSortOption(e.target.value)}
          className="border p-2 rounded"
        >
          <option value="name">Name</option>
          <option value="price">Price</option>
        </select>
      </div>

      <ul className="divide-y divide-gray-200">
        {filteredStocks.map((stock) => (
          <li key={stock.code} className="p-4 flex justify-between items-center">
            <div>
              <strong>{stock.name}</strong> ({stock.code})
            </div>
            <div className="text-right">
              <span className="block font-bold">${stock.price.toFixed(2)}</span>
            </div>
          </li>
        ))}
      </ul>

      {filteredStocks.length === 0 && (
        <div className="text-center text-gray-500">No stocks found.</div>
      )}
    </div>
  );
};

export default StocksPage;
