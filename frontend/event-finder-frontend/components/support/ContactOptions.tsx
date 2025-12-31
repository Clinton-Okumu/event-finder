import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Mail, Phone, Clock, ArrowRight } from "lucide-react";

export default function ContactOptions() {
  const contactMethods = [
    {
      icon: Mail,
      title: "Email Support",
      description: "support@eventfinder.com",
      action: "Email Us",
    },
    {
      icon: Phone,
      title: "Phone Support",
      description: "+254 7XX XXX XXX",
      action: "Call Us",
    },
    {
      icon: Clock,
      title: "Working Hours",
      description: "Mon - Fri: 9AM - 6PM",
      action: "View Hours",
    },
  ];

  return (
    <Card className="shadow-lg">
      <CardHeader>
        <CardTitle>Other Ways to Reach Us</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="space-y-4">
          {contactMethods.map((method, index) => (
            <div key={index} className="flex items-center justify-between p-4 border rounded-lg">
              <div className="flex items-center gap-4">
                <div className="p-2 bg-primary/10 rounded-full">
                  <method.icon className="w-5 h-5 text-primary" />
                </div>
                <div>
                  <p className="font-semibold">{method.title}</p>
                  <p className="text-sm text-muted-foreground">{method.description}</p>
                </div>
              </div>
              <Button variant="outline" size="sm">
                {method.action}
                <ArrowRight className="ml-2 w-4 h-4" />
              </Button>
            </div>
          ))}
        </div>
      </CardContent>
    </Card>
  );
}
