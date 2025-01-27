import api from './api';

const fetchCryptoPrice = async (code) => {
    const response = await api.get(`/crypto/fetch/${code}`);
    return response.data;
}

const fetchCryptoPrevPrice = async (code) => {
    const response = await api.get(`/crypto/fetch_prev/${code}`);
    return response;
}

export default {
    fetchCryptoPrice,
    fetchCryptoPrevPrice,
}
