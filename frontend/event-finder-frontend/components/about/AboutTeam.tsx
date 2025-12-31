import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";

export default function AboutTeam() {
  const team = [
    {
      name: "John Doe",
      role: "Founder & CEO",
      description:
        "Visionary leader with 10+ years in event management and tech.",
    },
    {
      name: "Jane Smith",
      role: "Head of Product",
      description:
        "Product expert focused on creating seamless user experiences.",
    },
    {
      name: "Mike Johnson",
      role: "Lead Developer",
      description: "Full-stack developer passionate about building great products.",
    },
    {
      name: "Sarah Williams",
      role: "Community Manager",
      description: "Connecting with users and building our vibrant community.",
    },
  ];

  return (
    <section className="py-16">
      <div className="text-center mb-12">
        <h2 className="text-3xl font-bold mb-4">Meet Our Team</h2>
        <p className="text-muted-foreground text-lg">
          The passionate people behind Event Finder
        </p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        {team.map((member, index) => (
          <Card key={index} className="shadow-lg">
            <CardHeader>
              <div className="flex items-start justify-between">
                <CardTitle className="text-xl">{member.name}</CardTitle>
                <Badge variant="secondary">{member.role}</Badge>
              </div>
            </CardHeader>
            <CardContent>
              <p className="text-muted-foreground">{member.description}</p>
            </CardContent>
          </Card>
        ))}
      </div>
    </section>
  );
}
