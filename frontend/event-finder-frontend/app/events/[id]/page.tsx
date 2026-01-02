"use client";

import { useEffect, useState, use } from "react";
import { getEventById } from "@/lib/api/events";
import EventDetails from "@/components/events/EventDetails";
import BookingWidget from "@/components/events/BookingWidget";
import { Event } from "@/lib/types/events";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";
import { ArrowLeft, Loader2 } from "lucide-react";

export default function EventDetailsPage({ params }: { params: Promise<{ id: string }> }) {
  const router = useRouter();
  const { id } = use(params);
  const [event, setEvent] = useState<Event | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchEvent = async () => {
      try {
        setLoading(true);
        const eventId = parseInt(id);
        const eventData = await getEventById(eventId);
        setEvent(eventData);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Failed to load event");
      } finally {
        setLoading(false);
      }
    };

    fetchEvent();
  }, [id]);

  if (loading) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="flex flex-col items-center justify-center py-20">
          <Loader2 className="w-12 h-12 animate-spin text-primary mb-4" />
          <p className="text-muted-foreground">Loading event details...</p>
        </div>
      </div>
    );
  }

  if (error || !event) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="max-w-md mx-auto">
          <Card className="text-center p-8 shadow-lg">
            <h2 className="text-2xl font-bold mb-2">Event Not Found</h2>
            <p className="text-muted-foreground mb-6">
              {error || "The event you're looking for doesn't exist or has been removed."}
            </p>
            <Button onClick={() => router.push("/events")}>
              <ArrowLeft className="w-4 h-4 mr-2" />
              Back to Events
            </Button>
          </Card>
        </div>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <Button
        variant="ghost"
        onClick={() => router.back()}
        className="mb-6"
      >
        <ArrowLeft className="w-4 h-4 mr-2" />
        Back
      </Button>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <div className="lg:col-span-2">
          <EventDetails event={event} />
        </div>

        <div className="lg:col-span-1">
          <BookingWidget
            price={event.price}
            eventTitle={event.title}
            eventId={event.id}
          />
        </div>
      </div>
    </div>
  );
}
