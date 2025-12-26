import { LoginRequest, RegisterRequest } from "../types/types";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/";

export async function login(data: LoginRequest) {
  const res = await fetch(`${API_URL}login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  if (!res.ok) {
    const error = await res.json();
    throw new Error(error.error || "Login failed");
  }

  return res.json();
}

export async function register(data: RegisterRequest) {
  const res = await fetch(`${API_URL}register`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  if (!res.ok) {
    const error = await res.json();
    throw new Error(error.error || "Registration failed");
  }

  return res.json();
}
