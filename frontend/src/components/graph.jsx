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

import fetchingServices from '../services/fetchingServices';

const CryptoGraph = () => {
  const [prices, setPrices] = useState([]);
  const [timestamps, setTimestamps] = useState([]);

  const updateGraph = (price) => {
    setPrices((prev) => [...prev, price]);
    setTimestamps((prev) => [...prev, new Date().toLocaleTimeString()]);
  };

  useEffect(() => {
    const fetchInitialData = async () => {
      const response = await fetchingServices.fetchCryptoPrevPrice("BTC-USD");
      console.log("Initial Data Response:", response);

      if (response && response.data) {
        // Split data and process prices from string representation
        const initialData = response.data.split(' ').map(price => parseFloat(price));
        console.log("Processed Initial Prices:", initialData);

        // Filter out zero prices and maintain matching timestamps
        const filteredData = [];
        const filteredTimestamps = [];

        // Only start generating timestamps after filtering out invalid (0) data
        initialData.forEach((price, _) => {
          if (price !== 0 && !isNaN(price)) {
            filteredData.push(price);
            // Generate timestamps only for real data points
            filteredTimestamps.push(new Date().toLocaleTimeString());
          }
        });

        console.log("Filtered Prices:", filteredData);
        console.log("Filtered Timestamps:", filteredTimestamps);

        // Check if the lengths of prices and timestamps are the same
        if (filteredData.length !== filteredTimestamps.length) {

          const lengthDifference = filteredData.length - filteredTimestamps.length;

          // If the lengths are not the same, trim the oldest data from the longer array
          if (lengthDifference > 0) {
            setTimestamps((prev) => prev.slice(lengthDifference));
          } else if (lengthDifference < 0) {
            setPrices((prev) => prev.slice(-lengthDifference));
          }
        }

        // Set the filtered data to the state
        setPrices(filteredData);
        setTimestamps(filteredTimestamps);
      } else {
        console.log("No initial data fetched.");
      }
    };

    fetchInitialData();

    const fetchPrice = async () => {
      const data = await fetchingServices.fetchCryptoPrice("BINANCE:BTCUSDT");
      console.log("New Price Data:", data);
      updateGraph(data);
    };

    // Setting up the interval to fetch prices every 10 seconds
    const interval = setInterval(fetchPrice, 10000);

    return () => clearInterval(interval);
  }, []);

  const options = {
    responsive: true,
    plugins: {
      legend: {
        position: 'top',
      },
      title: {
        display: true,
        text: 'BTC Price Chart',
      },
    },
  };

  const data = {
    labels: timestamps,
    datasets: [
      {
        label: 'BTC',
        data: prices,
        borderColor: 'rgb(255, 99, 132)',
        backgroundColor: 'rgba(255, 99, 132, 0.5)',
      },
    ],
  };

  return <Line options={options} data={data} />;
};

export default CryptoGraph;
