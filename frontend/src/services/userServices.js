import api from './api';

const createUser = async (createUsername) => {
    const response = await api.post('/user/create', {
        username: createUsername,
    })
    return response.data;
}

const getUsername = async () => {
    const response = await api.get('/user/username');
    return response.data;
}

const getBalance = async () => {
    const response = await api.get('/user/balance');
    return response.data;
}

const getInvestment = async () => {
    const response = await api.get('/user/investment');
    return response.data;
}

const getCryptoPortfolio = async () => {
    const response = await api.get('/user/crypto_portfolio');
    return response.data;
}

const getStockPortfolio = async () => {
    const response = await api.get('/user/stock_portfolio');
    return response.data;
}

const getTradeHistory = async () => {
    const response = await api.get('/user/trade_history');
    return response.data;
}

export default {
    createUser,
    getUsername,
    getBalance,
    getInvestment,
    getCryptoPortfolio,
    getStockPortfolio,
    getTradeHistory,
}