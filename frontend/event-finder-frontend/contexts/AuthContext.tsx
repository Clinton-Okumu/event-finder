"use client";

import React, { createContext, useContext, useEffect, useState } from "react";
import { User } from "@/lib/types/types";
import { validateToken, logout as apiLogout } from "@/lib/api/auth";

interface AuthContextType {
  user: User | null;
  loading: boolean;
  login: (user: User, token: string) => void;
  logout: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const checkAuth = async () => {
      try {
        const data = await validateToken();
        if (data && data.user) {
          setUser(data.user);
        } else {
          localStorage.removeItem("token");
          localStorage.removeItem("user");
        }
      } catch (error) {
        console.error("Auth check failed:", error);
        localStorage.removeItem("token");
        localStorage.removeItem("user");
      } finally {
        setLoading(false);
      }
    };

    checkAuth();
  }, []);

  const login = (user: User, token: string) => {
    localStorage.setItem("token", token);
    localStorage.setItem("user", JSON.stringify(user));
    setUser(user);
  };

  const logout = async () => {
    await apiLogout();
    setUser(null);
  };

  return (
    <AuthContext.Provider value={{ user, loading, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}
