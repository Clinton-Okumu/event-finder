import { Card, CardContent } from "@/components/ui/card";
import { CheckCircle2 } from "lucide-react";

export default function AboutStats() {
  const stats = [
    { value: "1000+", label: "Events Listed" },
    { value: "5000+", label: "Happy Users" },
    { value: "50+", label: "Event Categories" },
    { value: "24/7", label: "Support Available" },
  ];

  return (
    <section className="py-16">
      <div className="grid grid-cols-2 md:grid-cols-4 gap-6">
        {stats.map((stat, index) => (
          <Card key={index} className="text-center shadow-lg">
            <CardContent className="pt-6">
              <div className="flex justify-center mb-3">
                <CheckCircle2 className="w-8 h-8 text-primary" />
              </div>
              <div className="text-3xl font-bold text-primary mb-2">
                {stat.value}
              </div>
              <div className="text-sm text-muted-foreground">{stat.label}</div>
            </CardContent>
          </Card>
        ))}
      </div>
    </section>
  );
}
