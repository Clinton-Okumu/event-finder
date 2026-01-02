import { useState } from "react";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Calendar, MapPin, Minus, Plus, Ticket } from "lucide-react";

interface BookingWidgetProps {
  price?: number;
  eventTitle: string;
  eventId: number;
}

export default function BookingWidget({ price, eventTitle, eventId }: BookingWidgetProps) {
  const [quantity, setQuantity] = useState(1);
  const totalPrice = price ? price * quantity : 0;

  const handleQuantityChange = (newQuantity: number) => {
    if (newQuantity >= 1 && newQuantity <= 10) {
      setQuantity(newQuantity);
    }
  };

  return (
    <Card className="shadow-lg sticky top-6">
      <CardContent className="p-6 space-y-6">
        <div>
          <h3 className="text-2xl font-bold mb-2">Book Tickets</h3>
          <p className="text-muted-foreground text-sm">{eventTitle}</p>
        </div>

        <div className="space-y-4">
          <div className="flex items-center justify-between p-4 bg-muted rounded-lg">
            <div className="flex items-center gap-2">
              <Ticket className="w-5 h-5 text-primary" />
              <span className="font-semibold">Quantity</span>
            </div>
            <div className="flex items-center gap-3">
              <Button
                variant="outline"
                size="icon"
                onClick={() => handleQuantityChange(quantity - 1)}
                disabled={quantity <= 1}
              >
                <Minus className="w-4 h-4" />
              </Button>
              <Input
                type="number"
                value={quantity}
                onChange={(e) => handleQuantityChange(parseInt(e.target.value) || 1)}
                min={1}
                max={10}
                className="w-20 text-center"
              />
              <Button
                variant="outline"
                size="icon"
                onClick={() => handleQuantityChange(quantity + 1)}
                disabled={quantity >= 10}
              >
                <Plus className="w-4 h-4" />
              </Button>
            </div>
          </div>

          {price !== undefined && (
            <div className="space-y-2">
              <div className="flex justify-between text-sm">
                <span className="text-muted-foreground">Price per ticket</span>
                <span className="font-medium">ksh.{price}</span>
              </div>
              <div className="flex justify-between text-lg font-bold border-t pt-2">
                <span>Total Amount</span>
                <span className="text-primary">ksh.{totalPrice}</span>
              </div>
            </div>
          )}
        </div>

        <Button size="lg" className="w-full">
          Pay with M-Pesa
        </Button>

        <div className="text-xs text-muted-foreground text-center">
          By clicking "Pay with M-Pesa", you agree to our Terms of Service and Privacy Policy
        </div>
      </CardContent>
    </Card>
  );
}
