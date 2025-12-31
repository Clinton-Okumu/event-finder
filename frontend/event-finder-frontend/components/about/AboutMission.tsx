import { Card, CardContent } from "@/components/ui/card";
import { Target, Users, TrendingUp } from "lucide-react";

export default function AboutMission() {
  const values = [
    {
      icon: Target,
      title: "Our Mission",
      description:
        "To make it easy for everyone to discover, book, and experience amazing events in their area.",
    },
    {
      icon: Users,
      title: "Our Vision",
      description:
        "To be the leading event discovery platform that brings communities together through shared experiences.",
    },
    {
      icon: TrendingUp,
      title: "Our Values",
      description:
        "We believe in transparency, accessibility, and creating memorable experiences for every user.",
    },
  ];

  return (
    <section className="py-16">
      <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
        {values.map((value, index) => (
          <Card key={index} className="text-center shadow-lg">
            <CardContent className="pt-6">
              <div className="flex justify-center mb-4">
                <div className="p-4 bg-primary/10 rounded-full">
                  <value.icon className="w-8 h-8 text-primary" />
                </div>
              </div>
              <h3 className="text-xl font-semibold mb-3">{value.title}</h3>
              <p className="text-muted-foreground">{value.description}</p>
            </CardContent>
          </Card>
        ))}
      </div>
    </section>
  );
}
