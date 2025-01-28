import cs from '../services/cryptoServices.js'

const Buy = async (type, code, price, amount) => {
    if (type == 'crypto') {
        const response = await cs.BuyCrypto(code, price, amount);
        console.log(response);
        if (response.Status !== HttpStatusCode.OK) {
            alert("error")
        } else {
            alert("success")
        }
    }
}

const Sell = (type, code, price, amount) => {

}

// Determine if stock or crypto on view page
const BuySellDetails = (type) => {

}

export default BuySellDetails;