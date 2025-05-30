
import Container from '../../components/ui/Container';

export default function CustomerDashboard() {
  return (
    <Container>
      <div className="py-12">
        <h1 className="text-4xl font-bold mb-8">Customer Dashboard</h1>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div className="bg-gray-800 p-6 rounded-lg">
            <h2 className="text-2xl font-bold mb-4">Current Rentals</h2>
            {/* Add current rentals content */}
          </div>
          <div className="bg-gray-800 p-6 rounded-lg">
            <h2 className="text-2xl font-bold mb-4">Rental History</h2>
            {/* Add rental history content */}
          </div>
        </div>
      </div>
    </Container>
  );
}