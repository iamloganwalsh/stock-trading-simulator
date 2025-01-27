import { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import userServices from './services/userServices.js';
import NavBar from './components/navbar.jsx';
import Home from './pages/homePage.jsx';
import AccountPage from './pages/accountPage.jsx';
import PortfolioPage from './pages/portfolioPage.jsx'

import fetchingServices from './services/fetchingServices.js'

function App() {

  const [balance, setBalance] = useState(null)
  const [username, setUsername] = useState(null)
  const [profitloss, setProfitLoss] = useState(null)
  const [error, setError] = useState(null)
  const [loading, setLoading] = useState(null)

  useEffect(() => {
    const fetchUserData = async () => {
      try{
        const [balanceData, usernameData, profitlossData] = await Promise.all([
          userServices.getBalance(),
          userServices.getUsername(),
          userServices.getProfitLoss(),
        ]);

        setBalance(balanceData);
        setUsername(usernameData);
        setProfitLoss(profitlossData)
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
              profitloss={profitloss}
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
              profitloss={profitloss}
              loading={loading}
              error={error}
            />
          }
        />
        <Route
          path="/"
          element={<Home />}
        />
      </Routes>
    </>
  );
};

export default App