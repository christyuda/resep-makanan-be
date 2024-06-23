const express = require('express');
const router = express.Router();
const { createOrder, getOrdersByUserId } = require('../controllers/orderController.js');
const verifyToken = require('../middleware/auth.js');


router.post('/', verifyToken, createOrder);
router.get('/', verifyToken, getOrdersByUserId)



module.exports = router;