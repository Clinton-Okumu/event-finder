"use client";

import { register } from "@/lib/api/auth";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import React, { useState } from "react";

interface SignUpFormProps {
    onSuccess?: () => void;
}

const SignupForm = ({ onSuccess }: SignUpFormProps) => {
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [success, setSuccess] = useState(false);
    const [loading, setLoading] = useState(false);

    const handleRegister = async (e: React.FormEvent) => {
        e.preventDefault();
        setError("");
        setSuccess(false);
        setLoading(true);

        try {
            await register({ username, email, password });
            setSuccess(true);
            if (onSuccess) {
                setTimeout(() => onSuccess(), 1500);
            }
        } catch (err: unknown) {
            setError(err instanceof Error ? err.message : "Registration failed");
        } finally {
            setLoading(false);
        }
    };

    return (
        <form className="space-y-4" onSubmit={handleRegister}>
            <div className="space-y-2">
                <Label htmlFor="username">Username</Label>
                <Input
                    type="text"
                    id="username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    placeholder="johndoe"
                    required
                    disabled={loading}
                />
            </div>

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
                <p className="text-xs text-muted-foreground">
                    Password must be at least 6 characters
                </p>
            </div>

            {error && (
                <p className="text-destructive text-sm">{error}</p>
            )}

            {success && (
                <p className="text-green-600 dark:text-green-400 text-sm">
                    Account created! Redirecting to login...
                </p>
            )}

            <Button type="submit" className="w-full" disabled={loading}>
                {loading ? "Creating account..." : "Create Account"}
            </Button>
        </form>
    );
};

export default SignupForm;
