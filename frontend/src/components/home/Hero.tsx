import React, { useState, useEffect } from 'react';
import Button from '../ui/Button';
import { motion } from 'framer-motion';

const Hero: React.FC = () => {
  const [currentTextIndex, setCurrentTextIndex] = useState(0);
  const locations = ['ASTANA', 'EXPO', 'AIRPORT', 'DOWNTOWN'];

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentTextIndex((prevIndex) => (prevIndex + 1) % locations.length);
    }, 3000);

    return () => clearInterval(interval);
  }, []);

  return (
    <div className="relative min-h-screen flex items-center justify-center overflow-hidden bg-gradient-to-b from-black via-gray-900 to-black">
      

      {/* Animated gradient orbs */}
      <div className="absolute inset-0 overflow-hidden pointer-events-none z-10">
        <div className="absolute -top-40 -right-40 w-80 h-80 bg-purple-600/30 rounded-full blur-3xl animate-pulse" />
        <div className="absolute -bottom-40 -left-40 w-80 h-80 bg-purple-800/20 rounded-full blur-3xl animate-pulse" />
      </div>

      {/* Content */}
      <div className="container mx-auto px-4 sm:px-6 lg:px-8 relative z-20 text-white text-center mb-32">
        <motion.h1
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8 }}
          className="text-4xl md:text-6xl lg:text-7xl font-bold mb-6 bg-clip-text text-transparent bg-gradient-to-r from-white to-purple-400"
        >
          PREMIUM CARS
          <br />
          <div className="h-20">
            <motion.span
              key={currentTextIndex}
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              exit={{ opacity: 0, y: -20 }}
              transition={{ duration: 0.5 }}
              className="inline-block bg-gradient-to-r from-purple-400 to-purple-600 bg-clip-text text-transparent"
            >
              IN {locations[currentTextIndex]}
            </motion.span>
          </div>
          <span className="bg-gradient-to-r from-white to-purple-200 bg-clip-text text-transparent">NOW</span>
        </motion.h1>

        <motion.p
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8, delay: 0.2 }}
          className="text-gray-300 max-w-2xl mx-auto mb-8 text-lg"
        >
          Experience luxury driving in the heart of Kazakhstan. From city tours to airport transfers, we offer
          24/7 support and free cancellation.
        </motion.p>

        {/* 24/7 badge */}
        <motion.div
          initial={{ opacity: 0, scale: 0.8 }}
          animate={{ opacity: 1, scale: 1 }}
          transition={{ duration: 0.5, delay: 0.4 }}
          className="inline-block bg-gradient-to-r from-purple-600 to-purple-800 text-white text-sm font-bold px-6 py-2 rounded-full mb-8 shadow-lg shadow-purple-500/20"
        >
          24/7 AVAILABILITY
        </motion.div>

        <div className="mt-8">
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8, delay: 0.6 }}
            className="space-x-4"
          >
            <Button
              variant="primary"
              size="lg"
              className="px-10 py-5 text-lg rounded-full shadow-xl shadow-purple-500/20 hover:shadow-purple-500/40"
            >
              Book my ride →
            </Button>
            <Button
              variant="outline"
              size="lg"
              className="px-10 py-5 text-lg rounded-full shadow-xl"
            >
              View Fleet
            </Button>
          </motion.div>
        </div>
      </div>

      {/* Car image - positioned at the bottom */}
      <div className="absolute bottom-0 left-0 right-0 z-10">
        <div className="relative max-w-6xl mx-auto">
          <div className="absolute inset-0 bg-gradient-to-t from-black via-transparent to-transparent" />
          
        </div>
      </div>
    </div>
  );
};

export default Hero;