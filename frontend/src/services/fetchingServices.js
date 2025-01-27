import api from './api';

const fetchCryptoPrice = async (code) => {
    const response = await api.get(`/crypto/fetch/${code}`);
    return response.data;
}

const fetchStockPrice = async (code) => {
    const response = await api.get(`/stock/fetch/${code}`);
    return response.data;
}

export default {
    fetchCryptoPrice,
    fetchStockPrice,
}
