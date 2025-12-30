"use client";

import Image from "next/image";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Event } from "@/lib/types/events";

export default function Events({ events }: { events: Event[] }) {
  if (events.length === 0) {
    return (
      <div className="text-center py-12">
        <p className="text-muted-foreground text-lg">No events found</p>
      </div>
    );
  }

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-8">
      {events.map((event) => (
        <Card key={event.id} className="group hover:shadow-lg transition-shadow duration-200">
          <div className="relative h-52 w-full overflow-hidden">
            <Image
              src={event.image}
              alt={event.title}
              fill
              className="object-cover group-hover:scale-105 transition-transform duration-200"
            />
            <Badge variant="secondary" className="absolute top-3 right-3 shadow-md">
              {event.date}
            </Badge>
          </div>

          <CardContent className="pt-5">
            <h3 className="text-lg font-semibold mb-1">{event.title}</h3>
            <p className="text-sm text-muted-foreground mb-2">{event.location}</p>
            <p className="text-sm text-muted-foreground line-clamp-2">
              {event.description}
            </p>
          </CardContent>

          <CardFooter className="flex items-center justify-between pt-0">
            <span className="text-primary font-bold">From ${event.price}</span>
            <Button size="sm">Get Tickets</Button>
          </CardFooter>
        </Card>
      ))}
    </div>
  );
}
