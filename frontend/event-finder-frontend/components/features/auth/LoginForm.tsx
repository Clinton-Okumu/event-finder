"use client";

import { login } from "@/lib/api/auth";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { User } from "@/lib/types/types";
import React, { useState } from "react";
import { useAuth } from "@/contexts/AuthContext";

interface LoginFormProps {
    onSuccess?: (user: User) => void;
}

const LoginForm = ({ onSuccess }: LoginFormProps) => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);
    const { login: authLogin } = useAuth();

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError("");
        setLoading(true);

        try {
            const res = await login({ email, password });
            authLogin(res.user, res.token);

            if (onSuccess) {
                onSuccess(res.user);
            }
        } catch (err: unknown) {
            setError(err instanceof Error ? err.message : "Login failed");
        } finally {
            setLoading(false);
        }
    };

    return (
        <form className="space-y-4" onSubmit={handleSubmit}>
            {error && (
                <div className="text-destructive text-sm font-medium p-3 rounded-md bg-destructive/10 border border-destructive/20">
                    {error}
                </div>
            )}

            <div className="space-y-2">
                <Label htmlFor="email">Email Address</Label>
                <Input
                    type="email"
                    id="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    placeholder="you@example.com"
                    required
                    disabled={loading}
                />
            </div>

            <div className="space-y-2">
                <Label htmlFor="password">Password</Label>
                <Input
                    type="password"
                    id="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    placeholder="•••••••"
                    required
                    disabled={loading}
                />
            </div>

            <Button type="submit" className="w-full" disabled={loading}>
                {loading ? "Logging in..." : "Login"}
            </Button>
        </form>
    );
};

export default LoginForm;
