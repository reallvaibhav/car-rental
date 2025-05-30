
import Container from '../components/ui/Container';
import Button from '../components/ui/Button';

export default function PremiumPlans() {
  const plans = [
    {
      name: 'Basic',
      price: '$99/month',
      features: ['Access to standard fleet', '24/7 support', 'Free cancellation'],
    },
    {
      name: 'Premium',
      price: '$199/month',
      features: ['Access to luxury fleet', 'Priority booking', 'Concierge service'],
    },
    {
      name: 'Elite',
      price: '$299/month',
      features: ['Access to entire fleet', 'Personal driver option', 'VIP treatment'],
    },
  ];

  return (
    <Container>
      <div className="py-12">
        <h1 className="text-4xl font-bold text-center mb-8">Premium Plans</h1>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          {plans.map((plan) => (
            <div key={plan.name} className="bg-gray-800 p-6 rounded-lg">
              <h2 className="text-2xl font-bold mb-4">{plan.name}</h2>
              <p className="text-3xl font-bold mb-6">{plan.price}</p>
              <ul className="mb-6 space-y-2">
                {plan.features.map((feature) => (
                  <li key={feature} className="flex items-center">
                    <span className="mr-2">✓</span>
                    {feature}
                  </li>
                ))}
              </ul>
              <Button className="w-full">Subscribe Now</Button>
            </div>
          ))}
        </div>
      </div>
    </Container>
  );
}