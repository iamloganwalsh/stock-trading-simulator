import React from 'react';
import CryptoGraph from '../components/cryptoGraph';
import { useParams } from "react-router-dom";
import TradeInput from '../components/tradeInput';

const CryptoView = () => {
  const styles = {
    container: {
      display: 'flex',
      justifyContent: 'center',
      flexDirection: 'column',
      padding: '20px',
      width: '100% - 100px',
      marginLeft: '100px',
      alignItems: 'center',
    },
    heading: {
      fontSize: '24px',
      marginBottom: '10px',
    },
    paragraph: {
      fontSize: '16px',
    },
  };

  const { crypto_code } = useParams();
  const cryptoCode = (crypto_code || "BTC").toUpperCase();
  var yahoo_code;
  var finnhub_code;
  var crypto_name;

  // Decode crypto_code for cross functionality with YAHOO Finance & Finnhub

  if (cryptoCode.includes("-USD")) {

    yahoo_code = cryptoCode;
    crypto_name = crypto_code.replace("-USD", "");
    finnhub_code = `BINANCE:${crypto_name}USDT`;

  } else if (cryptoCode.startsWith("BINANCE:") && crypto_code.endsWith("USDT")) {

    finnhub_code = cryptoCode;
    crypto_name = crypto_code.replace("BINANCE:", "").replace("USDT", "");
    yahoo_code = `${crypto_name}-USD`;

  } else {
    crypto_name = cryptoCode;
    yahoo_code = `${crypto_name}-USD`;
    finnhub_code = `BINANCE:${crypto_name}USDT`;
  }


  return (
    <div style={styles.container}>

      <CryptoGraph crypto_name={crypto_name} yahoo_code={yahoo_code} finnhub_code={finnhub_code} />
      <TradeInput type="crypto" code={crypto_name} finnhub_code={finnhub_code}/>
      
    </div>
    
  );
};

export default CryptoView;