import { API_URL } from "../constants";
import { Event } from "../types/events";

export async function getEvents(): Promise<Event[]> {
  const res = await fetch(`${API_URL}events`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
    cache: "no-store",
  });

  if (!res.ok) {
    const error = await res.json();
    throw new Error(error.error || "Failed to fetch events");
  }

  const data = await res.json();
  return data.events || [];
}
