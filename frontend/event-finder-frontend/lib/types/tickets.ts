export interface Ticket {
  id: number;
  user_id: number;
  event_id: number;
  event_title: string;
  event_date: string;
  event_location: string;
  event_image_url: string;
  price_paid: number;
  booking_date: string;
  qr_code?: string;
}
