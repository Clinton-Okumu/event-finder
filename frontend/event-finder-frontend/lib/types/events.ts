export interface Event {
  id: number;
  title: string;
  start_time: string;
  end_time: string;
  image_url: string;
  location: string;
  description: string;
  price?: number;
  status?: string;
  category_id?: number;
  total_tickets?: number;
  tickets_remaining?: number;
}
