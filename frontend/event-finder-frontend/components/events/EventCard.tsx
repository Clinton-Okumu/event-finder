"use client";

import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import Image from "next/image";
export default function EventCard({
  id,
  title,
  date,
  image,
  location,
  description,
  price,
}: EventCardProps) {
  return (
    <Card className="group hover:shadow-lg transition-shadow duration-200">
      <div className="relative h-52 w-full overflow-hidden">
        <Image
          src={image}
          alt={title}
          fill
          className="object-cover group-hover:scale-105 transition-transform duration-200"
        />
        <Badge variant="secondary" className="absolute top-3 right-3 shadow-md">
          {date}
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
        <span className="text-primary font-bold">From ${price}</span>
        <Button size="sm">Get Tickets</Button>
      </CardFooter>
    </Card>
  );
}
