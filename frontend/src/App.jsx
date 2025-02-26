import { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import userServices from './services/userServices.js';
import NavBar from './components/navbar.jsx';
import Home from './pages/homePage.jsx';
import AccountPage from './pages/accountPage.jsx';
import PortfolioPage from './pages/portfolioPage.jsx'
import CryptoView from './pages/cryptoView.jsx'
import StockView from './pages/stockView.jsx'


import Trade from './pages/tradePage.jsx';
import StocksPage from './pages/marketPage.jsx';
import Buysell from './pages/buysellPage.jsx';

function App() {

  const [balance, setBalance] = useState(null)
  const [username, setUsername] = useState(null)
  const [investment, setinvestment] = useState(null)
  const [error, setError] = useState(null)
  const [loading, setLoading] = useState(null)

  useEffect(() => {
    const fetchUserData = async () => {
      try{
        const [balanceData, usernameData, investmentData] = await Promise.all([
          userServices.getBalance(),
          userServices.getUsername(),
          userServices.getInvestment(),
        ]);

        setBalance(balanceData);
        setUsername(usernameData);
        setinvestment(investmentData)
      } catch (err) {
        setError('Failed to fetch user data:', err);
      } finally {
        setLoading(false);
      }
    };

    fetchUserData();
  }, []);

  return (
    <>
      <NavBar />
      <Routes>
        <Route
          path="/portfolio"
          element={
            <PortfolioPage
              balance={balance}
              investment={investment}
              loading={loading}
              error={error}
            />
          }
        />
        <Route
          path="/account"
          element={
            <AccountPage
              balance={balance}
              username={username}
              investment={investment}
              loading={loading}
              error={error}
            />
          }
        />
        <Route
          path="/trade"
          element={<Trade />}
        />
        <Route
          path="/market"
          element={<StocksPage />}
        />
        <Route
          path="/"
          element={<Home />}
        />
        <Route
        path="/view/crypto/:crypto_code?"
        element={<CryptoView />}
        />
        <Route 
        path="/view/stock/:stock_code?"
        element={<StockView />}
        />
      </Routes>
    </>
  );
};

export default App