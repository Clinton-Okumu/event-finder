"use client";

import { Button } from "@/components/ui/button";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog";
import { Calendar, MapPin, Ticket, CheckCircle, Loader2 } from "lucide-react";
import Image from "next/image";
import { useAuth } from "@/contexts/AuthContext";
import { useRouter } from "next/navigation";
import { bookTicket } from "@/lib/api/tickets";
import { useState } from "react";

interface BookingModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  eventId: number;
  eventTitle: string;
  eventDate: string;
  eventLocation: string;
  eventImage: string;
  price?: number;
}

export default function BookingModal({
  open,
  onOpenChange,
  eventId,
  eventTitle,
  eventDate,
  eventLocation,
  eventImage,
  price,
}: BookingModalProps) {
  const { user } = useAuth();
  const router = useRouter();
  const [status, setStatus] = useState<"idle" | "loading" | "success" | "error">("idle");
  const [error, setError] = useState<string | null>(null);

  const handleBooking = async () => {
    if (!user) {
      onOpenChange(false);
      router.push("/");
      return;
    }

    setStatus("loading");
    setError(null);

    try {
      await bookTicket(eventId);
      setStatus("success");
      setTimeout(() => {
        onOpenChange(false);
        router.push("/tickets");
      }, 2000);
    } catch (err) {
      setStatus("error");
      setError(err instanceof Error ? err.message : "Failed to book ticket");
    }
  };

  const handleClose = () => {
    if (status === "loading") return;
    setStatus("idle");
    setError(null);
    onOpenChange(false);
  };

  return (
    <AlertDialog open={open} onOpenChange={handleClose}>
      <AlertDialogContent className="max-w-2xl">
        {status === "idle" && (
          <>
            <AlertDialogHeader>
              <AlertDialogTitle>Book Event Ticket</AlertDialogTitle>
              <AlertDialogDescription>
                {user ? (
                  "Confirm your booking details below"
                ) : (
                  "Please login to book this event"
                )}
              </AlertDialogDescription>
            </AlertDialogHeader>

            <div className="my-6">
              <div className="flex gap-4">
                <div className="relative w-32 h-32 shrink-0 rounded-lg overflow-hidden">
                  {eventImage ? (
                    <Image
                      src={eventImage}
                      alt={eventTitle}
                      fill
                      className="object-cover"
                    />
                  ) : (
                    <div className="w-full h-full bg-muted flex items-center justify-center">
                      <Ticket className="w-8 h-8 text-muted-foreground" />
                    </div>
                  )}
                </div>

                <div className="flex-1 space-y-2">
                  <h3 className="font-semibold text-lg">{eventTitle}</h3>
                  <div className="flex items-center gap-2 text-sm text-muted-foreground">
                    <Calendar className="w-4 h-4" />
                    <span>{eventDate}</span>
                  </div>
                  <div className="flex items-center gap-2 text-sm text-muted-foreground">
                    <MapPin className="w-4 h-4" />
                    <span>{eventLocation}</span>
                  </div>
                  {price !== undefined && (
                    <p className="font-bold text-primary text-lg">
                      ksh.{price}
                    </p>
                  )}
                </div>
              </div>
            </div>

            <AlertDialogFooter>
              <AlertDialogCancel>Cancel</AlertDialogCancel>
              <AlertDialogAction onClick={handleBooking}>
                {user ? "Confirm Booking" : "Login to Book"}
              </AlertDialogAction>
            </AlertDialogFooter>
          </>
        )}

        {status === "loading" && (
          <div className="flex flex-col items-center justify-center py-12">
            <Loader2 className="w-12 h-12 animate-spin text-primary mb-4" />
            <p className="text-muted-foreground">Processing your booking...</p>
          </div>
        )}

        {status === "success" && (
          <div className="flex flex-col items-center justify-center py-12">
            <CheckCircle className="w-16 h-16 text-green-500 mb-4" />
            <h3 className="text-xl font-semibold mb-2">Booking Confirmed!</h3>
            <p className="text-muted-foreground">
              Your ticket has been booked successfully
            </p>
            <p className="text-sm text-muted-foreground mt-2">
              Redirecting to your tickets...
            </p>
          </div>
        )}

        {status === "error" && (
          <>
            <AlertDialogHeader>
              <AlertDialogTitle>Booking Failed</AlertDialogTitle>
              <AlertDialogDescription>
                {error || "An error occurred while booking your ticket"}
              </AlertDialogDescription>
            </AlertDialogHeader>

            <AlertDialogFooter>
              <AlertDialogCancel onClick={() => setStatus("idle")}>
                Try Again
              </AlertDialogCancel>
            </AlertDialogFooter>
          </>
        )}
      </AlertDialogContent>
    </AlertDialog>
  );
}
