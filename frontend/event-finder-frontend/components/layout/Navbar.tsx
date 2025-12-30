"use client";

import LoginForm from "@/components/features/auth/LoginForm";
import SignupForm from "@/components/features/auth/SignUpForm";
import { Button } from "@/components/ui/button";
import Modal from "@/components/ui/modal";
import { ThemeToggle } from "@/components/ui/theme-toggle";
import { User } from "@/lib/types/types";
import logo from "@/public/event.png";
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useEffect, useState } from "react";
import { useAuth } from "@/contexts/AuthContext";

const links = [
    { id: 1, title: "Home", href: "/" },
    { id: 2, title: "MyTickets", href: "/tickets" },
    { id: 3, title: "Events", href: "/events" },
    { id: 4, title: "About us", href: "/about" },
    { id: 5, title: "Customer Support", href: "/support" },
];

export default function Navbar() {
    const pathname = usePathname();
    const { user, loading, login, logout } = useAuth();

    const [scrolled, setScrolled] = useState(false);
    const [open, setOpen] = useState(false);
    const [showLoginModal, setShowLoginModal] = useState(false);
    const [showSignupModal, setShowSignupModal] = useState(false);

    const isActive = (href: string) => pathname === href;

    useEffect(() => {
        const handler = () => setScrolled(window.scrollY > 10);
        window.addEventListener("scroll", handler);
        return () => window.removeEventListener("scroll", handler);
    }, []);

    const handleLogout = async () => {
        await logout();
    };

    const handleLoginSuccess = (loggedInUser: User) => {
        const token = localStorage.getItem("token");
        if (token) {
            login(loggedInUser, token);
            setShowLoginModal(false);
        }
    };

    const handleSignupSuccess = () => {
        setShowSignupModal(false);
        setTimeout(() => {
            setShowLoginModal(true);
        }, 500);
    };

    return (
        <>
            {/* NAVBAR */}
            <header
                className={`sticky shadow-lg top-0 z-50 w-full border-b border-border transition-all duration-300 
          ${scrolled ? "bg-background shadow-md" : "bg-background/95 backdrop-blur-sm"}
        `}
            >
                <nav className="mx-auto flex items-center justify-between px-4 py-4">
                    {/* LOGO */}
                    <div className="flex items-center gap-3">
                        <Image
                            src={logo}
                            alt="logo"
                            width={40}
                            height={40}
                            className="rounded-md"
                            priority
                        />
                        <span className="font-semibold text-lg tracking-tight text-foreground">
                            Event-<span className="text-primary">Finder</span>
                        </span>
                    </div>

                    {/* DESKTOP LINKS */}
                    <ul className="hidden md:flex items-center gap-6">
                        {links.map((l) => (
                            <li key={l.id}>
                                <Link
                                    href={l.href}
                                    className={`relative font-medium transition-colors
                    ${isActive(l.href)
                                            ? "text-primary"
                                            : "text-foreground/80 hover:text-primary"
                                        }
                  `}
                                >
                                    {l.title}

                                    {/* underline animation */}
                                    <span
                                        className={`absolute left-0 bottom-0 h-[2px] bg-primary transition-all duration-300
                      ${isActive(l.href) ? "w-full" : "w-0 group-hover:w-full"}
                    `}
                                    />
                                </Link>
                            </li>
                        ))}
                    </ul>

                    {/* DESKTOP AUTH */}
                    <div className="hidden md:flex items-center gap-4">
                        <ThemeToggle />
                        {loading ? (
                            <span className="text-sm text-muted-foreground">Loading...</span>
                        ) : !user ? (
                            <>
                                <Button onClick={() => setShowLoginModal(true)}>
                                    Login
                                </Button>
                                <Button variant="outline" onClick={() => setShowSignupModal(true)}>
                                    Signup
                                </Button>
                            </>
                        ) : (
                            <>
                                <div className="w-10 h-10 rounded-full bg-primary text-primary-foreground flex items-center justify-center font-bold text-lg">
                                    {user?.username?.charAt(0)?.toUpperCase()}
                                </div>
                                <Button variant={"outline"}
                                    onClick={handleLogout}
                                    className="text-sm text-foreground/80 hover:text-primary transition-colors"
                                >
                                    Logout
                                </Button>
                            </>
                        )}
                    </div>

                    {/* MOBILE BUTTON */}
                    <button
                        className="md:hidden p-2 text-foreground hover:text-primary transition"
                        onClick={() => setOpen(!open)}
                    >
                        {open ? (
                            <svg
                                className="w-6 h-6"
                                fill="none"
                                stroke="currentColor"
                                viewBox="0 0 24 24"
                            >
                                <path
                                    strokeLinecap="round"
                                    strokeLinejoin="round"
                                    strokeWidth="2"
                                    d="M6 18L18 6M6 6l12 12"
                                />
                            </svg>
                        ) : (
                            <svg
                                className="w-6 h-6"
                                fill="none"
                                stroke="currentColor"
                                viewBox="0 0 24 24"
                            >
                                <path
                                    strokeLinecap="round"
                                    strokeLinejoin="round"
                                    strokeWidth="2"
                                    d="M4 6h16M4 12h16M4 18h16"
                                />
                            </svg>
                        )}
                    </button>
                </nav>

                {/* MOBILE MENU */}
                {open && (
                    <div className="md:hidden border-t border-border bg-background p-5 animate-fade-in">
                        <ul className="flex flex-col gap-4">
                            {links.map((l) => (
                                <li key={l.id}>
                                    <Link
                                        href={l.href}
                                        onClick={() => setOpen(false)}
                                        className={`block text-lg font-medium transition
                      ${isActive(l.href)
                                                ? "text-primary"
                                                : "text-foreground hover:text-primary"
                                            }
                    `}
                                    >
                                        {l.title}
                                    </Link>
                                </li>
                            ))}

                            {/* MOBILE AUTH */}
                            <div className="flex items-center justify-between pt-3">
                                <ThemeToggle />
                                {loading ? (
                                    <span className="text-sm text-muted-foreground">Loading...</span>
                                ) : !user ? (
                                    <div className="flex flex-col gap-3">
                                        <Button onClick={() => {
                                            setShowLoginModal(true);
                                            setOpen(false);
                                        }}>
                                            Login
                                        </Button>
                                        <Button variant="outline" onClick={() => {
                                            setShowSignupModal(true);
                                            setOpen(false);
                                        }}>
                                            Signup
                                        </Button>
                                    </div>
                                ) : (
                                    <div className="flex items-center gap-3">
                                        <div className="w-10 h-10 rounded-full bg-primary text-primary-foreground flex items-center justify-center font-bold text-lg">
                                            {user?.username?.charAt(0)?.toUpperCase()}
                                        </div>
                                        <button
                                            onClick={async () => {
                                                await handleLogout();
                                                setOpen(false);
                                            }}
                                            className="text-sm font-medium text-foreground/80 hover:text-primary"
                                        >
                                            Logout
                                        </button>
                                    </div>
                                )}
                            </div>
                        </ul>
                    </div>
                )}
            </header>

            {/* MODALS */}
            <Modal
                title="Login to your account"
                isOpen={showLoginModal}
                onClose={() => setShowLoginModal(false)}
            >
                <LoginForm onSuccess={handleLoginSuccess} />
            </Modal>

            <Modal
                title="Create a new account"
                isOpen={showSignupModal}
                onClose={() => setShowSignupModal(false)}
            >
                <SignupForm onSuccess={handleSignupSuccess} />
            </Modal>
        </>
    );
}
