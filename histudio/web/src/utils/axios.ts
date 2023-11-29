/**
 * axios setup to use mock service
 */

import axios from 'axios';

import env from "../environment"

const axiosServices = axios.create({ baseURL: env?.API_URL || 'http://localhost:3010/' });

// interceptor for http
axiosServices.interceptors.response.use(
    (response) => response,
    (error) => Promise.reject((error.response && error.response.data) || 'Wrong Services')
);

export default axiosServices;
