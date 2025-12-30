import { API_URL } from "../constants";
import { LoginRequest, RegisterRequest } from "../types/types";

function getAuthHeaders() {
  const token = localStorage.getItem("token");
  return {
    "Content-Type": "application/json",
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
  };
}

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

export async function validateToken() {
  const token = localStorage.getItem("token");
  if (!token) {
    return null;
  }

  const res = await fetch(`${API_URL}validate-token`, {
    method: "GET",
    headers: getAuthHeaders(),
  });

  if (!res.ok) {
    return null;
  }

  return res.json();
}

export async function logout() {
  const token = localStorage.getItem("token");
  if (!token) {
    return null;
  }

  try {
    const res = await fetch(`${API_URL}logout`, {
      method: "POST",
      headers: getAuthHeaders(),
    });

    if (!res.ok) {
      console.error("Logout request failed");
    }
  } catch (error) {
    console.error("Logout error:", error);
  } finally {
    localStorage.removeItem("token");
    localStorage.removeItem("user");
  }

  return null;
}
