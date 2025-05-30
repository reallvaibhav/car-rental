import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import Container from '../ui/Container';
import { Menu, X, ChevronDown, Car, LogOut } from 'lucide-react';
import { cn } from '../../utils/cn';
import { useAuth } from '../../contexts/AuthContext';
import Button from '../ui/Button';

const Header: React.FC = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const [isScrolled, setIsScrolled] = useState(false);
  const { user, signInWithGoogle, logout } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    const handleScroll = () => {
      setIsScrolled(window.scrollY > 50);
    };

    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  const handleLogout = async () => {
    await logout();
    navigate('/');
  };

  return (
    <header
      className={cn(
        'fixed top-0 left-0 right-0 z-50 transition-all duration-300',
        isScrolled ? 'bg-black bg-opacity-90 backdrop-blur-sm py-3' : 'bg-transparent py-4'
      )}
    >
      <Container>
        <div className="flex items-center justify-between">
          <Link to="/" className="text-3xl font-bold text-white flex items-center gap-2">
            <Car className="h-8 w-8 text-purple-500" />
            <span className="bg-gradient-to-r from-white to-purple-400 bg-clip-text text-transparent">
              CAR-GO
            </span>
          </Link>

          {/* Desktop Navigation */}
          <nav className="hidden md:flex items-center space-x-8">
            <Link to="/" className="text-white hover:text-purple-400 transition-colors">
              Home
            </Link>
            <div className="relative group">
              <button className="flex items-center text-white hover:text-purple-400 transition-colors">
                Destinations
                <ChevronDown className="ml-1 h-4 w-4" />
              </button>
              <div className="absolute left-0 mt-2 w-48 bg-black bg-opacity-90 backdrop-blur-sm rounded-md shadow-lg py-2 opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-300">
                <Link
                  to="/destinations/downtown"
                  className="block px-4 py-2 text-white hover:bg-purple-800 transition-colors"
                >
                  Downtown
                </Link>
                <Link
                  to="/destinations/expo"
                  className="block px-4 py-2 text-white hover:bg-purple-800 transition-colors"
                >
                  EXPO
                </Link>
                <Link
                  to="/destinations/airport"
                  className="block px-4 py-2 text-white hover:bg-purple-800 transition-colors"
                >
                  Airport
                </Link>
              </div>
            </div>
            <Link
              to="/popular-routes"
              className="text-white hover:text-purple-400 transition-colors"
            >
              Popular Routes
            </Link>
            <Link
              to="/listings"
              className="text-white hover:text-purple-400 transition-colors"
            >
              Car Listings
            </Link>
            <Link to="/about" className="text-white hover:text-purple-400 transition-colors">
              About
            </Link>
          </nav>

          {/* Auth Section */}
          <div className="hidden md:flex items-center space-x-4">
            {user ? (
              <div className="flex items-center space-x-4">
                <Button variant="primary\" size="sm\" onClick={() => navigate('/book')}>
                  Book Now
                </Button>
                <div className="relative group">
                  <button className="flex items-center space-x-2 text-white">
                    <img
                      src={user.photoURL || 'https://via.placeholder.com/32'}
                      alt={user.displayName || 'User'}
                      className="w-8 h-8 rounded-full"
                    />
                    <ChevronDown className="h-4 w-4" />
                  </button>
                  <div className="absolute right-0 mt-2 w-48 bg-black bg-opacity-90 backdrop-blur-sm rounded-md shadow-lg py-2 opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-300">
                    <button
                      onClick={handleLogout}
                      className="w-full px-4 py-2 text-white hover:bg-purple-800 transition-colors flex items-center"
                    >
                      <LogOut className="h-4 w-4 mr-2" />
                      Sign Out
                    </button>
                  </div>
                </div>
              </div>
            ) : (
              <div className="flex items-center space-x-4">
                <Button variant="primary" size="sm" onClick={() => navigate('/book')}>
                  Book Now
                </Button>
                <button
                  onClick={() => signInWithGoogle()}
                  className="bg-white text-gray-900 px-3 py-1.5 rounded-lg text-sm hover:bg-gray-100 transition-colors flex items-center space-x-2"
                >
                  <img
                    src="https://www.google.com/favicon.ico"
                    alt="Google"
                    className="w-4 h-4"
                  />
                  <span>Sign in</span>
                </button>
              </div>
            )}
          </div>

          {/* Mobile Menu Button */}
          <button className="md:hidden text-white" onClick={() => setIsMenuOpen(!isMenuOpen)}>
            {isMenuOpen ? <X className="h-6 w-6" /> : <Menu className="h-6 w-6" />}
          </button>
        </div>

        {/* Mobile Navigation */}
        {isMenuOpen && (
          <div className="md:hidden mt-4 pb-4 space-y-4">
            <Link
              to="/"
              className="block text-white hover:text-purple-400 py-2 transition-colors"
            >
              Home
            </Link>
            <div className="space-y-2">
              <button className="flex items-center w-full text-white hover:text-purple-400 py-2 transition-colors">
                Destinations
                <ChevronDown className="ml-1 h-4 w-4" />
              </button>
              <div className="pl-4 space-y-2">
                <Link
                  to="/destinations/downtown"
                  className="block text-white hover:text-purple-400 py-1 transition-colors"
                >
                  Downtown
                </Link>
                <Link
                  to="/destinations/expo"
                  className="block text-white hover:text-purple-400 py-1 transition-colors"
                >
                  EXPO
                </Link>
                <Link
                  to="/destinations/airport"
                  className="block text-white hover:text-purple-400 py-1 transition-colors"
                >
                  Airport
                </Link>
              </div>
            </div>
            <Link
              to="/popular-routes"
              className="block text-white hover:text-purple-400 py-2 transition-colors"
            >
              Popular Routes
            </Link>
            <Link
              to="/listings"
              className="block text-white hover:text-purple-400 py-2 transition-colors"
            >
              Car Listings
            </Link>
            <Link
              to="/about"
              className="block text-white hover:text-purple-400 py-2 transition-colors"
            >
              About
            </Link>
            <Button
              variant="primary"
              size="sm"
              fullWidth
              onClick={() => navigate('/book')}
              className="mb-2"
            >
              Book Now
            </Button>
            {user ? (
              <button
                onClick={handleLogout}
                className="w-full flex items-center text-white hover:text-purple-400 py-2 transition-colors"
              >
                <LogOut className="h-4 w-4 mr-2" />
                Sign Out
              </button>
            ) : (
              <button
                onClick={() => signInWithGoogle()}
                className="w-full bg-white text-gray-900 px-3 py-1.5 rounded-lg text-sm hover:bg-gray-100 transition-colors flex items-center justify-center space-x-2"
              >
                <img
                  src="https://www.google.com/favicon.ico"
                  alt="Google"
                  className="w-4 h-4"
                />
                <span>Sign in with Google</span>
              </button>
            )}
          </div>
        )}
      </Container>
    </header>
  );
}

export default Header;