"use client";

import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import Image from "next/image";
import Link from "next/link";
import { formatDate } from "@/lib/utils";

interface EventCardProps {
  id: number;
  title: string;
  start_time: string;
  image_url: string;
  location: string;
  description: string;
  price?: number;
}

export default function EventCard({
  id,
  title,
  start_time,
  image_url,
  location,
  description,
  price,
}: EventCardProps) {
  const formattedDate = formatDate(start_time);

  return (
    <Card className="group shadow-lg hover:shadow-xl transition-shadow duration-200">
      <div className="relative h-52 w-full overflow-hidden">
        {image_url ? (
          <Image
            src={image_url}
            alt={title}
            fill
            className="object-cover group-hover:scale-105 transition-transform duration-200"
          />
        ) : (
          <div className="w-full h-full bg-muted flex items-center justify-center">
            <span className="text-muted-foreground">No Image</span>
          </div>
        )}
        <Badge variant="secondary" className="absolute top-3 right-3 shadow-md">
          {formattedDate}
        </Badge>
      </div>

      <CardContent className="pt-5">
        <h3 className="text-lg font-semibold mb-1">{title}</h3>
        <p className="text-sm text-muted-foreground mb-2">{location}</p>
        <p className="text-sm text-muted-foreground line-clamp-2">
          {description}
        </p>
      </CardContent>

      <CardFooter className="flex items-center justify-between pt-0">
        <span className="text-primary font-bold">
          {price !== undefined ? `From ksh.${price}` : "Price TBD"}
        </span>
        <Link href={`/events/${id}`}>
          <Button size="sm">Get Tickets</Button>
        </Link>
      </CardFooter>
    </Card>
  );
}
