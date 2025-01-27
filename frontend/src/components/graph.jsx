import { Line } from 'react-chartjs-2';
import { useState, useEffect } from 'react';
import React from 'react';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
);
import fetchingServices from '../services/fetchingServices'

const CryptoGraph = () => {
  const [prices, setPrices] = useState([]);
  const [timestamps, setTimestamps] = useState([]);

  const updateGraph = (price) => {
    setPrices((prev) => [...prev, price]);
    setTimestamps((prev) => [...prev, new Date().toLocaleTimeString()]);
  };

  useEffect(() => {
    const fetchPrice = async () => {
      const data = await fetchingServices.fetchCryptoPrice("BINANCE:BTCUSDT");
      updateGraph(data);
    };

    const interval = setInterval(fetchPrice, 5000);
    return () => clearInterval(interval);
  }, []);


  const options = {
    responsive: true,
    plguins: {
        legend: {
            position: 'top',
        },
        title: {
            display: true,
            text: 'cool chart',
        },
    },
  };

  const data = {
    labels: timestamps, // Add labels here for the x-axis
    datasets: [
      {
        label: 'btc',
        data: prices,
        borderColor: 'rgb(255, 99, 132)',
        backgroundColor: 'rgba(255, 99, 132, 0.5)',
      },
    ],
  };
  


  return <Line options={options} data={data} />

};

export default CryptoGraph;
