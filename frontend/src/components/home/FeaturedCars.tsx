import React from 'react';
import Container from '../ui/Container';
import { featuredCars } from '../../data/cars';
import { motion } from 'framer-motion';
import { Fuel, Users, Zap } from 'lucide-react';

const FeaturedCars: React.FC = () => {
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
            Featured Premium Cars
          </motion.h2>
          <motion.p
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.6, delay: 0.1 }}
            viewport={{ once: true }}
            className="text-gray-400 max-w-2xl mx-auto"
          >
            Experience the pinnacle of automotive luxury with our handpicked selection of premium vehicles
          </motion.p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          {featuredCars.map((car, index) => (
            <motion.div
              key={car.id}
              initial={{ opacity: 0, y: 30 }}
              whileInView={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.5, delay: index * 0.1 }}
              viewport={{ once: true }}
              className="bg-gray-900 rounded-xl overflow-hidden hover:shadow-purple-900/20 hover:shadow-xl transition-all group"
            >
              {/* Car Image */}
              <div className="h-56 overflow-hidden">
                <img
                  src={car.image}
                  alt={car.name}
                  className="w-full h-full object-cover object-center group-hover:scale-105 transition-transform duration-500"
                />
              </div>

              {/* Car Details */}
              <div className="p-6">
                <div className="flex justify-between items-start mb-3">
                  <h3 className="text-xl font-bold">{car.name}</h3>
                  <div className="bg-purple-600 px-3 py-1 rounded-full text-sm font-medium">
                    €{car.price.daily}/day
                  </div>
                </div>

                <div className="text-gray-400 mb-6">{car.year} • Premium</div>

                <div className="grid grid-cols-3 gap-4 mb-6">
                  <div className="flex flex-col items-center text-center">
                    <Users size={20} className="text-purple-500 mb-1" />
                    <span className="text-sm text-gray-400">{car.seats} Seats</span>
                  </div>
                  <div className="flex flex-col items-center text-center">
                    <Zap size={20} className="text-purple-500 mb-1" />
                    <span className="text-sm text-gray-400">{car.transmission}</span>
                  </div>
                  <div className="flex flex-col items-center text-center">
                    <Fuel size={20} className="text-purple-500 mb-1" />
                    <span className="text-sm text-gray-400">{car.fuelType}</span>
                  </div>
                </div>

                <button className="w-full py-3 text-center border border-purple-600 text-purple-600 rounded-lg hover:bg-purple-600 hover:text-white transition-colors">
                  View Details
                </button>
              </div>
            </motion.div>
          ))}
        </div>

        <div className="text-center mt-12">
          <motion.button
            initial={{ opacity: 0, y: 10 }}
            whileInView={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.4, delay: 0.3 }}
            viewport={{ once: true }}
            className="inline-flex items-center text-purple-500 hover:text-purple-400 font-medium transition-colors"
          >
            View All Cars <span className="ml-2">→</span>
          </motion.button>
        </div>
      </Container>
    </section>
  );
};

export default FeaturedCars;