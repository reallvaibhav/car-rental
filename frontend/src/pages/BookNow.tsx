
import BookingForm from '../components/home/BookingForm';
import Container from '../components/ui/Container';

export default function BookNow() {
  return (
    <Container>
      <div className="py-12">
        <h1 className="text-4xl font-bold text-center mb-8">Book Your Ride</h1>
        <BookingForm />
      </div>
    </Container>
  );
}