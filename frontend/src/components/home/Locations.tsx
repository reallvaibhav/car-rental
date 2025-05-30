import React from 'react';
import Container from '../ui/Container';
import { locations } from '../../data/locations';
import { MapPin, Car } from 'lucide-react';
import { motion } from 'framer-motion';

const Locations: React.FC = () => {
  return (
    <section className="py-20 bg-gray-950 text-white">
      <Container>
        <div className="text-center mb-16">
          <motion.h2
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.6 }}
            viewport={{ once: true }}
            className="text-4xl font-bold mb-4"
          >
            Locations
          </motion.h2>
          <motion.p
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.6, delay: 0.1 }}
            viewport={{ once: true }}
            className="text-gray-400 max-w-2xl mx-auto"
          >
            Find premium car rental services across Astana
          </motion.p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          {locations.map((location, index) => (
            <motion.div
              key={location.id}
              initial={{ opacity: 0, y: 30 }}
              whileInView={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.5, delay: index * 0.1 }}
              viewport={{ once: true }}
              className="bg-gray-900 rounded-lg p-6 hover:shadow-purple-900/20 hover:shadow-lg transition-all group"
            >
              <div className="flex justify-between items-start mb-4">
                <h3 className="text-xl font-bold">{location.name}</h3>
                <div className="p-2 bg-purple-900/30 rounded-full group-hover:bg-purple-600/30 transition-colors">
                  <MapPin className="h-5 w-5 text-purple-500" />
                </div>
              </div>
              <p className="text-gray-400 mb-4">{location.address}</p>
              <div className="flex items-center text-sm text-gray-400">
                <Car className="h-4 w-4 mr-2 text-purple-500" />
                <span>{location.availableCars} cars available</span>
              </div>
              <button className="mt-6 w-full py-2 text-center border border-gray-800 text-gray-300 rounded hover:bg-gray-800 transition-colors">
                View Location
              </button>
            </motion.div>
          ))}
        </div>
      </Container>
    </section>
  );
};

export default Locations;