import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";

export default function FAQSection() {
  const faqs = [
    {
      question: "How do I book an event?",
      answer:
        "Simply browse our events page, select the event you're interested in, and click 'Get Tickets'. Follow the booking process to complete your purchase.",
    },
    {
      question: "Can I cancel my booking?",
      answer:
        "Yes, you can cancel your ticket from the 'My Tickets' page. Note that cancellation policies vary by event, so please check the event details before booking.",
    },
    {
      question: "How do I get my ticket?",
      answer:
        "After booking, your ticket will appear in the 'My Tickets' section. You'll receive a QR code that you can present at the event venue.",
    },
    {
      question: "What payment methods do you accept?",
      answer:
        "We accept M-Pesa, credit cards, and bank transfers. All transactions are secure and processed through our trusted payment partners.",
    },
  ];

  return (
    <Card className="shadow-lg">
      <CardHeader>
        <CardTitle>Frequently Asked Questions</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="space-y-4">
          {faqs.map((faq, index) => (
            <div key={index} className="border-b pb-4 last:border-0">
              <div className="flex items-start gap-3 mb-2">
                <Badge variant="secondary" className="mt-0.5 shrink-0">
                  {index + 1}
                </Badge>
                <h4 className="font-semibold flex-1">{faq.question}</h4>
              </div>
              <p className="text-sm text-muted-foreground ml-9">{faq.answer}</p>
            </div>
          ))}
        </div>
      </CardContent>
    </Card>
  );
}
