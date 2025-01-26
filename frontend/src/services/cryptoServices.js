import api from './api';

const buyCrypto = async (code, cost, amount) => {
    const response = await api.post('/crypto/buy', {
        code: code,
        cost: cost,
        crypto_count: amount,
    })
    return response.data;
}

const sellCrypto = async (code, cost, amount) => {
    const response = await api.post('/crypto/sell', {
        code: code,
        cost: cost,
        crypto_count: amount,
    })
    return response.data
}

export default {
    buyCrypto,
    sellCrypto,
}