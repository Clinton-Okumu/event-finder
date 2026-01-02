import { Badge } from "@/components/ui/badge";
import { Card, CardContent } from "@/components/ui/card";
import { Calendar, MapPin, Clock, Users } from "lucide-react";
import Image from "next/image";
import { formatDate } from "@/lib/utils";

interface EventDetailsProps {
  event: {
    id: number;
    title: string;
    start_time: string;
    end_time: string;
    image_url: string;
    location: string;
    description: string;
    price?: number;
  };
}

export default function EventDetails({ event }: EventDetailsProps) {
  const formattedDate = formatDate(event.start_time);

  const startTime = new Date(event.start_time);
  const endTime = new Date(event.end_time);
  const formattedStartTime = startTime.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit' });
  const formattedEndTime = endTime.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit' });

  return (
    <div className="space-y-6">
      <div className="relative h-[400px] w-full rounded-xl overflow-hidden shadow-lg">
        {event.image_url ? (
          <Image
            src={event.image_url}
            alt={event.title}
            fill
            className="object-cover"
            priority
          />
        ) : (
          <div className="w-full h-full bg-muted flex items-center justify-center">
            <Calendar className="w-16 h-16 text-muted-foreground" />
          </div>
        )}
      </div>

      <Card className="shadow-lg">
        <CardContent className="p-6 space-y-6">
          <div>
            <Badge variant="secondary" className="mb-4">
              {formattedDate}
            </Badge>
            <h1 className="text-4xl font-bold mb-4">{event.title}</h1>

            <div className="flex flex-wrap gap-6 text-muted-foreground">
              <div className="flex items-center gap-2">
                <Calendar className="w-5 h-5" />
                <span>{formattedDate}</span>
              </div>
              <div className="flex items-center gap-2">
                <MapPin className="w-5 h-5" />
                <span>{event.location}</span>
              </div>
              <div className="flex items-center gap-2">
                <Clock className="w-5 h-5" />
                <span>{formattedStartTime} - {formattedEndTime}</span>
              </div>
              <div className="flex items-center gap-2">
                <Users className="w-5 h-5" />
                <span>21+ Event</span>
              </div>
            </div>
          </div>

          <div className="border-t pt-6">
            <h2 className="text-2xl font-semibold mb-4">About This Event</h2>
            <p className="text-muted-foreground leading-relaxed whitespace-pre-line">
              {event.description}
            </p>
          </div>

          <div className="border-t pt-6">
            <h2 className="text-xl font-semibold mb-4">Event Details</h2>
            <div className="space-y-3">
              <div className="flex justify-between">
                <span className="text-muted-foreground">Date</span>
                <span className="font-medium">{formattedDate}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-muted-foreground">Time</span>
                <span className="font-medium">{formattedStartTime} - {formattedEndTime}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-muted-foreground">Location</span>
                <span className="font-medium">{event.location}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-muted-foreground">Age Restriction</span>
                <span className="font-medium">21+</span>
              </div>
              {event.price !== undefined && (
                <div className="flex justify-between">
                  <span className="text-muted-foreground">Starting From</span>
                  <span className="font-medium text-primary">ksh.{event.price}</span>
                </div>
              )}
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
