import React, { useState } from 'react';
import Button from '../ui/Button';
import { Calendar, MapPin, Clock } from 'lucide-react';
import { locations } from '../../data/locations';
import { motion } from 'framer-motion';

const BookingForm: React.FC = () => {
  const [formData, setFormData] = useState({
    pickupLocation: '',
    dropoffLocation: '',
    pickupDate: '',
    dropoffDate: '',
  });

  const handleChange = (e: React.ChangeEvent<HTMLSelectElement | HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    console.log('Booking data:', formData);
    // Handle booking submission
  };

  return (
    <div className="relative py-16 bg-black">
      <div className="container mx-auto px-4 sm:px-6 lg:px-8">
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          whileInView={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.6 }}
          viewport={{ once: true }}
          className="bg-gray-900 rounded-2xl shadow-2xl p-6 md:p-10 max-w-5xl mx-auto"
        >
          <h2 className="text-3xl font-bold text-white mb-8 text-center">Book Your Premium Ride</h2>

          <form onSubmit={handleSubmit} className="space-y-6">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
              {/* Pickup Location */}
              <div className="space-y-2">
                <label htmlFor="pickupLocation" className="block text-gray-400 font-medium">
                  Pickup Location
                </label>
                <div className="relative">
                  <MapPin className="absolute left-3 top-1/2 transform -translate-y-1/2 text-purple-500" size={20} />
                  <select
                    id="pickupLocation"
                    name="pickupLocation"
                    value={formData.pickupLocation}
                    onChange={handleChange}
                    className="w-full pl-10 pr-4 py-3 bg-gray-800 border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500 text-white appearance-none"
                    required
                  >
                    <option value="" disabled>
                      Select pickup location
                    </option>
                    {locations.map((location) => (
                      <option key={location.id} value={location.id}>
                        {location.name}
                      </option>
                    ))}
                  </select>
                </div>
              </div>

              {/* Dropoff Location */}
              <div className="space-y-2">
                <label htmlFor="dropoffLocation" className="block text-gray-400 font-medium">
                  Dropoff Location
                </label>
                <div className="relative">
                  <MapPin className="absolute left-3 top-1/2 transform -translate-y-1/2 text-purple-500" size={20} />
                  <select
                    id="dropoffLocation"
                    name="dropoffLocation"
                    value={formData.dropoffLocation}
                    onChange={handleChange}
                    className="w-full pl-10 pr-4 py-3 bg-gray-800 border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500 text-white appearance-none"
                    required
                  >
                    <option value="" disabled>
                      Select dropoff location
                    </option>
                    {locations.map((location) => (
                      <option key={location.id} value={location.id}>
                        {location.name}
                      </option>
                    ))}
                  </select>
                </div>
              </div>

              {/* Pickup Date */}
              <div className="space-y-2">
                <label htmlFor="pickupDate" className="block text-gray-400 font-medium">
                  Pickup Date & Time
                </label>
                <div className="relative">
                  <Calendar className="absolute left-3 top-1/2 transform -translate-y-1/2 text-purple-500" size={20} />
                  <input
                    type="datetime-local"
                    id="pickupDate"
                    name="pickupDate"
                    value={formData.pickupDate}
                    onChange={handleChange}
                    className="w-full pl-10 pr-4 py-3 bg-gray-800 border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500 text-white"
                    required
                  />
                </div>
              </div>

              {/* Dropoff Date */}
              <div className="space-y-2">
                <label htmlFor="dropoffDate" className="block text-gray-400 font-medium">
                  Dropoff Date & Time
                </label>
                <div className="relative">
                  <Clock className="absolute left-3 top-1/2 transform -translate-y-1/2 text-purple-500" size={20} />
                  <input
                    type="datetime-local"
                    id="dropoffDate"
                    name="dropoffDate"
                    value={formData.dropoffDate}
                    onChange={handleChange}
                    className="w-full pl-10 pr-4 py-3 bg-gray-800 border border-gray-700 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500 text-white"
                    required
                  />
                </div>
              </div>
            </div>

            <div className="mt-8 text-center">
              <Button
                type="submit"
                variant="primary"
                size="lg"
                className="px-12 py-4 bg-purple-600 hover:bg-purple-700 rounded-lg transition-colors"
              >
                Search Available Cars
              </Button>
            </div>
          </form>
        </motion.div>
      </div>
    </div>
  );
};

export default BookingForm;