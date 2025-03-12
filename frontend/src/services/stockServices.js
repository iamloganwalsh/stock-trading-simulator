import api from './api';

const buyStock = async (code, cost, amount) => {
    const response = await api.post('/stock/buy', {
        code: code,
        cost: cost,
        stock_count: amount,
    })
    return response;
}

const sellStock = async (code, cost, amount) => {
    const response = await api.post('/stock/sell', {
        code: code,
        cost: cost,
        stock_count: amount,
    })
    return response;
}

export default {
    buyStock,
    sellStock,
}