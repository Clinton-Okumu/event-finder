"use client";

import { Ticket } from "@/lib/types/tickets";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Calendar, MapPin, Ticket as TicketIcon, X } from "lucide-react";
import { useEffect, useState } from "react";
import { getUserTickets, cancelTicket } from "@/lib/api/tickets";
import { useAuth } from "@/contexts/AuthContext";
import { useRouter } from "next/navigation";
import Image from "next/image";

export default function MyTicketsPage() {
  const { user, loading } = useAuth();
  const router = useRouter();
  const [tickets, setTickets] = useState<Ticket[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [cancellingId, setCancellingId] = useState<number | null>(null);

  useEffect(() => {
    if (!loading && !user) {
      router.push("/");
      return;
    }

    if (user) {
      fetchTickets();
    }
  }, [user, loading, router]);

  const fetchTickets = async () => {
    try {
      const data = await getUserTickets();
      setTickets(data);
      setError(null);
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message.toLowerCase() : "";
      if (errorMessage.includes("no tickets") || errorMessage.includes("not found")) {
        setTickets([]);
        setError(null);
      } else {
        setError("Failed to load your tickets. Please try again.");
      }
    }
  };

  const handleCancelTicket = async (ticketId: number) => {
    if (!confirm("Are you sure you want to cancel this ticket?")) return;

    setCancellingId(ticketId);
    try {
      await cancelTicket(ticketId);
      setTickets(tickets.filter((t) => t.id !== ticketId));
    } catch (err) {
      setError("Failed to cancel ticket. Please try again.");
    } finally {
      setCancellingId(null);
    }
  };

  if (loading) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center py-12">
          <p className="text-muted-foreground">Loading...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="mb-8">
        <h1 className="text-3xl font-bold mb-2">My Tickets</h1>
        <p className="text-muted-foreground">
          View and manage your booked event tickets
        </p>
      </div>

      {error && (
        <div className="bg-destructive/10 text-destructive p-4 rounded-lg mb-6">
          {error}
        </div>
      )}

      {tickets.length === 0 ? (
        <Card className="shadow-lg">
          <CardContent className="flex flex-col items-center justify-center py-12">
            <TicketIcon className="w-16 h-16 text-muted-foreground mb-4" />
            <h3 className="text-lg font-semibold mb-2">No tickets yet</h3>
            <p className="text-muted-foreground mb-6">
              You haven't booked any events yet
            </p>
            <Button onClick={() => router.push("/")}>Browse Events</Button>
          </CardContent>
        </Card>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {tickets.map((ticket) => (
            <Card key={ticket.id} className="overflow-hidden shadow-lg">
              <div className="relative h-48 w-full">
                {ticket.event_image_url ? (
                  <Image
                    src={ticket.event_image_url}
                    alt={ticket.event_title}
                    fill
                    className="object-cover"
                  />
                ) : (
                  <div className="w-full h-full bg-muted flex items-center justify-center">
                    <TicketIcon className="w-16 h-16 text-muted-foreground" />
                  </div>
                )}
                <Badge className="absolute top-3 right-3 bg-primary">
                  Booked
                </Badge>
              </div>

              <CardHeader>
                <CardTitle className="text-xl">{ticket.event_title}</CardTitle>
              </CardHeader>

              <CardContent className="space-y-3">
                <div className="flex items-center gap-2 text-sm text-muted-foreground">
                  <Calendar className="w-4 h-4" />
                  <span>{ticket.event_date}</span>
                </div>

                <div className="flex items-center gap-2 text-sm text-muted-foreground">
                  <MapPin className="w-4 h-4" />
                  <span>{ticket.event_location}</span>
                </div>

                <div className="pt-2 border-t">
                  <div className="flex justify-between items-center">
                    <div>
                      <p className="text-xs text-muted-foreground">Amount Paid</p>
                      <p className="font-bold text-primary text-lg">
                        ksh.{ticket.price_paid}
                      </p>
                    </div>
                    <Button
                      variant="outline"
                      size="sm"
                      onClick={() => handleCancelTicket(ticket.id)}
                      disabled={cancellingId === ticket.id}
                    >
                      {cancellingId === ticket.id ? (
                        "Cancelling..."
                      ) : (
                        <>
                          <X className="w-4 h-4 mr-1" />
                          Cancel
                        </>
                      )}
                    </Button>
                  </div>
                </div>

                {ticket.qr_code && (
                  <div className="pt-2 border-t">
                    <p className="text-xs text-muted-foreground mb-2">QR Code</p>
                    <div className="bg-white p-3 rounded-lg inline-block border">
                      <Image
                        src={ticket.qr_code}
                        alt="Ticket QR Code"
                        width={120}
                        height={120}
                        className="w-32 h-32"
                      />
                    </div>
                  </div>
                )}
              </CardContent>
            </Card>
          ))}
        </div>
      )}
    </div>
  );
}
