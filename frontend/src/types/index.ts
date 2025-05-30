export interface Car {
  id: string;
  name: string;
  make: string;
  model: string;
  year: number;
  image: string;
  category: 'economy' | 'luxury' | 'suv' | 'sports' | 'electric';
  price: {
    daily: number;
    weekly: number;
    monthly: number;
  };
  seats: number;
  transmission: 'automatic' | 'manual';
  fuelType: 'gasoline' | 'diesel' | 'electric' | 'hybrid';
  features: string[];
  available: boolean;
}

export interface Booking {
  pickupLocation: string;
  dropoffLocation: string;
  pickupDate: string;
  dropoffDate: string;
  carId: string;
}

export interface Testimonial {
  id: string;
  name: string;
  avatar: string;
  rating: number;
  comment: string;
  date: string;
}

export interface Location {
  id: string;
  name: string;
  address: string;
  city: string;
  availableCars: number;
}

export interface User {
  id: string;
  email: string;
  name: string;
  displayName?: string;
  photoURL?: string;
}