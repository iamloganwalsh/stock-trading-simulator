import { useState, useEffect } from "react";
import fetchingServices from '../services/fetchingServices';
import cs from '../services/cryptoServices';
//import ss from '../services/stockServices';

const TradeInput = ({ type, code, finnhub_code }) => {
  const [amount, setAmount] = useState(""); // Allow empty string initially
  const [currPrice, setCurrPrice] = useState(0);

  // Fetch the latest price from Finnhub
  const fetchPrice = async () => {
    try {
      const data = await fetchingServices.fetchCryptoPrice(finnhub_code);
      console.log("New Price Data:", data);
      setCurrPrice(data);
    } catch (error) {
      console.error("Error fetching price:", error);
    }
  };

  useEffect(() => {
    fetchPrice(); // Initial fetch

    const interval = setInterval(fetchPrice, 10000); // Fetch price every 10 seconds

    return () => clearInterval(interval); // Cleanup on unmount
  }, [finnhub_code]);

  const handleTrade = async (action) => {
    const parsedAmount = parseFloat(amount); // Convert input to number
    if (isNaN(parsedAmount) || parsedAmount <= 0) {
      alert("Please enter a valid amount.");
      return;
    }

    try {
      const response =
        type === "crypto"
          ? await (action === "buy"
              ? cs.buyCrypto(code, currPrice, parsedAmount)
              : cs.sellCrypto(code, currPrice, parsedAmount))
          : await (action === "buy"
              ? ss.buyStock(code, currPrice, parsedAmount)
              : ss.sellStock(code, currPrice, parsedAmount));

      if (response?.status !== 201) {
        console.log(response);
        alert("Error processing the transaction");
      } else {
        alert("Transaction successful");
      }
    } catch (error) {
      console.error("Trade Error:", error);
      alert(`An error occurred while processing your trade:\n${error.response.data}`);
    }
  };

  return (
    <div className="p-4 border rounded-lg shadow-lg w-80 bg-white">

<input
  type="number"
  placeholder="Enter amount"
  value={amount}
  onChange={(e) => setAmount(e.target.value)}
  style={{
    width: "100% - 20px",
    padding: "10px",
    marginTop: "8px",
    border: "1px solid #ccc",
    borderRadius: "8px",
    fontSize: "16px",
    outline: "none",
    transition: "border-color 0.3s ease-in-out, box-shadow 0.3s ease-in-out",
  }}
  onFocus={(e) => {
    e.target.style.borderColor = "#007bff";
    e.target.style.boxShadow = "0 0 5px rgba(0, 123, 255, 0.5)";
  }}
  onBlur={(e) => {
    e.target.style.borderColor = "#ccc";
    e.target.style.boxShadow = "none";
  }}
/>



      <p className="mt-2 text-gray-700">
        Total: {isNaN(amount * currPrice) ? "NaN" : `$${(currPrice * amount).toFixed(2)}`}
      </p>

      <div style={{
        display: "flex",
        justifyContent: "center",
        gap: "15px",
      }}>
        <button
          className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600"
          onClick={() => handleTrade("buy")}
        >
          Buy
        </button>

        <button
          className="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600"
          onClick={() => handleTrade("sell")}
        >
          Sell
        </button>
      </div>
    </div>
  );
};

export default TradeInput;
