import { useState, useEffect } from 'react'
import './App.css'
import userServices from './services/userServices.js'



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
        setError('Failed to fetch user data.');
      } finally {
        setLoading(false);
      }
    };

    fetchUserData();
  }, []);

  return (
    <>
      {loading ? (
        <p>Loading user data...</p>
      ) : error ? (
        <p style={{ color: 'red'}}>{error}</p>
      ) : (
        <>
          <p>Balance: {balance}</p>
          <p>Username: {username}</p>
          <p>Profit / Loss: {profitloss}</p>
        </>
      )}
    </>
  );
}

export default App
