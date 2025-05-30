
import Container from '../components/ui/Container';

export default function About() {
  return (
    <Container>
      <div className="py-12">
        <h1 className="text-4xl font-bold text-center mb-8">About Us</h1>
        <div className="max-w-3xl mx-auto">
          <p className="text-lg mb-6">
            We are a premium car rental service dedicated to providing exceptional vehicles
            and outstanding customer service to our clients.
          </p>
          <p className="text-lg">
            With years of experience in the industry, we pride ourselves on our extensive
            fleet of luxury vehicles and our commitment to making your rental experience
            seamless and enjoyable.
          </p>
        </div>
      </div>
    </Container>
  );
}