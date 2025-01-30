import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import userServices from '../services/userServices';
import stockServices from '../services/stockServices';
import cryptoServices from '../services/cryptoServices'; // Add the crypto services
import fetchingServices from '../services/fetchingServices';

const Buysell = () => {
  const { stockCode, cryptoCode } = useParams(); 
  const [selectedStock, setSelectedStock] = useState(stockCode || '');
  const [selectedCrypto, setSelectedCrypto] = useState(cryptoCode || '');
  const [quantity, setQuantity] = useState('');
  const [portfolio, setPortfolio] = useState([]);
  const [cryptoPortfolio, setCryptoPortfolio] = useState([]); 
  const [balance, setBalance] = useState(0);
  const [username, setUsername] = useState('');
  const [stockPrice, setStockPrice] = useState(null);
  const [cryptoPrice, setCryptoPrice] = useState(null);
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

  useEffect(() => {
    if (selectedCrypto) {
      fetchCryptoPrice(selectedCrypto);
    }
  }, [selectedCrypto]);

  const fetchUserData = async () => {
    try {
      const [balanceData, usernameData, portfolioData, cryptoPortfolioData] = await Promise.all([
        userServices.getBalance(),
        userServices.getUsername(),
        userServices.getStockPortfolio(),
        userServices.getCryptoPortfolio()  
      ]);

      setBalance(balanceData);
      setUsername(usernameData);
      setPortfolio(portfolioData);
      setCryptoPortfolio(cryptoPortfolioData);  
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

  const fetchCryptoPrice = async (code) => {
    if (!code) return;
    try {
      const data = await fetchingServices.fetchCryptoPrice(code);
      setCryptoPrice(data);
      setError('');
    } catch (error) {
      setError('Failed to fetch crypto price');
      setCryptoPrice(null);
    }
  };

  const handleTransaction = async (type) => {
    if ((!selectedStock && !selectedCrypto) || !quantity || isNaN(parseFloat(quantity)) || parseFloat(quantity) <= 0 || (stockPrice === null && cryptoPrice === null)) {
      setError('Please fill in all fields correctly');
      return;
    }

    setLoading(true);
    setError('');

    try {
      if (selectedStock) {
        if (type === 'buy') {
          await stockServices.buyStock(selectedStock, stockPrice, parseFloat(quantity));
        } else {
          await stockServices.sellStock(selectedStock, stockPrice, parseFloat(quantity));
        }
      }

      if (selectedCrypto) {
        if (type === 'buy') {
          await cryptoServices.buyCrypto(selectedCrypto, cryptoPrice, parseFloat(quantity));
        } else {
          await cryptoServices.sellCrypto(selectedCrypto, cryptoPrice, parseFloat(quantity));
        }
      }

      await fetchUserData();
      setQuantity('');
      setStockPrice(null);
      setCryptoPrice(null); 
    } catch (error) {
      setError(error.response?.data?.error || `Failed to ${type} ${selectedStock ? 'stock' : 'crypto'}`);
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
    backgroundColor: 'white',
    color: 'black'
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
        <h2>Trade Stocks / Crypto</h2>
        <div>
          <label>Code</label>
          <input
            type="text"
            value={selectedStock || selectedCrypto}
            onChange={(e) => {
              const value = e.target.value.toUpperCase();
              setSelectedStock(value);
              setSelectedCrypto(value);
            }}
            placeholder="Enter stock or crypto code (e.g., AAPL, BTC)"
            style={inputStyle}
          />
        </div>

        {(stockPrice !== null || cryptoPrice !== null) && (
          <p>
            {selectedStock ? `Stock Price: $${stockPrice?.toFixed(2)}` : `Crypto Price: $${cryptoPrice?.toFixed(2)}`}
          </p>
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

        {(stockPrice !== null || cryptoPrice !== null) && quantity && !isNaN(parseFloat(quantity)) && (
          <p>Total Cost: ${(selectedStock ? stockPrice : cryptoPrice) * parseFloat(quantity).toFixed(2)}</p>
        )}

        <button onClick={() => handleTransaction('buy')} disabled={loading} style={buyButtonStyle}>
          {loading ? 'Processing...' : 'Buy'}
        </button>
        <button onClick={() => handleTransaction('sell')} disabled={loading} style={sellButtonStyle}>
          {loading ? 'Processing...' : 'Sell'}
        </button>
      </div>

      <div style={cardStyle}>
        <h2>Your Stock Portfolio</h2>
        {portfolio.length === 0 ? (
          <p>No stocks in portfolio</p>
        ) : (
          portfolio.map((item) => (
            <div key={item.code} style={{ borderBottom: '1px solid #eee', padding: '10px 0', color: 'black' }}>
              <h3>{item.code}</h3>
              <p>{item.stock_count} shares</p>
            </div>
          ))
        )}
      </div>

      <div style={cardStyle}>
        <h2>Your Crypto Portfolio</h2>
        {cryptoPortfolio.length === 0 ? (
          <p>No cryptocurrencies in portfolio</p>
        ) : (
          cryptoPortfolio.map((item) => (
            <div key={item.code} style={{ borderBottom: '1px solid #eee', padding: '10px 0', color: 'black' }}>
              <h3>{item.code}</h3>
              <p>{item.crypto_count} units</p>
            </div>
          ))
        )}
      </div>
    </div>
  );
};

export default Buysell;
