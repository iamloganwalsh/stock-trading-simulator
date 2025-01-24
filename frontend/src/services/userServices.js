import api from './api';

const createUser = async (createUsername) => {
    const response = await api.post('/user/create', {
        username: createUsername,
    })
    return response.data;
}

const getUsers = async () => {
    const response = await api.get('/user/username');
    return response.data;
}

const getBalance = async () => {
    const response = await api.get('/user/balance');
    return response.data;
}

export default {
    createUser,
    getUsers,
    getBalance,
}