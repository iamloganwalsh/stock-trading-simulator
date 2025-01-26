import React from "react";

function Account({ balance, username, profitloss, loading, error }) {
  return (
    <div style={{ padding: '20px', marginLeft: '80px' }}>
      {loading ? (
        <p>Loading user data...</p>
      ) : error ? (
        <p style={{ color: 'red' }}>{error}</p>
      ) : (
        <>
          <h1>Profile</h1>
          <p>Username: {username}</p>
          <p>Balance: {balance}</p>
          <p>Profit / Loss: {profitloss}</p>
        </>
      )}
    </div>
  );
}

export default Account;