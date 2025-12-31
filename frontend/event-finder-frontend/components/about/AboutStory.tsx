import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

export default function AboutStory() {
  return (
    <section className="py-16">
      <Card className="shadow-lg">
        <CardHeader>
          <CardTitle className="text-3xl">Our Story</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4 text-lg">
          <p>
            Event Finder was born from a simple frustration: finding great events
            was harder than it should be. We noticed people missing out on amazing
            experiences simply because they didn't know where to look.
          </p>
          <p>
            In 2025, we set out to change that. Our platform brings together
            events of all types—from music festivals to workshops, from sports to
            cultural gatherings—in one easy-to-use place.
          </p>
          <p>
            Today, we're proud to connect thousands of event-goers with
            unforgettable experiences. But this is just the beginning. We're
            constantly improving our platform to make event discovery even easier
            and more personalized.
          </p>
          <p className="text-primary font-semibold">
            Join us in discovering what's happening around you!
          </p>
        </CardContent>
      </Card>
    </section>
  );
}
