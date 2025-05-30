import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { Toaster } from 'react-hot-toast';
import Header from './components/layout/Header';
import Home from './pages/Home';
import BookNow from './pages/BookNow';
import About from './pages/About';
import PremiumPlans from './pages/PremiumPlans';
import Listings from './pages/Listings';
import CustomerDashboard from './pages/dashboard/CustomerDashboard';
import FleetOwnerDashboard from './pages/dashboard/FleetOwnerDashboard';
import ProtectedRoute from './components/auth/ProtectedRoute';
import { AuthProvider } from './contexts/AuthContext';
import Footer from './components/layout/Footer';

function App() {
  return (
    <Router>
      <AuthProvider>
        <div className="min-h-screen bg-gradient-to-b from-gray-900 to-black text-white">
          <Header />
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/book" element={<BookNow />} />
            <Route path="/about" element={<About />} />
            <Route path="/plans" element={<PremiumPlans />} />
            <Route path="/listings" element={<Listings />} />
            <Route
              path="/dashboard/customer"
              element={
                <ProtectedRoute role="customer">
                  <CustomerDashboard />
                </ProtectedRoute>
              }
            />
            <Route
              path="/dashboard/fleet-owner"
              element={
                <ProtectedRoute role="fleet-owner">
                  <FleetOwnerDashboard />
                </ProtectedRoute>
              }
            />
          </Routes>
          <Footer />
          <Toaster position="top-right" />
        </div>
      </AuthProvider>
    </Router>
  );
}

export default App;