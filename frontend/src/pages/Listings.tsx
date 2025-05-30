import React, { useState } from 'react';
import Container from '../components/ui/Container';
import { cars } from '../data/cars';
import { motion } from 'framer-motion';
import { Car as CarIcon, Filter } from 'lucide-react';

const categories = ['All', 'Economy', 'Luxury', 'SUV', 'Sports', 'Electric'];
const priceRanges = [
  { min: 0, max: 50, label: 'Under €50' },
  { min: 50, max: 100, label: '€50 - €100' },
  { min: 100, max: 150, label: '€100 - €150' },
  { min: 150, max: Infinity, label: 'Over €150' },
];

export default function Listings() {
  const [selectedCategory, setSelectedCategory] = useState('All');
  const [selectedPriceRange, setSelectedPriceRange] = useState<{ min: number; max: number }>(
    { min: 0, max: Infinity }
  );

  const filteredCars = cars.filter((car) => {
    const matchesCategory = selectedCategory === 'All' || 
      car.category.toLowerCase() === selectedCategory.toLowerCase();
    const matchesPrice = car.price.daily >= selectedPriceRange.min && 
      car.price.daily <= selectedPriceRange.max;
    return matchesCategory && matchesPrice;
  });

  return (
    <div className="py-24 bg-gray-950">
      <Container>
        <div className="flex flex-col md:flex-row gap-8">
          {/* Filters */}
          <div className="w-full md:w-64 bg-gray-900 p-6 rounded-lg h-fit">
            <div className="flex items-center gap-2 mb-6">
              <Filter className="w-5 h-5 text-purple-500" />
              <h2 className="text-xl font-bold">Filters</h2>
            </div>

            <div className="space-y-6">
              {/* Category Filter */}
              <div>
                <h3 className="font-medium mb-3">Category</h3>
                <div className="space-y-2">
                  {categories.map((category) => (
                    <button
                      key={category}
                      onClick={() => setSelectedCategory(category)}
                      className={`w-full text-left px-3 py-2 rounded ${
                        selectedCategory === category
                          ? 'bg-purple-600 text-white'
                          : 'hover:bg-gray-800'
                      }`}
                    >
                      {category}
                    </button>
                  ))}
                </div>
              </div>

              {/* Price Range Filter */}
              <div>
                <h3 className="font-medium mb-3">Price Range</h3>
                <div className="space-y-2">
                  {priceRanges.map((range) => (
                    <button
                      key={range.label}
                      onClick={() => setSelectedPriceRange(range)}
                      className={`w-full text-left px-3 py-2 rounded ${
                        selectedPriceRange.min === range.min
                          ? 'bg-purple-600 text-white'
                          : 'hover:bg-gray-800'
                      }`}
                    >
                      {range.label}
                    </button>
                  ))}
                </div>
              </div>
            </div>
          </div>

          {/* Car Grid */}
          <div className="flex-1">
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {filteredCars.map((car, index) => (
                <motion.div
                  key={car.id}
                  initial={{ opacity: 0, y: 20 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ duration: 0.3, delay: index * 0.1 }}
                  className="bg-gray-900 rounded-lg overflow-hidden hover:shadow-purple-500/10 hover:shadow-lg transition-all"
                >
                  <div className="relative h-48">
                    <img
                      src={car.image}
                      alt={car.name}
                      className="w-full h-full object-cover"
                    />
                    <div className="absolute top-4 right-4 bg-purple-600 px-3 py-1 rounded-full text-sm font-medium">
                      €{car.price.daily}/day
                    </div>
                  </div>
                  <div className="p-6">
                    <h3 className="text-xl font-bold mb-2">{car.name}</h3>
                    <div className="flex items-center gap-2 text-gray-400 mb-4">
                      <CarIcon className="w-4 h-4" />
                      <span>{car.category}</span>
                    </div>
                    <button className="w-full py-2 bg-purple-600 hover:bg-purple-700 rounded-lg transition-colors">
                      Book Now
                    </button>
                  </div>
                </motion.div>
              ))}
            </div>
          </div>
        </div>
      </Container>
    </div>
  );
}