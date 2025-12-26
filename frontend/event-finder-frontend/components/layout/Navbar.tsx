"use client";

import { Button } from "@/components/ui/button";
import logo from "@/public/event.png";
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useEffect, useState } from "react";

const links = [
    { id: 1, title: "Home", href: "/" },
    { id: 2, title: "MyTickets", href: "/tickets" },
    { id: 3, title: "Events", href: "/events" },
    { id: 4, title: "Help", href: "/help" },
];

export default function Navbar() {
    const pathname = usePathname();

    const [scrolled, setScrolled] = useState(false);
    const [open, setOpen] = useState(false);

    const isActive = (href: string) => pathname === href;

    useEffect(() => {
        const handler = () => setScrolled(window.scrollY > 10);
        window.addEventListener("scroll", handler);
        return () => window.removeEventListener("scroll", handler);
    }, []);

    // Mock user — replace with real logic
    const user = null;

    return (
        <>
            {/* NAVBAR */}
            <header
                className={`sticky top-0 z-50 w-full border-b border-border transition-all duration-300 
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
                        {!user ? (
                            <>
                                <Button variant="default">
                                    Login
                                </Button>
                                <Button variant="outline">
                                    Signup
                                </Button>
                            </>
                        ) : (
                            <>
                                <div className="w-10 h-10 rounded-full bg-primary text-primary-foreground flex items-center justify-center font-bold text-lg">
                                    {user.name.charAt(0).toUpperCase()}
                                </div>
                                <button className="text-sm text-foreground/80 hover:text-primary">
                                    Logout
                                </button>
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
                            {!user ? (
                                <div className="flex flex-col gap-3 pt-3">
                                    <Button variant="default">
                                        Login
                                    </Button>
                                    <Button>
                                        Signup
                                    </Button>
                                </div>
                            ) : (
                                <div className="flex items-center gap-3 pt-4">
                                    <div className="w-10 h-10 rounded-full bg-primary text-primary-foreground flex items-center justify-center font-bold text-lg">
                                        {user.name.charAt(0)}
                                    </div>
                                    <button className="text-sm font-medium text-foreground/80 hover:text-primary">
                                        Logout
                                    </button>
                                </div>
                            )}
                        </ul>
                    </div>
                )}
            </header>
        </>
    );
}
