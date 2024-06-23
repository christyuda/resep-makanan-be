const mongoose = require('mongoose');
const { Schema } = mongoose;

const userSchema = new Schema({
    fullName: { type: String, required: true },
    nickname: { type: String, required: true },
    email: { type: String, required: true, unique: true },
    phoneNumber: { type: String, required: true ,unique: true },
    password: { type: String, required: true },
    isActive: { type: String, default: true },
    role: { type: String, default: 'user' }
})

const User = mongoose.model('Users', userSchema);
module.exports = User;