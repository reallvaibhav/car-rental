
import Container from '../../components/ui/Container';

export default function FleetOwnerDashboard() {
  return (
    <Container>
      <div className="py-12">
        <h1 className="text-4xl font-bold mb-8">Fleet Owner Dashboard</h1>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div className="bg-gray-800 p-6 rounded-lg">
            <h2 className="text-2xl font-bold mb-4">My Fleet</h2>
            {/* Add fleet management content */}
          </div>
          <div className="bg-gray-800 p-6 rounded-lg">
            <h2 className="text-2xl font-bold mb-4">Rental Requests</h2>
            {/* Add rental requests content */}
          </div>
        </div>
      </div>
    </Container>
  );
}