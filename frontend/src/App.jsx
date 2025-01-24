import { useState, useEffect } from 'react'
import './App.css'
import userServices from './services/userServices.js'



function App() {

  const [balance, setBalance] = useState(null)

  useEffect(() => {
    const fetchBalance = async () => {
      const balanceData = await userServices.getBalance();
      setBalance(balanceData);
    }

    fetchBalance();
  }, []);

  return (
    <>
    <p>{balance ? `Balance: ${balance}` : 'Loading balance...'}</p>
    </>
  )
}

export default App
