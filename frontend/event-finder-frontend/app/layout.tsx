import Footer from "@/components/layout/Footer";
import Navbar from "@/components/layout/Navbar";
import { ThemeProvider } from "@/components/layout/ThemeProvider";
import { AuthProvider } from "@/contexts/AuthContext";
import type { Metadata } from "next";
import { Geist, Geist_Mono, Outfit } from "next/font/google";
import "./globals.css";

const outfit = Outfit({ subsets: ['latin'], variable: '--font-sans' });

const geistSans = Geist({
    variable: "--font-geist-sans",
    subsets: ["latin"],
});

const geistMono = Geist_Mono({
    variable: "--font-geist-mono",
    subsets: ["latin"],
});

export const metadata: Metadata = {
    title: "Event Finder",
    description: "Discover amazing events near you",
    icons: {
        icon: [
            { url: '/favicon.ico', type: 'image/png' },
        ],
    },
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="en" className={outfit.variable} suppressHydrationWarning>
            <body
                className={`${geistSans.variable} ${geistMono.variable} antialiased`}
            >
                <ThemeProvider
                    attribute="class"
                    defaultTheme="system"
                    enableSystem
                    disableTransitionOnChange
                >
                    <AuthProvider>
                        <Navbar />
                        {children}
                        <Footer />
                    </AuthProvider>
                </ThemeProvider>
            </body>
        </html>
    );
}
