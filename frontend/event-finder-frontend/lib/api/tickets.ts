import { API_URL } from "../constants";
import { Ticket } from "../types/tickets";

function getAuthHeaders() {
  const token = localStorage.getItem("token");
  return {
    "Content-Type": "application/json",
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
  };
}

export async function bookTicket(eventId: number): Promise<Ticket> {
  const res = await fetch(`${API_URL}tickets`, {
    method: "POST",
    headers: getAuthHeaders(),
    body: JSON.stringify({ event_id: eventId }),
  });

  if (!res.ok) {
    const error = await res.json();
    throw new Error(error.error || "Failed to book ticket");
  }

  return res.json();
}

export async function getUserTickets(): Promise<Ticket[]> {
  const res = await fetch(`${API_URL}tickets`, {
    method: "GET",
    headers: getAuthHeaders(),
  });

  if (!res.ok) {
    const error = await res.json();
    throw new Error(error.error || "Failed to fetch tickets");
  }

  const data = await res.json();
  return data.tickets || [];
}

export async function cancelTicket(ticketId: number): Promise<void> {
  const res = await fetch(`${API_URL}tickets/${ticketId}`, {
    method: "DELETE",
    headers: getAuthHeaders(),
  });

  if (!res.ok) {
    const error = await res.json();
    throw new Error(error.error || "Failed to cancel ticket");
  }
}
