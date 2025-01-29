import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import userServices from '../services/userServices';
import stockServices from '../services/stockServices';
import fetchingServices from '../services/fetchingServices';

const Buysell = () => {
  const { stockCode } = useParams(); // Get stock code from URL parameter
  const [selectedStock, setSelectedStock] = useState(stockCode || '');
  const [quantity, setQuantity] = useState('');
  const [portfolio, setPortfolio] = useState([]);
  const [balance, setBalance] = useState(0);
  const [username, setUsername] = useState('');
  const [stockPrice, setStockPrice] = useState(null);
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    fetchUserData();
  }, []);

  useEffect(() => {
    if (selectedStock) {
      fetchStockPrice(selectedStock);
    }
  }, [selectedStock]);

  const fetchUserData = async () => {
    try {
      const [balanceData, usernameData, portfolioData] = await Promise.all([
        userServices.getBalance(),
        userServices.getUsername(),
        userServices.getStockPortfolio()
      ]);

      setBalance(balanceData);
      setUsername(usernameData);
      setPortfolio(portfolioData);
    } catch (error) {
      setError('Failed to fetch user data');
    }
  };

  const fetchStockPrice = async (code) => {
    if (!code) return;
    try {
      const data = await fetchingServices.fetchStockPrice(code);
      setStockPrice(data);
      setError('');
    } catch (error) {
      setError('Failed to fetch stock price');
      setStockPrice(null);
    }
  };

  const handleTransaction = async (type) => {
    if (!selectedStock || !quantity || !stockPrice) {
      setError('Please fill in all fields');
      return;
    }

    setLoading(true);
    setError('');

    try {
      if (type === 'buy') {
        await stockServices.buyStock(selectedStock, stockPrice, parseFloat(quantity));
      } else {
        await stockServices.sellStock(selectedStock, stockPrice, parseFloat(quantity));
      }
      
      await fetchUserData();
      setQuantity('');
      setSelectedStock('');
      setStockPrice(null);
    } catch (error) {
      setError(error.response?.data?.error || `Failed to ${type} stock`);
    } finally {
      setLoading(false);
    }
  };

  const containerStyle = {
    maxWidth: '800px',
    margin: '0 auto',
    padding: '20px',
    fontFamily: 'Arial, sans-serif'
  };

  const cardStyle = {
    border: '1px solid #ddd',
    borderRadius: '8px',
    padding: '20px',
    marginBottom: '20px',
    backgroundColor: 'white'
  };

  const inputStyle = {
    width: '100%',
    padding: '8px',
    marginBottom: '10px',
    border: '1px solid #ddd',
    borderRadius: '4px'
  };

  const buttonStyle = {
    padding: '10px 20px',
    border: 'none',
    borderRadius: '4px',
    cursor: 'pointer',
    marginRight: '10px',
    color: 'white'
  };

  const buyButtonStyle = {
    ...buttonStyle,
    backgroundColor: '#4CAF50'
  };

  const sellButtonStyle = {
    ...buttonStyle,
    backgroundColor: '#f44336'
  };

  const errorStyle = {
    color: 'red',
    marginBottom: '10px'
  };

  return (
    <div style={containerStyle}>
      {error && <div style={errorStyle}>{error}</div>}

      <div style={cardStyle}>
        <h2>Account Overview</h2>
        <p>Welcome, {username}</p>
        <p style={{ fontSize: '24px', fontWeight: 'bold' }}>${balance ? balance.toFixed(2) : '0.00'}</p>
      </div>

      <div style={cardStyle}>
        <h2>Trade Stocks</h2>
        <div>
          <label>Stock Code</label>
          <input
            type="text"
            value={selectedStock}
            onChange={(e) => setSelectedStock(e.target.value.toUpperCase())}
            placeholder="Enter stock code (e.g., AAPL)"
            style={inputStyle}
          />
        </div>

        {stockPrice && (
          <p>Current Price: ${stockPrice.toFixed(2)}</p>
        )}

        <div>
          <label>Quantity</label>
          <input
            type="number"
            value={quantity}
            onChange={(e) => setQuantity(e.target.value)}
            placeholder="Enter quantity"
            min="0"
            step="0.01"
            style={inputStyle}
          />
        </div>

        {stockPrice && quantity && (
          <p>Total Cost: ${(stockPrice * parseFloat(quantity || 0)).toFixed(2)}</p>
        )}

        <div>
          <button 
            onClick={() => handleTransaction('buy')}
            disabled={loading}
            style={buyButtonStyle}
          >
            {loading ? 'Processing...' : 'Buy'}
          </button>
          <button 
            onClick={() => handleTransaction('sell')}
            disabled={loading}
            style={sellButtonStyle}
          >
            {loading ? 'Processing...' : 'Sell'}
          </button>
        </div>
      </div>

      <div style={cardStyle}>
        <h2>Your Portfolio</h2>
        {portfolio.length === 0 ? (
          <p>No stocks in portfolio</p>
        ) : (
          portfolio.map((item) => (
            <div key={item.code} style={{ borderBottom: '1px solid #eee', padding: '10px 0' }}>
              <h3>{item.code}</h3>
              <p>{item.stock_count} shares</p>
            </div>
          ))
        )}
      </div>
    </div>
  );
};

export default Buysell;