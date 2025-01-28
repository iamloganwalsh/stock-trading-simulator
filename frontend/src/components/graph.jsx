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

const CryptoGraph = ({crypto_name, yahoo_code, finnhub_code}) => {
  const [prices, setPrices] = useState([]);
  const [timestamps, setTimestamps] = useState([]);

  const updateGraph = (price) => {
    const currentTime = Date.now();
    
    setPrices((prev) => [...prev, price]);
    setTimestamps((prev) => [...prev, currentTime]); // Store timestamp in ms
  };

  useEffect(() => {
    const fetchInitialData = async () => {
      const response = await fetchingServices.fetchCryptoPrevPrice(yahoo_code);
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
            filteredTimestamps.push(Date.now()); // Start with the current time
          }
        });

        console.log("Filtered Prices:", filteredData);
        console.log("Filtered Timestamps:", filteredTimestamps);

        // Adjust the timestamps to go backward in time (each timestamp 10 seconds apart)
        const adjustedTimestamps = filteredTimestamps.map((timestamp, index) => {
          return timestamp - (filteredTimestamps.length - 1 - index) * 10000; // Subtract 10s for each step
        });

        console.log("Adjusted Timestamps:", adjustedTimestamps);

        // Set the filtered data and adjusted timestamps to the state
        setPrices(filteredData);
        setTimestamps(adjustedTimestamps);
      } else {
        console.log("No initial data fetched.");
      }
    };

    fetchInitialData();

    const fetchPrice = async () => {
      const data = await fetchingServices.fetchCryptoPrice(finnhub_code);
      console.log("New Price Data:", data);
      updateGraph(data);
    };

    // Setting up the interval to fetch prices every 10 seconds
    const interval = setInterval(fetchPrice, 10000);

    return () => clearInterval(interval);
  }, []); // The dependency array is empty to ensure this runs only once

  const options = {
    responsive: true,
    plugins: {
      legend: {
        display: false,
      },
      title: {
        display: true,
        text: `${crypto_name} Price Chart`,
      },
    },
  };

  const data = {
    labels: timestamps.map((timestamp) => new Date(timestamp).toLocaleTimeString()), // Convert timestamps to readable format
    datasets: [
      {
        label: crypto_name,
        data: prices,
        borderColor: 'rgb(255, 99, 132)',
        backgroundColor: 'rgba(255, 99, 132, 0.5)',
      },
    ],
  };

  return <Line options={options} data={data} />;
};

export default CryptoGraph;
