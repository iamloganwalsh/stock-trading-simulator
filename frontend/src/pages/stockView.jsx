import React from 'react';
import StockGraph from '../components/stockGraph';
import { useParams } from "react-router-dom";
import TradeInput from '../components/tradeInput';

const StockView = () => {
  const styles = {
    container: {
      marginLeft: '120px', // Offset for the fixed NavBar
      padding: '20px',
      width: '100%',
    },
    heading: {
      fontSize: '24px',
      marginBottom: '10px',
    },
    paragraph: {
      fontSize: '16px',
    },
  };

  const { stock_code } = useParams();
  const stockCode = (stock_code || "AAPL").toUpperCase();

  return (
    <div style={styles.container}>

      <div>
        <StockGraph stock_name={stockCode} yahoo_code={stockCode} finnhub_code={stockCode} />
        <TradeInput type="stock" code={stockCode} finnhub_code={stockCode}/>
      </div>
      
    </div>
    
  );
};

export default StockView;