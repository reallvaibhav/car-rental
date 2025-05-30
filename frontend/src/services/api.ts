import axios from 'axios';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:3000/api';

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token && config.headers) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export const auth = {
  login: async (email: string, password: string) => {
    const response = await api.post('/auth/login', { email, password });
    return response.data;
  },
  register: async (data: {
    email: string;
    password: string;
    role: 'customer' | 'fleet-owner';
    name: string;
  }) => {
    const response = await api.post('/auth/register', data);
    return response.data;
  },
};

export const bookings = {
  create: async (data: {
    pickupLocation: string;
    dropoffLocation: string;
    pickupDate: string;
    dropoffDate: string;
    carId: string;
  }) => {
    const response = await api.post('/bookings', data);
    return response.data;
  },
  getUserBookings: async () => {
    const response = await api.get('/bookings/user');
    return response.data;
  },
  getFleetOwnerBookings: async () => {
    const response = await api.get('/bookings/fleet-owner');
    return response.data;
  },
  cancel: async (bookingId: string) => {
    const response = await api.delete(`/bookings/${bookingId}`);
    return response.data;
  },
};

export const cars = {
  getAll: async () => {
    const response = await api.get('/cars');
    return response.data;
  },
  getById: async (id: string) => {
    const response = await api.get(`/cars/${id}`);
    return response.data;
  },
  create: async (data: FormData) => {
    const response = await api.post('/cars', data, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  },
  update: async (id: string, data: FormData) => {
    const response = await api.put(`/cars/${id}`, data, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  },
  delete: async (id: string) => {
    const response = await api.delete(`/cars/${id}`);
    return response.data;
  },
};

export default api;